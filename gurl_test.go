package gurl

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strconv"
	"testing"
)

var (
	//cities = []string{"tokyo", "london", "newark", "atlanta", "dallas", "fremont"}
	cities    = []string{"fremont"}
	linFmt    = "http://%s1.linode.com/100MB-%s.bin"
	server    = httptest.NewServer(&testHandler{})
	local_url = fmt.Sprintf("%s/%s", server.URL, "foobar.tar.gz")
)

type testHandler struct {
}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	size := 1024 * 1024
	b := make([]byte, size)
	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("Content-Length", strconv.Itoa(size))
	w.Write(b)
}

func TestLocal(t *testing.T) {
	defer os.Remove(path.Base(local_url))
	if err := Download("./", local_url); err != nil {
		t.Errorf("Download : %v", err)
	}
}

func TestLocalAll(t *testing.T) {
	urls := []string{local_url, local_url}
	defer os.Remove(path.Base(local_url))
	if err := DownloadAll("./", urls); err != nil {
		t.Errorf("Download : %v", err)
	}
}

func testRemote(t *testing.T) {
	for _, c := range cities {
		if err := Download("./", fmt.Sprintf(linFmt, c, c)); err != nil {
			t.Errorf("Download : %v", err)
		}
	}
}
