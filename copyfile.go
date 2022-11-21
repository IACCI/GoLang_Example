package main

import (
	"io"
	"log"
	"os"
)

func main() {

	//
	originfile, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originfile.Close()

	// yeni bir dosya olustur
	newfile, err := os.Create("demo2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newfile.Close()

	// byte'leri kaynaktan hedefe kopyala
	byteswritten, err := io.Copy(newfile, originfile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("copied %d bytes", byteswritten)

	//dosya icerigini isle
	//bellegi diske bosalt
	err = newfile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
