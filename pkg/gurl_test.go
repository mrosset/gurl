package gurl

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type testHandler struct {
}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	b := new(bytes.Buffer)
	b.Write(make([]byte, 1024*1024))
	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("Content-Length", strconv.Itoa(len(b.Bytes())))
	io.Copy(w, b)
}

func TestLocal(t *testing.T) {
	testfile := "/foobar.tar.gz"
	server := httptest.NewServer(&testHandler{})
	url := server.URL + testfile
	if err := Download(url, "./"); err != nil {
		t.Errorf("Download : %v", err)
	}
	fmt.Println()
	server.Close()
}

func TestRemote(t *testing.T) {
	//url := "http://ftp.osuosl.org/pub/archlinux/iso/2011.08.19/archlinux-2011.08.19-core-dual.iso"
	url := "http://samba.org/ftp/ccache/ccache-3.1.7.tar.gz"
	if err := Download(url, "./"); err != nil {
		t.Errorf("Download : %v", err)
	}
	fmt.Println()
}

func TestHttpd(t *testing.T) {
	url := "http://localhost:8080/bash-4.2.tar.gz"
	//url := "http://ftp.osuosl.org/pub/archlinux/iso/2011.08.19/randome"
	if err := Download(url, "./"); err == nil {
		t.Errorf("Download : %v", "should be nil")
	}
	fmt.Println()
}
