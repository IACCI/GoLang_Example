// package main

// import "fmt"

// func main() {

//		defer fmt.Println("world")
//		fmt.Println("hello")
//	}
package main

import "fmt"

var isconnected bool = false

func main() {
	fmt.Printf("connection open: %v\n", isconnected)
	databaseprocessing()
	fmt.Printf("connection open: %v\n", isconnected)
}

func databaseprocessing() {
	connect()
	fmt.Println("deferring disconnect")
	defer disconnect()
	fmt.Printf("connection open: %v\n", isconnected)
	fmt.Println("doing something")
}
func connect() {
	isconnected = true
	fmt.Println("connected to database!")
}
func disconnect() {
	isconnected = false
	fmt.Println("disconnected!")
}
