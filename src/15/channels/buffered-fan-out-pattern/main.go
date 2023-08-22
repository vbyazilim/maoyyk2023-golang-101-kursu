package main

import (
	"fmt"
)

func main() {
	maxGoroutines := 20 // 20 goroutine kullanacağız.

	ch := make(chan int, maxGoroutines) // buffered channel, goroutine sayısı kadar kapasite
	done := make(chan struct{})         // done channel, sinyalizasyon için

	for g := 0; g < maxGoroutines; g++ {
		// nolint her goroutine'nin kendi bufferı var, paralel olarak çalışıyor.
		go func(n int) {
			ch <- n
			fmt.Println("ch <- kanala yolluyoruz (send)", n)
		}(g)
	}
	close(done)

	// tüketiyoruz, gelenleri alıyoruz.
	for maxGoroutines > 0 {
		fmt.Println("<- kanaldan alıyoruz (receive)", <-ch)
		maxGoroutines--
	}

	<-done
}
