package main

import (
	"log"
	"os"
)

// dosya kapama acma islemleri

func main() {

	file, err := os.OpenFile("demo.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

}

/*file, err := os.Open("demo,txt")
defer file.Close()

if err != nil {
	log.Fatal(err)

}
*/
