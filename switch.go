package main

import (
	"fmt"
)

func main() {

	/*
		acc := 4

		switch {
		case acc == 1:
			fmt.Println("bir")
		case acc == 2:
			fmt.Println("iki")
		case acc == 4:
			fmt.Println("dogru terim")
		default:
			fmt.Println("bilinmeyen sayi")
		}
	*/

	var puan int
	fmt.Println("sinav notunu gir: ")
	fmt.Scanf("%v", &puan)

	switch {
	case puan <= 0:
		fmt.Println("lutfen sayi giriniz...")
	case puan <= 49:
		fmt.Println("harf notunuz: FF", "'KALDI'")
	case puan <= 59:
		fmt.Println("harf notunuz: DD")
	case puan <= 69:
		fmt.Println("harf notunuz: CC")
	case puan <= 79:
		fmt.Println("harf notunuz: BB")
	case puan <= 89:
		fmt.Println("harf notunuz: BA")
	case puan <= 100:
		fmt.Println("harf notunuz: AA")
	default:
		fmt.Println("0 ile 100 arasinda bir not giriniz...")

	}

}
