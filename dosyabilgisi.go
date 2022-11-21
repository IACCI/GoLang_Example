package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {

	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file name : ", fileInfo.Name())
	fmt.Println("size of bytes : ", fileInfo.Size())
	fmt.Println("permission : ", fileInfo.Mode())
	fmt.Println("last modified : ", fileInfo.ModTime())
	fmt.Println("is dictionary : ", fileInfo.IsDir())
	fmt.Println("system interface type : ", fileInfo.Sys())
	fmt.Printf("system info : %+v\n\n ", fileInfo.Sys())

}
