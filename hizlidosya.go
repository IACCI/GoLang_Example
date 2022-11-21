package main

import (
	"io/ioutil"
	"log"
)

// hizli bir sekilde soya olusturarak icine byte yazdirir

func main() {

	err := ioutil.WriteFile("demo3.txt", []byte("budeneme"), 0666)
	if err != nil {
		log.Fatal(err)
	}

}
