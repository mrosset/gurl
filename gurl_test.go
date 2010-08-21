package gurl

import "testing"


func TestDownload(t *testing.T) {
    Download("http://curl.haxx.se/download/curl-7.21.1.tar.bz2")
}

func TestVersion(t *testing.T) {
	if Version() != string("libcurl/7.21.1 OpenSSL/1.0.0a zlib/1.2.5") {
		t.Fail()
	}
}
