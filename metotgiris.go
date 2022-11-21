package main

import "fmt"

/*
type sayilar struct {
 	x int
 	y int
 }

 func main() {

	fmt.Println(sayilar{2, 3})

 }

*/

type sayilar struct {
	x int
	y int
}

func main() {

	s := sayilar{1, 2}

	s.x = 4
	fmt.Println(s.x)

}
