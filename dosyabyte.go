package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteslices := []byte("bu dosyaya yaz!\n")
	byteswritten, err := file.Write(byteslices)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WROTE %d bytes\n ", byteswritten)

}
