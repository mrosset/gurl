package main

import (
	"gurl/pkg"
	"log"
	"os"
)

func init() {
	log.SetPrefix("gurl: ")
	log.SetFlags(0)
}

func main() {
	if len(os.Args) <= 1 {
		println("Error no urls specified")
		os.Exit(1)
	}
	err := gurl.DownloadAll("./", os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
