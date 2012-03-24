package main

import (
	"os"
	"path"
	"testing"
)

func TestMain(t *testing.T) {
	url := "http://mirrors.us.kernel.org/gnu/bash/bash-4.2.tar.gz.sig"
	defer os.Remove(path.Base(url))
	urls := []string{url, url}
	os.Args = append(os.Args, urls...)
	main()
}
