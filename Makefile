include $(GOROOT)/src/Make.$(GOARCH)

TARG=gurl
CGOFILES=gurl.go
CGO_CFLAGS  = `curl-config --cflags`
CGO_LDFLAGS = `curl-config --libs`
GOFMT=gofmt -l=true -tabwidth=4 -w

include $(GOROOT)/src/Make.pkg

format:
	    ${GOFMT} .
