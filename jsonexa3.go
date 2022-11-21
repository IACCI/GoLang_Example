package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type people1 struct {
	ID   int    `json:"ID"`
	Name string `json:"isim"`
	Age  int    `json:"yas"`
}

type people2 struct {
	ID   int
	Name string
	Age  int
}

func main() {

	p1 := people1{
		ID:   123,
		Name: "ahmet can",
		Age:  23,
	}
	j, err := json.Marshal(p1)
	if err != nil {
		log.Fatal("wrong %s", err.Error())
	}
	fmt.Printf("people1 json:%s\n", string(j))

	p2 := people2{
		ID:   123,
		Name: "ahmet can",
		Age:  23,
	}
	j, err = json.Marshal(p2)
	if err != nil {
		log.Fatal("wrong %s", err.Error())
	}
	fmt.Printf("people2 json:%s\n", string(j))

}
