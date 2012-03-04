package gurl

import (
	"bytes"
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
	server.Close()
}

func TestRemote(t *testing.T) {
	//url := "http://ftp.osuosl.org/pub/archlinux/iso/2011.08.19/archlinux-2011.08.19-core-dual.iso"
	//url := "http://localhost:8080/src/bash-4.2.tar.gz"
	url := "http://localhost:8080/src/Mac.OSX.Lion.10.7.2.dmg"
	if err := Download(url, "./"); err != nil {
		t.Errorf("Download : %v", err)
	}
}

func testRemoteNoneExists(t *testing.T) {
	url := "http://localhost:8080/src/bash-4.2.tar.gz"
	//url := "http://ftp.osuosl.org/pub/archlinux/iso/2011.08.19/randome"
	if err := Download(url, "./"); err == nil {
		t.Errorf("Download : %v", "should be nil")
	}
}
