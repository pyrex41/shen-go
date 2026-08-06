package kl

import (
	"context"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/go-zeromq/zmq4"
)

// shen.x.zmq host waist for shen-go, backed by github.com/go-zeromq/zmq4
// (pure Go, no cgo/libzmq). Binds the 10 host primitives described in the
// shen.x.zmq spec and sets shen.x.*zmq-backend* to the symbol host.
// Disable with SHEN_X_ZMQ=off.

type zmqRecvResult struct {
	msg zmq4.Msg
	err error
}

type zsock struct {
	sock     zmq4.Socket
	endpoint string // last bound endpoint ("" = not bound)
	rcvtimeo int    // ms; -1 = block forever (default)

	// In-flight recv pattern: at most ONE outstanding Recv() goroutine per
	// socket, feeding a 1-buffered channel. A recv that times out leaves the
	// goroutine in flight; the next recv-host/poll-host on this socket
	// consumes the same channel (never a second concurrent Recv). poll-host
	// stashes an arrived message in pending; recv-host drains pending first.
	// These fields are only touched from the (single) interpreter thread.
	inflight chan zmqRecvResult
	pending  *zmqRecvResult
}

var (
	zmqMu     sync.Mutex
	zmqSocks  = make(map[int]*zsock)
	zmqNextID = 1
)

func installShenXZmq() {
	if os.Getenv("SHEN_X_ZMQ") == "off" {
		return
	}
	bind1 := func(name string, fn func(Obj) Obj) {
		BindSymbolFunc(MakeSymbol(name), MakePrimitive(name, 1, fn))
	}
	bind2 := func(name string, fn func(x, y Obj) Obj) {
		BindSymbolFunc(MakeSymbol(name), MakePrimitive(name, 2, fn))
	}
	bind3 := func(name string, fn func(x, y, z Obj) Obj) {
		BindSymbolFunc(MakeSymbol(name), MakePrimitive(name, 3, fn))
	}
	bind1("shen.x.zmq.socket-host", primZmqSocketHost)
	bind2("shen.x.zmq.bind-host", primZmqBindHost)
	bind2("shen.x.zmq.connect-host", primZmqConnectHost)
	bind2("shen.x.zmq.send-host", primZmqSendHost)
	bind1("shen.x.zmq.recv-host", primZmqRecvHost)
	bind3("shen.x.zmq.setopt-host", primZmqSetoptHost)
	bind2("shen.x.zmq.poll-host", primZmqPollHost)
	bind1("shen.x.zmq.endpoint-host", primZmqEndpointHost)
	bind1("shen.x.zmq.close-host", primZmqCloseHost)
	// term-host takes no arguments; MakePrimitive only wraps arity 1..3.
	BindSymbolFunc(MakeSymbol("shen.x.zmq.term-host"),
		MakeNative(func(e *ControlFlow) { e.Return(primZmqTermHost()) }, 0))
	PrimSet(MakeSymbol("shen.x.*zmq-backend*"), MakeSymbol("host"))
}

// ---------- helpers ----------

func zmqError(format string, args ...interface{}) Obj {
	return MakeError("shen.x.zmq: " + fmt.Sprintf(format, args...))
}

func zmqTimeoutSym() Obj { return MakeSymbol("shen.x.zmq.timeout") }

// zmqImmediateGrace is the tiny wait used for "immediate" recv/poll (ms == 0).
// A pure non-blocking channel check would deterministically miss a message the
// just-spawned in-flight goroutine is about to deliver (the goroutine has not
// been scheduled yet), so ms == 0 waits this long instead — long enough for an
// already-queued message to surface, short enough to still read as a snapshot.
const zmqImmediateGrace = time.Millisecond

// zmqInt narrows a Shen number to an integral Go int. ok=false for
// non-numbers AND fractional values: GetInteger would silently truncate, so
// handle 1.5 would alias live socket 1 and octet 1.5 would become byte 1.
func zmqInt(o Obj) (int, bool) {
	if !IsNumber(o) {
		return 0, false
	}
	f := GetNumber(o)
	if !isPreciseInteger(f) || !fitsInt(f) {
		return 0, false
	}
	return int(f), true
}

