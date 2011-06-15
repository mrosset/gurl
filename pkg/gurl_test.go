package gurl

import (
	"bytes"
	"fmt"
	"http"
	"http/httptest"
	"io"
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
	start := time.Seconds()
	server := httptest.NewServer(&testHandler{})
	url := server.URL + testfile
	gurl := new(Client)
	if err := gurl.Download("./", url); err != nil {
		t.Errorf("Download : %v", err)
	}
	total := time.Seconds() - start
	fmt.Println("Finished in", total)
	server.Close()
}


func TestRemote(t *testing.T) {
	testfile := "archlinux-2010.05-netinstall-i686.iso"
	start := time.Seconds()
	//url := "http://localhost/" + testfile
	url := "http://mirrors.kernel.org/archlinux/iso/latest/" + testfile
	gurl := new(Client)
	if err := gurl.Download("./", url); err != nil {
		t.Errorf("Download : %v", err)
	}
	total := time.Seconds() - start
	fmt.Println("Finished in", total)
}
