package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("dil gir:  ")
	dil, _ := reader.ReadString('\n')

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("ismini gir:  ")
	isim, _ := reader1.ReadString('\n')

	fmt.Print("yasini gir:  ")
	str, _ := reader.ReadString('\n')
	yas, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("")
	fmt.Println("girilen dil: ", dil)
	fmt.Println("girilen isim: ", isim)
	fmt.Println("girilen yas: ", yas)
	fmt.Println(" ")

}
