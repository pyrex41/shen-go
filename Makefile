.PHONY: all kl shen docker test certify test-all precompile yggdrasil-build

all: kl shen

kl:
	go install github.com/tiancaiamao/shen-go/cmd/kl

shen:
	go build -o shen github.com/tiancaiamao/shen-go/cmd/shen

# Yggdrasil stage-2 static builder: turn a Yggdrasil shaken directory into a
# Go module that builds a single static, cross-compilable binary.
#   make yggdrasil-build
#   ./yggdrasil-build path/to/shaken-out path/to/gen && (cd path/to/gen && go build -o prog .)
yggdrasil-build:
	go build -o yggdrasil-build ./cmd/yggdrasil-build

# AOT-compile a whole Shen (or KL) file to a Go plugin .so, to be loaded at
# startup with `./shen -precompiled OUT`. The .so MUST be built with the same Go
# toolchain + module as the shen binary that loads it (both from this tree).
#   make precompile FILE=bench/hot.shen           # -> bench/hot.so
#   make precompile FILE=bench/hot.shen OUT=x.so  # -> x.so
FILE ?=
OUT  ?= $(basename $(FILE)).so
precompile:
	@test -n "$(FILE)" || { echo "usage: make precompile FILE=path/to/file.shen [OUT=out.so]"; exit 1; }
	@go build -o klc ./cmd/kl
	@rm -rf compiled/_in* compiled/plugintmp*
	@cp "$(FILE)" compiled/_in.shen
	@cd compiled && ../klc < precompile.kl > /tmp/shen-precompile.log 2>&1 || { cat /tmp/shen-precompile.log; exit 1; }
	@grep -q '\[go-build-plugin\]' /tmp/shen-precompile.log || { echo "precompile failed:"; cat /tmp/shen-precompile.log; exit 1; }
	@mkdir -p "$(dir $(OUT))"
	@mv compiled/_in.so "$(OUT)"
	@mv compiled/_in.fns "$(OUT).fns"
	@rm -rf compiled/_in* compiled/plugintmp*
	@echo "precompiled $(FILE) -> $(OUT) (+ $(OUT).fns)"

shen-exe:
	go build -o shen.exe github.com/tiancaiamao/shen-go/cmd/shen



docker:
	docker build -t shen-go .
	docker run -i -t --rm -v /tmp:/tmp shen-go \
		/bin/sh -c 'cp -a /usr/local/bin/shen /tmp/'
	cp -a /tmp/shen ./shen

# certify: run ONLY the canonical Shen kernel certification suite (the official
# ShenOSKernel acceptance tests under kernel/tests). This is the external bar
# for "Certified" — distinct from our own tests below.
#
# -count=1 disables Go's test result cache: the cert execs the freshly-built
# binary against kernel/tests (inputs Go's cache can't see), so a cached PASS
# could be stale. Certification must always actually run.
certify:
	go test -v -count=1 ./certification/

# test: run ONLY our own Go unit tests (the kl evaluator/VM and the shen CLI).
# Skips the slow canonical certification (-short).
test:
	go test -short ./kl/ ./cmd/shen/

# test-all: everything — our unit tests (cached) plus the canonical
# certification (always run for real). Composed from the targets above so the
# cert is never served stale from cache.
test-all: test certify

