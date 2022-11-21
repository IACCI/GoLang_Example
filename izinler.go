package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("dizi.go", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("hata : yazma izni reddedildi. ")
		}
	}
	file.Close()

}
