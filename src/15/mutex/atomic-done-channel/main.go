package main

import (
	"fmt"
	"sync/atomic"
)

var counter int64

func main() {
	done := make(chan struct{})

	fmt.Printf("[start] - %d\n", counter)

	// 10 tane goroutine
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			done <- struct{}{} // goroutine işi bitti
		}()
	}

	// 10 goroutine var, 10 kere okumamız lazım
	for i := 0; i < 10; i++ {
		<-done // biteni al
	}
	close(done)

	fmt.Printf("[end] - %d\n", counter)
}
