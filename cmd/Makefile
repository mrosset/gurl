include $(GOROOT)/src/Make.inc

TARG=gurl
GOFILES=main.go
GOFMT=gofmt -l -w

include $(GOROOT)/src/Make.cmd

test: clean all
	./${TARG} http://localhost/gcc-go-snapshot-4.7.20110423-1-x86_64.pkg.tar.xz

format:
	    ${GOFMT} .

