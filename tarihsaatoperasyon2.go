package main

import (
	"fmt"
	"time"
)

func main() {

	x := fmt.Println

	xtime := time.Date(1923, 4, 23, 21, 44, 0, 0, time.UTC)

	now := time.Now()

	x(now)

	x("---------------------")

	/*
		x(now.Year())
		x(now.Month())
		x(now.Day())
		x(now.Weekday())
		x(now.Hour())
		x(now.Minute())
		x(now.Second())
		x(now.Location())
	*/

	x("---------------------")

	// tarih karsilastirma

	x(xtime.Before(now))
	x(xtime.After(now))
	x(xtime.Equal(now))

	x(now.Sub(xtime))

}
