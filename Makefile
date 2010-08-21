include $(GOROOT)/src/Make.$(GOARCH)

TARG=gurl
CGOFILES=gurl.go
CGO_CFLAGS  = `curl-config --cflags`
CGO_LDFLAGS = `curl-config --libs`

include $(GOROOT)/src/Make.pkg
