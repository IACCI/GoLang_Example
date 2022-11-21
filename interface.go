package main

import (
	"fmt"
	"strconv"
)

// interface

type carface interface {
	run() bool
	stop() bool
	information() string
}

func carexecute(c carface) {
	fmt.Printf("\n")
	fmt.Println("arac bilgisi : \n" + c.information())
	fmt.Printf("\n")

	msg := ""

	isrun := c.run()

	if isrun {
		msg = "calisiyor"
	} else {
		msg = "calismiyor"
	}
	fmt.Println("arac " + msg + ".")

	isstop := c.stop()

	if isstop {
		msg = "durdu"
	} else {
		msg = "durmadi"
	}
	fmt.Println("arac " + msg + ".")

}

func main() {
	ferr := new(ferrari)
	ferr.brand = "ferrari"
	ferr.model = "f50"
	ferr.color = "red"
	ferr.speed = 340
	ferr.price = 1.4
	ferr.special = true
	//fmt.Println(ferr.information())

	lamb := new(lamborghini)
	lamb.brand = "lamborghini"
	lamb.model = "specialclass"
	lamb.color = "black"
	lamb.speed = 290
	lamb.price = 0.4
	//fmt.Println(lamb.information())

	// carexecute
	carexecute(ferr)
	carexecute(lamb)

}

// base structs

type car struct {
	brand string
	model string
	color string
	speed int
	price float64
}

type specialproduction struct {
	special bool
}

// ferrari

type ferrari struct {
	car // composition yani car structindaki ozellikleri aliyor
	specialproduction
}

func (_ ferrari) run() bool {
	return true
}

func (_ ferrari) stop() bool {
	return true
}

func (x *ferrari) information() string {

	ret := "\t" + "model : " + x.brand + " " + x.model + "\n" + "\t" + "color : " + x.color + "\n" + "\t" + "speed : " + strconv.Itoa(x.speed) + "\n" + "\t" + "price : " + strconv.FormatFloat(x.price, 'g', -1, 64) + " million"
	add := "yes"

	if x.special {
		ret += "\n" + "\t" + "special : " + add
	}
	return ret

}

type lamborghini struct {
	car // composition yani car structindaki ozellikleri aliyor
	specialproduction
}

func (_ lamborghini) run() bool {
	return true
}

func (_ lamborghini) stop() bool {
	return false
}
func (x *lamborghini) information() string {

	ret := "\t" + "model : " + x.brand + " " + x.model + "\n" + "\t" + "color : " + x.color + "\n" + "\t" + "speed : " + strconv.Itoa(x.speed) + "\n" + "\t" + "price : " + strconv.FormatFloat(x.price, 'g', -1, 64) + " million"
	add := "no"

	if x.special {
		ret += "\n" + "\t" + "special : " + add
	}
	return ret

}
