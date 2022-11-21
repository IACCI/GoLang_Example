package main

import "fmt"

func main() {

	
		    // dizi yazdirma ornek

			myarray1 := [3]int{}
			myarray1[0] = 23
			myarray1[1] = 34
			myarray1[2] = 12
			fmt.Println(myarray1[0], myarray1[1], myarray1[2])
			fmt.Println(myarray1)
	

	
		       // yazdirma ve len func ile eleman sayisini yazdirma

			   var sayilar = [5]int{}
			   sayilar[0] = 12
			   sayilar[1] = 23
			   sayilar[2] = 34
			   sayilar[3] = 45
			   sayilar[4] = 67
			   fmt.Println(sayilar)
			   fmt.Println("eleman sayisi : ", len(sayilar))
	

	// int dizi icine ... koyarak size belirtilmez fakat {} icinde elemanlari yazariz
	
		myarray := [...]int{78, 54, 321, 98, 7, 321, 34}
		fmt.Println(myarray)
		fmt.Println("eleman sayisi : ", len(myarray))
	

	
		// index 0'dan basladigi icin ve len eleman sayisini gosterdigi icin bir fazla gosterecekti oy yuzden -1 yapildi.
		var cars [3]string
		cars[0] = "SUV"
		cars[1] = "SEDAN"
		cars[2] = "HATCHBACK"

		index := 0
		for index <= len(cars)-1 {
			fmt.Println(cars[index])
			index++
	

	var cars [3]string
	cars[0] = "SUV"
	cars[1] = "SEDAN"
	cars[2] = "HATCHBACK"

	for i, value := range cars {
		fmt.Println(i, "index : ", value)
	}

}

}