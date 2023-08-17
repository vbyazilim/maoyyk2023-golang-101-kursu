package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go count(ch)

	// ch channel'ı kapanana kadar içinde iterasyon yapıyoruz
	// bu yöntemle kanal açık mı? kapalı mı? bakmaya gerek kalmıyor...
	// bu bir syntactic sugar
	for msg := range ch {
		fmt.Println(msg)
	}
}

func count(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500) // yarım saniye bekletme, görmek içim
	}
	close(ch)
}
