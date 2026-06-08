.PHONY: all kl shen docker test certify test-all

all: kl shen

kl:
	go install github.com/tiancaiamao/shen-go/cmd/kl

shen:
	go build -o shen github.com/tiancaiamao/shen-go/cmd/shen

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
certify:
	go test -v ./certification/

# test: run ONLY our own Go unit tests (the kl evaluator/VM and the shen CLI).
# Skips the slow canonical certification (-short).
test:
	go test -short ./kl/ ./cmd/shen/

# test-all: everything — our unit tests plus the canonical certification.
test-all:
	go test ./...

