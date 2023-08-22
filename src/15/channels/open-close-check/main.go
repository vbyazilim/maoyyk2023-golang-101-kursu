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
		// channel'dan genel golang convention'ındaki gibi
		// value, ok := şeklinde
		// ch'daki değer ve channel'ın açık/kapalı olma bilgisini alıyoruz
		msg, open := <-ch
		if !open {
			break // eğer kapalıysa döngüden çık, main artık exit etsin, iş bitsin
		}

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
