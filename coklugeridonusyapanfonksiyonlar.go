package main

import "fmt"

// ornek 1
// func main() {
// 	sayhello()
// }

// func sayhello() {
// 	fmt.Println("hello")
// }

// ornek 2
// func main() {
// 	message := "acc"
// 	sayhello(message)
// }

// func sayhello(message string) {
// 	fmt.Println(message)

// }

func main() {
	fmt.Println(add(3, 7, 8, 12, 4))
}

func aadd(x, y, z, p, e int) int {
	return x*y + z - (p / e)
}
