prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-roundhouse
	cp roundhouse.go src/github.com/whosonfirst/go-whosonfirst-roundhouse/
	cp -r vendor/src/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	rmdeps deps fmt bin

deps:   
	@GOPATH=$(shell pwd) go get -u "github.com/whosonfirst/go-whosonfirst-crawl"
	@GOPATH=$(shell pwd) go get -u "github.com/whosonfirst/go-whosonfirst-csv"
	@GOPATH=$(shell pwd) go get -u "github.com/whosonfirst/go-whosonfirst-uri"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt *.go

bin: 	self
	@GOPATH=$(shell pwd) go build -o bin/wof-roundhouse cmd/wof-roundhouse.go
	@GOPATH=$(shell pwd) go build -o bin/wof-roundhoused cmd/wof-roundhoused.go
	@GOPATH=$(shell pwd) go build -o bin/wof-roundhouse-repod cmd/wof-roundhouse-repod.go
