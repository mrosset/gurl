package main

import (
	"os"
	"path"
	"testing"
)

func TestMain(t *testing.T) {
	url := "http://mirrors.us.kernel.org/gnu/bash/bash-4.2.tar.gz"
	defer os.Remove(path.Base(url))
	os.Args = []string{url, url}
	main()
}
