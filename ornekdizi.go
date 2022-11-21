package main

import "fmt"

func main() {
	var haftaningunleri string
	haftaningunleri1 := "pazartesi"
	haftaningunleri2 := "sali"
	haftaningunleri3 := "carsamba"
	haftaningunleri4 := "persembe"
	haftaningunleri5 := "cuma"
	haftaningunleri6 := "cumartesi"
	haftaningunleri7 := "pazar"

	fmt.Println("gun giriniz: ")
	fmt.Scan(&haftaningunleri)

	if haftaningunleri == haftaningunleri1 {
		fmt.Println("is basi")
	} else if haftaningunleri == haftaningunleri2 {
		fmt.Println("calismaya devam")
	} else if haftaningunleri == haftaningunleri3 {
		fmt.Println("calismaya devam")
	} else if haftaningunleri == haftaningunleri4 {
		fmt.Println("calismaya devam")
	} else if haftaningunleri == haftaningunleri5 {
		fmt.Println("tatil oncesi son gun")
	} else if haftaningunleri == haftaningunleri6 {
		fmt.Println("tatil")
	} else if haftaningunleri == haftaningunleri7 {
		fmt.Println("tatil")
	}
}
