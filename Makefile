include $(GOROOT)/src/Make.inc

TARG=gurl
CGOFILES=gurl.go
CGO_CFLAGS  = `curl-config --cflags`
CGO_LDFLAGS = `curl-config --libs`
GOFMT=gofmt -l=true -tabwidth=4 -w



test: clean format install 
	gotest

format:
	    ${GOFMT} .

include $(GOROOT)/src/Make.pkg
