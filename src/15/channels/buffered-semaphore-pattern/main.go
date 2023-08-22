package main

import (
	"fmt"
	"time"
)

func main() {
	maxGoroutines := 20
	ch := make(chan int, maxGoroutines) // 20 kapasitesi olan buffered channel

	maxSemaphore := 5
	sch := make(chan struct{}, maxSemaphore) // semafore channel'ı. kapasitesi 5

	// done := make(chan struct{})

	for g := 0; g < maxGoroutines; g++ {
		go func(n int) {
			sch <- struct{}{}       // kapasite dolana kadar blok yok (5 slot)
			time.Sleep(time.Second) // sadece görüntülemek amacıyla bu goroutine'i beklet

			ch <- n // kapasite dolana kadar blok yok (20 slot)
			<-sch   // 5 olunca blokla
		}(g)
	}

	// channel'dan okuyoruz, kuyruğu tüketiyoruz...
	for maxGoroutines > 0 {
		fmt.Println(maxGoroutines, "ch", <-ch)
		maxGoroutines--
	}

	fmt.Println("bitti...")
}
