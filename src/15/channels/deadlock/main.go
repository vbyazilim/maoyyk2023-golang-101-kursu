package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go count(ch)

	// sonsuz döngüde ch channel'ından receive yapıyoruz..
	for {
		fmt.Println(<-ch)
	}
}

func count(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500) // yarım saniye bekletme, görmek içim
	}
}

// 0
// 1
// 2
// 3
// 4
// fatal error: all goroutines are asleep - deadlock!
//
// goroutine 1 [chan receive]:
// main.main()
// 	.../src/15/channels/deadlock/main.go:15 +0x7c
// exit status 2
