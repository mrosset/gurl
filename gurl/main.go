package main

import (
	"gurl/pkg"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		println("Error no url specified")
		os.Exit(1)
	}
	url := os.Args[1]
	if err := gurl.Download(url,"./"); err != nil {
		println(err)
	}
}
