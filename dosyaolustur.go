package main

import (
	"log"
	"os"
)

var (
	newfile *os.File
	errr    error
)

func main() {

	newfile, err = os.Create("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

}
