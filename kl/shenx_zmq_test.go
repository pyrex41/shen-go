package kl

import (
	"strings"
	"testing"
)

// helpers ------------------------------------------------------------------

func zmqList(items ...Obj) Obj {
	out := Nil
	for i := len(items) - 1; i >= 0; i-- {
		out = cons(items[i], out)
	}
	return out
}

func zmqOctetsObj(s string) Obj {
	out := Nil
	for i := len(s) - 1; i >= 0; i-- {
		out = cons(MakeInteger(int(s[i])), out)
	}
	return out
}

// zmqCatch runs f, expecting it to panic with a Shen error, and returns the
// error message.
func zmqCatch(t *testing.T, f func()) string {
	t.Helper()
	var msg string
	caught := false
	func() {
		defer func() {
			r := recover()
			if r == nil {
				return
			}
			if o, ok := r.(Obj); ok && IsError(o) {
				msg = mustError(o).err
				caught = true
				return
			}
			panic(r)
		}()
		f()
	}()
	if !caught {
		t.Fatalf("expected a Shen error, got none")
	}
	return msg
}

func zmqSetRcvtimeo(h Obj, ms int) {
	primZmqSetoptHost(h, MakeSymbol("rcvtimeo"), MakeInteger(ms))
}

// tests --------------------------------------------------------------------

func TestZmqPushPullRoundtrip(t *testing.T) {
	defer primZmqTermHost()

	pull := primZmqSocketHost(MakeSymbol("pull"))
	push := primZmqSocketHost(MakeSymbol("push"))
	if GetInteger(pull) < 1 {
		t.Fatalf("handle should start at 1, got %d", GetInteger(pull))
	}

	primZmqBindHost(pull, MakeString("tcp://127.0.0.1:0"))
	epObj := primZmqEndpointHost(pull)
	ep := GetString(epObj)
	if !strings.HasPrefix(ep, "tcp://127.0.0.1:") || strings.HasSuffix(ep, ":0") {
		t.Fatalf("endpoint after :0 bind should have the real port, got %q", ep)
	}
	primZmqConnectHost(push, MakeString(ep))

	// Don't hang the test suite if recv is broken.
	zmqSetRcvtimeo(pull, 5000)

	// Single frame.
	msg := zmqList(zmqOctetsObj("hello"))
	if primZmqSendHost(push, msg) != True {
		t.Fatalf("send-host should return true")
	}
	got := primZmqRecvHost(pull)
	if !Equal(got, msg) {
		t.Fatalf("recv mismatch: got %s want %s", ObjString(got), ObjString(msg))
	}

	// Multipart: 3 frames including an empty one.
	multi := zmqList(zmqOctetsObj("a"), zmqOctetsObj(""), zmqOctetsObj("bc"))
	primZmqSendHost(push, multi)
	got = primZmqRecvHost(pull)
	if !Equal(got, multi) {
		t.Fatalf("multipart recv mismatch: got %s want %s", ObjString(got), ObjString(multi))
	}
}

func TestZmqTimeoutAndInflightReuse(t *testing.T) {
	defer primZmqTermHost()

	pull := primZmqSocketHost(MakeSymbol("pull"))
	push := primZmqSocketHost(MakeSymbol("push"))
	primZmqBindHost(pull, MakeString("tcp://127.0.0.1:0"))
	ep := GetString(primZmqEndpointHost(pull))
	primZmqConnectHost(push, MakeString(ep))

	// Timeout on an empty socket returns the sentinel symbol (never raises).
	zmqSetRcvtimeo(pull, 50)
	got := primZmqRecvHost(pull)
	if got != MakeSymbol("shen.x.zmq.timeout") {
		t.Fatalf("expected timeout symbol, got %s", ObjString(got))
	}

	// The next recv must reuse the SAME in-flight recv and see the message.
	msg := zmqList(zmqOctetsObj("late"))
	primZmqSendHost(push, msg)
	zmqSetRcvtimeo(pull, 5000)
	got = primZmqRecvHost(pull)
	if !Equal(got, msg) {
		t.Fatalf("recv after timeout mismatch: got %s want %s", ObjString(got), ObjString(msg))
	}
}

