package main

import "fmt"

func main() {

	

	numbers := make([]int, 5)
	numbers[0] = 123
	numbers[1] = 345
	numbers[2] = 563
	numbers[3] = 34
	numbers[4] = 453
	fmt.Println(numbers)

	numbers = append(numbers, 222)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

}
