package main

import (
	"github.com/str1ngs/gurl"
	"log"
	"os"
)

func init() {
	log.SetPrefix("gurl: ")
	log.SetFlags(0)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		println("Error no urls specified")
		os.Exit(1)
	}
	err := gurl.DownloadAll("./", args)
	if err != nil {
		log.Fatal(err)
	}
}
