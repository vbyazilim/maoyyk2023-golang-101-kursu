package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	var wg sync.WaitGroup

	fmt.Printf("[start] - %d\n", counter)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("[end] - %d\n", counter)
}
