package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2022, 8, 23, 16, 0, 0, 0, time.UTC)
	fmt.Printf("calisiyor: %s\n", t)

	fmt.Println("****************************")

	now := time.Now()
	fmt.Printf("mevcut zaman: %s\n", now)

	fmt.Println("****************************")

	fmt.Println("yil: ", now.Year())
	fmt.Println("ay: ", now.Month())
	fmt.Println("gun: ", now.Day())
	fmt.Println("gun ismi: ", now.Weekday())
	fmt.Println("saat: ", now.Hour(), ":", now.Minute(), ":", now.Second())

	fmt.Println("****************************")

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("tomorrow is %v,%v,%v,%v\n", tomorrow.Day(), tomorrow.Month(), tomorrow.Year(), tomorrow.Weekday())
}
