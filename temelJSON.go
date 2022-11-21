package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonStr := `
{
	"data": {
		"object":"card",
		"id": "card_111111111111",
		"first_name":"ahmetcan",
		"last_name":"ceylan",
		"balance":"54.950"
	}
}
`
	// map modunda cikti
	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}
	fmt.Println(m)

	fmt.Println("-------------")

	// json seklinde cikti
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
