package main

import (
	"fmt"
	"strconv"
)

func main() {

	// CONVERT

	// 1. ORNEK
	var mystring = "2"
	var x = 10
	var z = 5

	fmt.Println(mystring, x, z)

	// STRING'DEN INT'E DONUSTURMEDE STRCONV.ATOI KULLANILIR.

	sayi, _ := strconv.Atoi(mystring)
	fmt.Println(sayi)

	toplama := sayi + 5
	fmt.Println(toplama)

	fmt.Println("===============================")

	// 2. ORNEK
	var myint = 5
	var c = 34
	var d = 12

	fmt.Println(myint, c, d)

	//INT'DEN STRING'E DONUSTURMEDE STRCONV.ITOA KULLANILIR.

	acc := strconv.Itoa(myint)
	fmt.Println("string: ", acc)
	toplam := myint + c + d
	fmt.Println(toplam)

}
