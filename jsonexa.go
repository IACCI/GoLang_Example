package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user1 struct {
	Name     string `json:"ad"`
	Age      int    `json:"yas"`
	lastname string `json:"soyad"`
}
type user2 struct {
	Name     string
	Age      int
	lastname string
}

func main() {
	u1 := user1{
		Name:     "ahmet can",
		Age:      21,
		lastname: "ceylan",
	}
	j, err := json.Marshal(u1)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	fmt.Printf("user1 JSON: %s\n", string(j))
	u2 := user2{
		Name:     "ahmet can",
		Age:      21,
		lastname: "ceylan",
	}
	j, err = json.Marshal(u2)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	fmt.Printf("\nuser2 JSON: %s	\n", string(j))
}
