package gurl

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
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
	start := time.Now()
	server := httptest.NewServer(&testHandler{})
	url := server.URL + testfile
	gurl := new(Client)
	if err := gurl.Download(url, "./"); err != nil {
		t.Errorf("Download : %v", err)
	}
	total := time.Now().Sub(start)
	fmt.Println("Finished in", total)
	server.Close()
}

func TestRemote(t *testing.T) {
	start := time.Now()
	//url := "http://localhost/" + testfile
	url := "http://ftp.osuosl.org/pub/archlinux/iso/2011.08.19/archlinux-2011.08.19-core-dual.iso"
	gurl := new(Client)
	if err := gurl.Download(url, "./"); err != nil {
		t.Errorf("Download : %v", err)
	}
	total := time.Now().Sub(start)
	fmt.Println("Finished in", total)
}
