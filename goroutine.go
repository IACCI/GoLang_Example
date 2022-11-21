package main

import "time"

func main() {

	// goroutine bir thread gorevi gorur yani islemi ayirir ve farkili bir goroutine'de calistirir.
	go xfunc()
	time.Sleep(1000 * time.Millisecond)

}

func xfunc() {

	for l := byte('a'); l <= byte('z'); l++ {
		println(string(l))
	}
}
