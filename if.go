package main

import "fmt"

func main() {

	// if, else if ile cogaltilabilir, her ihtimale karsi bir sonuc icin else kullanilir.
	a := 20
	b := 20

	if b > a {
		fmt.Println("buyuktur")
	} else if b == a {
		fmt.Println("esittir")
	} else {
		fmt.Println("kucuktur")
	}

	/*

		     if icerisinde de degisken tanimlanir fakat if dongusunun disinda, tanimlanan degisken kullanilamaz.

			if a := 2; a == 1 {
				fmt.Println("dogru")
			} else {
				fmt.Println("yanlis")
			}
	*/
}
