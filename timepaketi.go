package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("suan ki unix zamani %v\n", time.Now().Unix())
	time.Sleep(2 * time.Second)
	fmt.Printf("suan ki unix zamani %v\n", time.Now().Unix())

}
