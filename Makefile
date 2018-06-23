prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/aaronland/go-artisanal-integers-redis; then rm -rf src/github.com/aaronland/go-artisanal-integers-redis; fi
	mkdir -p src/github.com/aaronland/go-artisanal-integers-redis/
	cp *.go src/github.com/aaronland/go-artisanal-integers-redis/
	cp -r engine src/github.com/aaronland/go-artisanal-integers-redis/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:
	@GOPATH=$(shell pwd) go get "github.com/aaronland/go-artisanal-integers"
	@GOPATH=$(shell pwd) go get "github.com/gomodule/redigo/redis"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt engine/*.go

bin:    self
	if test ! -d bin; then mkdir bin; fi
	@GOPATH=$(shell pwd) go build -o bin/intd-server cmd/intd-server.go
