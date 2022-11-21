package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	a1 := Address{"sdfgsgsg", "sgsdgsgs", 3623572}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "lkjljl", city: "jkljljk", Pincode: 277001}

	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "ahmet can ceylan", city: "duzce", Pincode: 81000}
	fmt.Println("Address3: ", a3)
}
