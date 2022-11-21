package main

import "fmt"

/*
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	for i, v := range pow {

		fmt.Printf("2**%d = %d\n", i, v)

	}
*/

func main() {

	/*
		// array ornegi
		a := [...]string{"a", "b", "c", "d"}
			for i := range a {

				fmt.Println("dizinin", i, ". indisi", a[i])

			}
	*/

	//map ornegi

	baskentler := map[string]string{"turkiye": "ankara", "japonya": "tokyo", "azerbaycan": "baku"}
	for key, val := range baskentler {

		fmt.Println(key, " ulkesinin baskenti", val)

	}

}
