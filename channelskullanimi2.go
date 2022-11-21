package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(8)
	ch := make(chan string)
	go zfunc(ch)
	go printer(ch)
	time.Sleep(100 * time.Millisecond)
}

func zfunc(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
}

func printer(ch chan string) {

	for {

		println(<-ch)
	}

}
