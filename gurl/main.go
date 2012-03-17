package main

import (
	"fmt"
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
		println("Error no url specified")
		os.Exit(1)
	}
	url := os.Args[1]
	err := gurl.Download(url, "./")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
