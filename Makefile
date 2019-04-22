GO_VERSION:=$(shell go version)

.PHONY: all clean init lint test contributors

all: clean init lint test

clean:
	go clean -modcache
	rm -rf ./*.log
	rm -rf ./*.svg
	rm -rf ./go.mod
	rm -rf ./go.sum
	rm -rf bench
	rm -rf pprof
	rm -rf vendor


init:
	GO111MODULE=on go mod init
	GO111MODULE=on go mod vendor

lint:
	gometalinter --enable-all . | rg -v comment

test: clean init
	GO111MODULE=on go test --race -v $(go list ./... | rg -v vendor)

contributors:
	git log --format='%aN <%aE>' | sort -fu > CONTRIBUTORS

deps: clean
	go mod init
	go mod vendor
	rm -rf vendor
