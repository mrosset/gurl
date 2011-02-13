package gurl

import "testing"


func TestDownload(t *testing.T) {
	//Download("http://curl.haxx.se/download/curl-7.21.3.tar.bz2", "./")
	Download("http://jupiter/curl-7.21.3.tar.gz", "./")
}

func TestVersion(t *testing.T) {
	if Version() != string("libcurl/7.21.3 OpenSSL/1.0.0d zlib/1.2.5") {
		t.Fail()
	}
}
