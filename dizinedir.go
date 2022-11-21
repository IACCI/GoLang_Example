package main

import "fmt"

func main() {

	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])

	primes := [6]int{1, 2, 3, 4, 5, 6}
	primes[0] = 1
	primes[1] = 2
	primes[2] = 3
	primes[3] = 4
	primes[4] = 5
	primes[5] = 6
	fmt.Println(primes[0], primes[1], primes[2], primes[3], primes[4], primes[5])

	fmt.Println(a)
	fmt.Println(primes)

}
