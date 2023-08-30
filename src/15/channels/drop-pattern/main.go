package main

/*
Package main implements drop pattern approach

Amacımız sanki çok yoğun bir ağ ortamında herkes canlı video izliyor.
Ağdaki yoğunluktan dolayı akışı bozmadan bazı paketleri drop edip hayatın
devam etmesini sağlıyoruz.
*/

import (
	"fmt"
	"time"
)

func main() {
	capacity := 5
	ch := make(chan string, capacity) // 5 string slotu olan buffered channel

	go func() {
		// bu goroutine'de receive eden tarafız, channel'a yazılanı (send) buradan
		// okuyoruz
		for v := range ch {
			fmt.Printf("(receive): %q\n", v)
		}
	}()

	packages := 20
	// küçük bir event-loop simülasyonu yapıyoruz,
	// sanki bir network içindeyiz ve tcp paketlerini okuyoruz
	for p := 0; p < packages; p++ { // 20 paket okur gibi...

		// hem send hem de receive aynı anda olmak üzere
		// select ile çoklu channel işlemleri yapabiliriz
		select {
		case ch <- fmt.Sprintf("paket %d", p): // channel'a yazıyoruz, buffer dolunca duracak
			fmt.Printf("(send): paket %d\n", p)
		default: // non-blocking, buffer dolunca burası hep çalışacak
			// buffer dolunca bloklamadan devam et
			// bu sayede;
			// network gecikme maaliyetinden kurtulduk
			// channel üzerinde baskı oluşturma maaliyetinden kurtulduk
			// bu iş bir timeout azaltması değil, kapasite azaltmasıdır.
			fmt.Printf("..(drop): paket %d\n", p) // buffer dolunca drop!
		}
	}

	close(ch) // for p := range ch burasını sonlandırır
	fmt.Println("bitiyor")

	time.Sleep(time.Second)
}
