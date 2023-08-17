package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	done := make(chan struct{})

	// bu kısım main'den farklı bir kulvarda çalışmaya başladı
	go func() {
		fmt.Println(<-ch) // gönder gelsin
		done <- struct{}{}
	}()

	ch <- "hello world" // gönderdim
	<-done

	fmt.Println("bitti")
}
