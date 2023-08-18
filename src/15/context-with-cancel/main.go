package main

import (
	"context"
	"fmt"
)

func main() {
	// burada başlayan goroutine "leak" etmeden "return" ediyor...
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1

		go func() {
			for {
				select {
				case <-ctx.Done():
					return // <- leak etmeden return...
				case dst <- n:
					n++
				}
			}
		}()

		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// aşağıdaki loop bitince artık main func'dan exit etme işlemi başlayacak
	// ve defer cancel() çalışacak!
	// cancel() bitince ctx.Done()'dan receive edilecek ve goroutine'den
	// çıkılacak.
	for n := range gen(ctx) {
		fmt.Println(n)

		if n == 5 {
			break
		}
	}
	// code buraya geldiğinde defer devreye girip cancel'ı tetikleyecek.
	// goroutine'daki ctx.Done()'a sinyal gelecek ve goroutine güvenli
	// bir şekilde işini bitirip return edecek.
}
