prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/whosonfirst/go-whosonfirst-roundhouse; then rm -rf src/github.com/whosonfirst/go-whosonfirst-roundhouse; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-roundhouse
	cp roundhouse.go src/github.com/whosonfirst/go-whosonfirst-roundhouse/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	rmdeps deps fmt bin

deps:   self
	@GOPATH=$(shell pwd) go get -u "github.com/whosonfirst/go-whosonfirst-utils"

fmt:
	go fmt cmd/*.go
	go fmt *.go

bin: 	self
	@GOPATH=$(shell pwd) go build -o bin/wof-roundhouse cmd/wof-roundhouse.go
	@GOPATH=$(shell pwd) go build -o bin/wof-roundhouse-server cmd/wof-roundhouse-server.go