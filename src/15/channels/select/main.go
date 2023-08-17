package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// goroutine 1
	go func() {
		for {
			ch1 <- "500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// goroutine 2
	go func() {
		for {
			time.Sleep(time.Second * 2)
			ch2 <- "\t2sn"
		}
	}()

	// sonsuz döngüde kanalları dinliyoruz.
	// çıkış için ctrl+c
	for {
		select {
		case m1 := <-ch1: // ch1'den gelirse
			fmt.Println("ch1:", m1)
		case m2 := <-ch2: // ch2'den gelirse
			fmt.Println("ch2:", m2)
		}
	}
}
