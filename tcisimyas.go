package main

import (
	"fmt"
	"strings"
)

func main() {

	var tc string
	fmt.Println("tc giriniz: ")
	fmt.Scan(&tc)
	if strings.Contains(tc, "11111111111") {
		fmt.Println("tc kimlik dogru verildi. ")
	} else {
		fmt.Println("tc kimlik yanlis verildi. ")
	}
	var isim string
	fmt.Println("isim giriniz: ")
	fmt.Scan(&isim)
	if strings.Contains(isim, "ahmetcanceylan") {
		fmt.Println("isim dogru girildi.")
	} else {
		fmt.Println("isim yanlis girildi.")
	}
	var yas string
	fmt.Println("yas giriniz: ")
	fmt.Scan(&yas)
	if strings.Contains(yas, "23") {
		fmt.Println("yas dogru girildi.")
	} else {
		fmt.Println("yas yanlis girildi.")
	}

	fmt.Println("")

	if tc == "11111111111" {
		fmt.Println("tc: ✓")
	} else {
		fmt.Println("tc: ×")
	}
	if isim == "ahmetcanceylan" {
		fmt.Println("isim: ✓")
	} else {
		fmt.Println("isim: ×")
	}
	if yas == "23" {
		fmt.Println("yas: ✓")
	} else {
		fmt.Println("yas: ×")
	}

}
