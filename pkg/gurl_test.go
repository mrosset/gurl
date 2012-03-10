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

//var cities = []string{"tokyo", "london", "newark", "atlanta", "dallas", "fremont"}
var cities = []string{"fremont"}

func TestRemote(t *testing.T) {
	//cities := []string{"fremont"}
	for _, city := range cities {
		url := fmt.Sprintf("http://%s1.linode.com/100MB-%s.bin", city, city)
		if err := Download(url, "./"); err != nil {
			t.Errorf("Download : %v", err)
		}
		fmt.Println()
	}
}

func TestHttpd(t *testing.T) {
	url := "http://localhost:8080/bash-4.2.tar.gz"
	//url := "http://ftp.osuosl.org/pub/archlinux/iso/2011.08.19/randome"
	if err := Download(url, "./"); err == nil {
		t.Errorf("Download : %v", "should be nil")
	}
	fmt.Println()
}