func TestZmqPoll(t *testing.T) {
	defer primZmqTermHost()

	pull := primZmqSocketHost(MakeSymbol("pull"))
	push := primZmqSocketHost(MakeSymbol("push"))
	primZmqBindHost(pull, MakeString("tcp://127.0.0.1:0"))
	ep := GetString(primZmqEndpointHost(pull))
	primZmqConnectHost(push, MakeString(ep))

	// Empty socket: immediate poll and short positive poll both return [].
	if got := primZmqPollHost(zmqList(pull), MakeInteger(0)); got != Nil {
		t.Fatalf("poll 0 on empty socket: got %s want []", ObjString(got))
	}
	if got := primZmqPollHost(zmqList(pull), MakeInteger(50)); got != Nil {
		t.Fatalf("poll 50 on empty socket: got %s want []", ObjString(got))
	}

	// After a send the socket becomes readable within the poll window.
	msg := zmqList(zmqOctetsObj("ping"))
	primZmqSendHost(push, msg)
	got := primZmqPollHost(zmqList(pull), MakeInteger(5000))
	if !Equal(got, zmqList(pull)) {
		t.Fatalf("poll after send: got %s want %s", ObjString(got), ObjString(zmqList(pull)))
	}
	// A ready socket stays ready on an immediate re-poll (pending slot).
	got = primZmqPollHost(zmqList(pull), MakeInteger(0))
	if !Equal(got, zmqList(pull)) {
		t.Fatalf("re-poll of ready socket: got %s", ObjString(got))
	}
	// recv drains the pending slot stashed by poll.
	zmqSetRcvtimeo(pull, 5000)
	if got := primZmqRecvHost(pull); !Equal(got, msg) {
		t.Fatalf("recv after poll mismatch: got %s want %s", ObjString(got), ObjString(msg))
	}
	// Drained again: not readable.
	if got := primZmqPollHost(zmqList(pull), MakeInteger(0)); got != Nil {
		t.Fatalf("poll after drain: got %s want []", ObjString(got))
	}
}

func TestZmqCloseErrors(t *testing.T) {
	defer primZmqTermHost()

	s := primZmqSocketHost(MakeSymbol("pair"))
	if primZmqCloseHost(s) != True {
		t.Fatalf("close-host should return true")
	}
	want := "shen.x.zmq: closed or unknown socket handle: " + ObjString(s)
	if msg := zmqCatch(t, func() { primZmqCloseHost(s) }); msg != want {
		t.Fatalf("double close: got %q want %q", msg, want)
	}
	if msg := zmqCatch(t, func() { primZmqRecvHost(s) }); msg != want {
		t.Fatalf("recv on closed handle: got %q want %q", msg, want)
	}
	if msg := zmqCatch(t, func() { primZmqBindHost(s, MakeString("inproc://x")) }); msg != want {
		t.Fatalf("bind on closed handle: got %q want %q", msg, want)
	}
}

func TestZmqMalformedSend(t *testing.T) {
	defer primZmqTermHost()

	s := primZmqSocketHost(MakeSymbol("push"))
	want := "shen.x.zmq: expected list of frames (lists of ints 0..255)"
	cases := map[string]Obj{
		"empty message":       Nil,
		"non-list message":    MakeString("hi"),
		"non-list frame":      zmqList(MakeInteger(1)),
		"octet out of range":  zmqList(zmqList(MakeInteger(300))),
		"negative octet":      zmqList(zmqList(MakeInteger(-1))),
		"non-number in frame": zmqList(zmqList(MakeString("x"))),
	}
	for name, payload := range cases {
		if msg := zmqCatch(t, func() { primZmqSendHost(s, payload) }); msg != want {
			t.Fatalf("%s: got %q want %q", name, msg, want)
		}
	}
}

func TestZmqContractErrors(t *testing.T) {
	defer primZmqTermHost()

	if msg := zmqCatch(t, func() { primZmqSocketHost(MakeSymbol("bogus")) }); msg != "shen.x.zmq: unknown socket type: bogus" {
		t.Fatalf("unknown type: got %q", msg)
	}
	s := primZmqSocketHost(MakeSymbol("pair"))
	sid := ObjString(s)
	if msg := zmqCatch(t, func() { primZmqEndpointHost(s) }); msg != "shen.x.zmq: socket not bound: "+sid {
		t.Fatalf("endpoint before bind: got %q", msg)
	}
	if msg := zmqCatch(t, func() { primZmqSetoptHost(s, MakeSymbol("bogus"), MakeInteger(1)) }); msg != "shen.x.zmq: unknown option: bogus" {
		t.Fatalf("unknown option: got %q", msg)
	}
	if msg := zmqCatch(t, func() { primZmqSetoptHost(s, MakeSymbol("sndtimeo"), MakeInteger(1)) }); msg != "shen.x.zmq: option sndtimeo not supported by this backend" {
		t.Fatalf("sndtimeo: got %q", msg)
	}
	// linger is an accepted no-op.
	if primZmqSetoptHost(s, MakeSymbol("linger"), MakeInteger(0)) != True {
		t.Fatalf("linger should be accepted")
	}
}

func TestZmqTermIdempotentAndResets(t *testing.T) {
	s := primZmqSocketHost(MakeSymbol("pair"))
	_ = s
	if primZmqTermHost() != True {
		t.Fatalf("term-host should return true")
	}
	if primZmqTermHost() != True {
		t.Fatalf("term-host should be idempotent")
	}
	// Context reset: handles start at 1 again.
	s2 := primZmqSocketHost(MakeSymbol("pair"))
	defer primZmqTermHost()
	if GetInteger(s2) != 1 {
		t.Fatalf("after term, handles should restart at 1, got %d", GetInteger(s2))
	}
}
