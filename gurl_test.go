package gurl

import (
	"testing"
)

func TestDownload(t *testing.T) {
	gurl := &Gurl{}
	if err := gurl.Download("/tmp/", "http://localhost/linux-2.6.38.tar.bz2"); err != nil {
		t.Errorf("Download : %v", err)
	}
}