// zmqMsDuration converts a millisecond count to a time.Duration, clamping
// values whose nanosecond product would overflow int64 (~292y — effectively
// block forever) instead of wrapping negative and firing instantly.
func zmqMsDuration(ms int) time.Duration {
	if ms > int(math.MaxInt64/int64(time.Millisecond)) {
		return time.Duration(math.MaxInt64)
	}
	return time.Duration(ms) * time.Millisecond
}

// zmqGetSock resolves a handle object to its socket, or panics with the
// contract error for closed/unknown handles.
func zmqGetSock(h Obj) (*zsock, int) {
	id, ok := zmqInt(h)
	if !ok {
		panic(zmqError("closed or unknown socket handle: %s", ObjString(h)))
	}
	zmqMu.Lock()
	z, ok := zmqSocks[id]
	zmqMu.Unlock()
	if !ok {
		panic(zmqError("closed or unknown socket handle: %d", id))
	}
	return z, id
}

// zmqOctets converts a proper list of ints 0..255 to bytes; ok=false on any
// malformed element (caller picks the contract error string).
func zmqOctets(l Obj) ([]byte, bool) {
	buf := []byte{}
	for l != Nil {
		if *l != scmHeadPair {
			return nil, false
		}
		p := mustPair(l)
		n, ok := zmqInt(p.car)
		if !ok || n < 0 || n > 255 {
			return nil, false
		}
		buf = append(buf, byte(n))
		l = p.cdr
	}
	return buf, true
}

func zmqFramesToObj(frames [][]byte) Obj {
	out := Nil
	for i := len(frames) - 1; i >= 0; i-- {
		frame := Nil
		b := frames[i]
		for j := len(b) - 1; j >= 0; j-- {
			frame = cons(MakeInteger(int(b[j])), frame)
		}
		out = cons(frame, out)
	}
	return out
}

// zmqTakeResult consumes a finished recv result: backend errors pass through
// as Shen errors, successes become a list of frames.
func zmqTakeResult(r zmqRecvResult) Obj {
	if r.err != nil {
		panic(zmqError("%s", r.err.Error()))
	}
	return zmqFramesToObj(r.msg.Frames)
}

// zmqEnsureInflight spawns the single recv goroutine if none is in flight.
// Only called when the user asks to recv/poll, so the REQ/REP state machine
// stays valid. The channel is 1-buffered: the goroutine can always deliver
// and exit, so close-host/term-host never deadlock on it.
func zmqEnsureInflight(z *zsock) {
	if z.inflight != nil {
		return
	}
	ch := make(chan zmqRecvResult, 1)
	z.inflight = ch
	s := z.sock
	go func() {
		msg, err := s.Recv()
		ch <- zmqRecvResult{msg: msg, err: err}
	}()
}

// ---------- primitives ----------

func primZmqSocketHost(t Obj) Obj {
	if !IsSymbol(t) {
		panic(zmqError("unknown socket type: %s", ObjString(t)))
	}
	name := GetSymbol(t)
	ctx := context.Background()
	var s zmq4.Socket
	switch name {
	case "req":
		s = zmq4.NewReq(ctx)
	case "rep":
		s = zmq4.NewRep(ctx)
	case "pub":
		s = zmq4.NewPub(ctx)
	case "sub":
		s = zmq4.NewSub(ctx)
	case "push":
		s = zmq4.NewPush(ctx)
	case "pull":
		s = zmq4.NewPull(ctx)
	case "pair":
		s = zmq4.NewPair(ctx)
	case "dealer":
		s = zmq4.NewDealer(ctx)
	case "router":
		s = zmq4.NewRouter(ctx)
	default:
		panic(zmqError("unknown socket type: %s", name))
	}
	zmqMu.Lock()
	id := zmqNextID
	zmqNextID++
	zmqSocks[id] = &zsock{sock: s, rcvtimeo: -1}
	zmqMu.Unlock()
	return MakeInteger(id)
}

