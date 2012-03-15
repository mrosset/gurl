package main

import (
	"fmt"
	"gurl/pkg"
	"log"
	"os"
)

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
