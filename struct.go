package main

import (
	"fmt"
	"strconv"
)

type humann struct {
	firstname string
	lastname  string
	age       int
}

func newhuman() *humann {
	h := new(humann)
	return h
}

func newhumanwithparams(firstname, lastname string, age int) *humann {

	h := new(humann)
	h.firstname = firstname
	h.lastname = lastname
	h.age = age
	return h

}

func main() {
	// x := newhuman()
	// x.age = 23
	// fmt.Println(x.age)
	x := newhumanwithparams("acc", "acc", 23)
	//fmt.Println(x.firstname)
	var data = x.firstname + " " + x.lastname + " / " + strconv.Itoa(x.age)
	fmt.Println(data)
}