func primZmqBindHost(h, ep Obj) Obj {
	z, _ := zmqGetSock(h)
	if !IsString(ep) {
		panic(zmqError("bind-host expects an endpoint string"))
	}
	endpoint := GetString(ep)
	if err := z.sock.Listen(endpoint); err != nil {
		panic(zmqError("%s", err.Error()))
	}
	z.endpoint = endpoint
	// Recover the real endpoint after an ephemeral tcp bind (":0").
	if strings.HasPrefix(endpoint, "tcp://") {
		if addr := z.sock.Addr(); addr != nil {
			z.endpoint = "tcp://" + addr.String()
		}
	}
	return True
}

func primZmqConnectHost(h, ep Obj) Obj {
	z, _ := zmqGetSock(h)
	if !IsString(ep) {
		panic(zmqError("connect-host expects an endpoint string"))
	}
	if err := z.sock.Dial(GetString(ep)); err != nil {
		panic(zmqError("%s", err.Error()))
	}
	return True
}

func primZmqSendHost(h, msg Obj) Obj {
	z, _ := zmqGetSock(h)
	malformed := func() { panic(zmqError("expected list of frames (lists of ints 0..255)")) }
	if msg == Nil || *msg != scmHeadPair {
		malformed()
	}
	var frames [][]byte
	for l := msg; l != Nil; {
		if *l != scmHeadPair {
			malformed()
		}
		p := mustPair(l)
		frame, ok := zmqOctets(p.car)
		if !ok {
			malformed()
		}
		frames = append(frames, frame)
		l = p.cdr
	}
	var err error
	if len(frames) == 1 {
		err = z.sock.Send(zmq4.NewMsg(frames[0]))
	} else {
		err = z.sock.SendMulti(zmq4.NewMsgFrom(frames...))
	}
	if err != nil {
		panic(zmqError("%s", err.Error()))
	}
	return True
}

func primZmqRecvHost(h Obj) Obj {
	z, _ := zmqGetSock(h)
	if z.pending != nil {
		r := *z.pending
		z.pending = nil
		return zmqTakeResult(r)
	}
	zmqEnsureInflight(z)
	if z.rcvtimeo < 0 { // block forever
		r := <-z.inflight
		z.inflight = nil
		return zmqTakeResult(r)
	}
	// rcvtimeo == 0 is standard non-blocking recv: return a queued message if
	// present. The in-flight goroutine may only just have been spawned, so an
	// unbuffered select-default would miss an already-queued message; wait the
	// grace period instead (see zmqImmediateGrace).
	d := zmqMsDuration(z.rcvtimeo)
	if z.rcvtimeo == 0 {
		d = zmqImmediateGrace
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case r := <-z.inflight:
		z.inflight = nil
		return zmqTakeResult(r)
	case <-timer.C:
		return zmqTimeoutSym() // keep the in-flight recv
	}
}

func primZmqSetoptHost(h, opt, val Obj) Obj {
	z, _ := zmqGetSock(h)
	if !IsSymbol(opt) {
		panic(zmqError("unknown option: %s", ObjString(opt)))
	}
	name := GetSymbol(opt)
	switch name {
	case "rcvtimeo":
		ms, ok := zmqInt(val)
		if !ok {
			panic(zmqError("rcvtimeo expects an integer (ms)"))
		}
		z.rcvtimeo = ms
	case "sndtimeo":
		// zmq4 has no per-socket send deadline that can be set after
		// construction; a goroutine+select send could still deliver after a
		// reported timeout, which is ambiguous. Take the spec's escape hatch.
		panic(zmqError("option sndtimeo not supported by this backend"))
	case "subscribe", "unsubscribe":
		topic, ok := zmqOctets(val)
		if !ok {
			panic(zmqError("%s expects a list of ints 0..255", name))
		}
		zopt := zmq4.OptionSubscribe
		if name == "unsubscribe" {
			zopt = zmq4.OptionUnsubscribe
		}
		if err := z.sock.SetOption(zopt, string(topic)); err != nil {
			panic(zmqError("%s", err.Error()))
		}
	case "linger":
		if _, ok := zmqInt(val); !ok {
			panic(zmqError("linger expects an integer (ms)"))
		}
		// Accepted no-op: zmq4 has no linger concept (spec allows this).
	default:
		panic(zmqError("unknown option: %s", name))
	}
	return True
}

