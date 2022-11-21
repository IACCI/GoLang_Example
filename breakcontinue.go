package main

import "fmt"

func main() {

	//break
	/*
		i := 0
		for {
			if i >= 3 {
				break
			}
			fmt.Println("i'nin degeri: ", i)
			i++
		}
	*/

	//continue

	for i := 0; i < 7; i++ {
		if i == 3 { // ifte belirtilen sayida continue verdiyse belirtilen sayi ciktida gozukmez!
			continue
		} else if i == 6 {
			break
		}
		fmt.Println("cikti: ", i)
	}
	fmt.Println("islem devam ediyor...")

}
