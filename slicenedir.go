package main

import "fmt"

func main() {

	/*
		//basit slice ornegi
		primes := []int{14, 15, 16, 17, 18}

			var s []int = primes[1:4]

			fmt.Println(s)
	*/

	names := []string{
		"example1",
		"example2",
		"example3",
		"example4",
		"example5",
		"example6",
		"example7",
		"example8",
	}

	fmt.Println(names)

	a := names[0:]

	a[0] = "example9"
	a[1] = "example10"
	a[2] = "example11"
	a[3] = "example12"
	a[4] = "example13"
	a[5] = "example14"
	a[6] = "example15"
	a[7] = "example16"

	fmt.Println(a)

}
