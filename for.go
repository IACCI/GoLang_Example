package main

import (
	"fmt"
	"time"
)

func main() {

	

		    // for
		    for i := 1; i <= 5; i++ {
			fmt.Println("sayi : ", i)
	

	// for dongusunu while gibi kullanma

	a := 1

	for a < 10 {

		a += a
		fmt.Println(a)
		time.Sleep(500 * time.Millisecond)
	}

}
