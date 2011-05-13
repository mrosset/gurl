package gurl

import (
	"testing"
)

func TestDownload(t *testing.T) {
	url := "http://localhost/gcc-go-snapshot-4.7.20110423-1-x86_64.pkg.tar.xz"
	gurl := &Client{}
	if err := gurl.Download("./", url); err != nil {
		t.Errorf("Download : %v", err)
	}
}