func primZmqPollHost(hs, msObj Obj) Obj {
	ms, msOK := zmqInt(msObj)
	if !msOK {
		panic(zmqError("poll-host expects an integer (ms)"))
	}
	type entry struct {
		h Obj
		z *zsock
	}
	var entries []entry
	for l := hs; l != Nil; {
		if *l != scmHeadPair {
			panic(zmqError("poll-host expects a list of handles"))
		}
		p := mustPair(l)
		z, _ := zmqGetSock(p.car)
		entries = append(entries, entry{h: p.car, z: z})
		l = p.cdr
	}
	ready := make([]bool, len(entries))
	// Non-blocking sweep: pending counts as readable; otherwise ensure the
	// single in-flight recv and check its channel without blocking, stashing
	// any arrived message into the pending slot for recv-host to drain.
	sweep := func() bool {
		any := false
		for i := range entries {
			z := entries[i].z
			if ready[i] {
				any = true
				continue
			}
			if z.pending != nil {
				ready[i] = true
				any = true
				continue
			}
			zmqEnsureInflight(z)
			select {
			case r := <-z.inflight:
				z.inflight = nil
				r2 := r
				z.pending = &r2
				ready[i] = true
				any = true
			default:
			}
		}
		return any
	}
	if !sweep() && len(entries) > 0 {
		// Wait for the first socket to become readable (or the deadline).
		// ms == 0 ("immediate snapshot") also waits — for zmqImmediateGrace —
		// because the in-flight goroutines may only just have been spawned and
		// an already-queued message would otherwise be invisible.
		seen := make(map[*zsock]bool)
		var cases []reflect.SelectCase
		var caseEntry []int
		for i := range entries {
			z := entries[i].z
			if seen[z] {
				continue
			}
			seen[z] = true
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(z.inflight),
			})
			caseEntry = append(caseEntry, i)
		}
		timerIdx := -1
		if ms >= 0 {
			d := zmqMsDuration(ms)
			if ms == 0 {
				d = zmqImmediateGrace
			}
			timer := time.NewTimer(d)
			defer timer.Stop()
			timerIdx = len(cases)
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(timer.C),
			})
		}
		chosen, recv, _ := reflect.Select(cases)
		if chosen != timerIdx {
			i := caseEntry[chosen]
			z := entries[i].z
			r := recv.Interface().(zmqRecvResult)
			z.inflight = nil
			z.pending = &r
			ready[i] = true
			sweep() // pick up anything else that arrived meanwhile
		}
	}
	out := Nil
	for i := len(entries) - 1; i >= 0; i-- {
		if ready[i] {
			out = cons(entries[i].h, out)
		}
	}
	return out
}

func primZmqEndpointHost(h Obj) Obj {
	z, id := zmqGetSock(h)
	if z.endpoint == "" {
		panic(zmqError("socket not bound: %d", id))
	}
	return MakeString(z.endpoint)
}

func primZmqCloseHost(h Obj) Obj {
	z, id := zmqGetSock(h)
	zmqMu.Lock()
	delete(zmqSocks, id)
	zmqMu.Unlock()
	// Closing the socket makes any in-flight Recv return an error; the
	// goroutine delivers it into its buffered channel and exits (dropped).
	if err := z.sock.Close(); err != nil {
		panic(zmqError("%s", err.Error()))
	}
	return True
}

func primZmqTermHost() Obj {
	zmqMu.Lock()
	socks := zmqSocks
	zmqSocks = make(map[int]*zsock)
	zmqNextID = 1
	zmqMu.Unlock()
	for _, z := range socks {
		_ = z.sock.Close() // in-flight recvs error out and are dropped
	}
	return True
}
