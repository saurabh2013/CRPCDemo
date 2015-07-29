GOROOT=/usr/local/go
GOPATH=$HOME/demo #<-- Provide your working dir path where github.com dir exists
GOCMD=go

all: compile
	export GOPATH=$(GOPATH) 
	./CRPCDemo

compile:
	$(GOCMD) build
clean: 
	$(GOCMD) clean

