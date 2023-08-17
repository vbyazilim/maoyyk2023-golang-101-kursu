package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	amount := 10

	fmt.Println("channel'a", amount, "adet sayı yollayacağız")

	ch := make(chan int)
	done := make(chan struct{})

	// channel'a gönder (send)
	go func() {
		for i := 0; i < amount; i++ {
			randomInt, _ := rand.Int(rand.Reader, big.NewInt(100))
			ch <- int(randomInt.Int64()) + 1 // randon sayı 0-100 araası
		}
	}()

	// channel'dan al (receive)
	go func() {
		fmt.Println()
		for i := 0; i < amount; i++ {
			fmt.Println("gelen (received) sayı", <-ch)
		}
		fmt.Println()
		close(ch)

		close(done) // channel'ı kapat ya da
		// done <- struct{}{} // done için kullanılan channel'a boş struct koy, channel'a veri gitti
	}()

	<-done // veri gelene kadar blokla...

	fmt.Println("bitti...")
}
