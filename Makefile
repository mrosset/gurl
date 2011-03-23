include $(GOROOT)/src/Make.inc

TARG=gurl
GOFILES=gurl.go
GOFMT=gofmt -l -w


format:
	    ${GOFMT} .

include $(GOROOT)/src/Make.pkg
