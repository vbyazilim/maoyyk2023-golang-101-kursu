# Bölüm 15/02: Concurrency

## Channels

`chan` (channel) go’da goroutine’ler arasında veri taşımaya yarayan ön-tanımlı
(built-in) bir tiptir. Yani `int` gibi `map` gibi bir tiptir. Bir goroutine
channel’a veri yazarken, başka bir goroutine o channel’ı dinleyip okuyabilir.

Asıl amaç channel’ların goroutine’ler arasında paylaşılabilmesidir. Hatta
sevgili Rob Pike’ın o efsane sözlerinden biri olan;

> Don't communicate by sharing memory, share memory by communicating

hatırlatmak isterim; **hafızayı paylaşarak** iletişim kurmayın, 
**iletişim kurarak** hafızayı paylaşın! Yani channel’ları kullanın! diyor...

Normal şartlarda fonksiyon geriye değer dönebilir ama `go` anahtar kelimesiyle
tetiklenen fonksiyonların dönüşünü `main` beklemediği için, goroutine ile
çalışan fonksiyon ile `main` arasındaki iletişim channel’lar üzerinden
sağlanır.

- Birden fazla goroutine aynı channel’a yazabilir
- Birden fazla goroutine aynı channel’dan okuyabilir

Tabi tüm bunlar sizin uygulamak istediğiniz concurrency pattern’ine bağlı
olarak değişen şeylerdir.

- Channel’daki veri mutlaka bir tipte olmalı, yani `make(chan)` ile channel
  oluşturulmalı, aynı `map`’deki gibi...
- Default olarak channel read/write (okuma/yazma) senkron işlemlerdir
- Channel’a bir değer atamak, aynı fonksiyona parametre geçmek gibidir
- Channel’a reference types (pointer, map, slice vs...) geçerken dikkatli olmak
  gerekir. Çünkü DATA RACE olması yüksek ihtimaldir. Yani bir goroutine değer okurken,
  diğer bir goroutine’de yeni değer atamaya çalışırsa durum race-condition durumu olur.

Bir goroutine, channel’a yazdığı zaman, diğer bir goroutine o channel’dan
okuyana kadar pause / block olur. Yani goroutine channel’dan okumak istedi ve
okuyacağı bir değer (value) yoksa, değer gelene kadar o goroutine beklemeye
devam eder...

Aslında bu aynı linux/unix pipe’lara benzer:

```bash
$ cat foo | grep 'bar' | cut -d'.' -F1
#               |           |
#               ↓           |
#          cat’i bekler     |      
#                           ↓
#                         grep’i bekler 
```

Channel’ı kapatmak, sizin artık o channel ile işinizin kalmadığı anlamına
gelir.

4 Tür channel var;

1. Unbuffered
1. Buffered
1. Closed
1. Nil

3 Tür channel state var;

1. **nil** -> zero-value
1. **open** -> `make(...)`
1. **close** -> `close(...)`

**State’ine göre send/receive uygunluğu**

|         | example      | nil     | open    | closed  |
|:--------|:-------------|:--------|:--------|:--------|
| send    | `ch <- true` | blocked | allowed | panic!  |
| receive | `<-ch`       | blocked | allowed | allowed |

**Garantileme durumuna göre**

|           | Garanti    | Garanti Yok   | Gecikmeli Garanti |
|:----------|:-----------|:--------------|:------------------|
| `channel` | Unbuffered | Buffered `ch > 1` | Buffered `ch = 1` |

Normalde channel’ı fonksiyona geçerken **bi-directional** yani hem **send**
hem de **receive** edebilir şekilde geçiyoruz;

```go
func Foo(ch chan int){}
```

Bunu limitlemek mümkün. Yani duruma göre **receive-only** ya da **send-only**
yapmak da mümkün;

```go
func receiveOnly(ch <-chan int) {
	fmt.Println("read/receive", <-ch)
}

func sendOnly(ch chan<- int) {
	ch <- 1
}
```

[Örnek](../../src/15/channels/send-only)

```go
package main

import (
	"fmt"
)

func returnReceiveOnly() <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)

		// bu fonksiyondan dönen channel'ı receive
		// eden her kimse, en fazla 100 tane sayı
		// receive edebilir.

		// loop bitiminde defer ile channel kapandığı
		// için, 100+ zero-value => 0 gelir...
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()

	return c
}

func main() {
	r := returnReceiveOnly() // returns receive-only channel

	// r <- 10 // invalid operation: cannot send to receive-only type <-chan int

	for i := 0; i < 200; i++ {
		fmt.Println(<-r)
	}
}

// read/write-only channels are distinct types, the compiler can use its existing
// type-checking mechanisms to ensure the caller does not try to write stuff
// into a channel it has no business writing to.
```

Bi-directional channel ile teknik/performans farkı yok sadece compile-time’da;
**receive only** kullanmak istediğiniz channel’a send etmek isterseniz;

```bash
invalid operation: cannot send to receive-only type <-chan int
```

gibi hata vererek kodu düzeltmenizi ister. Şimdi ilk yaptığımız, hani
`time.Sleep` kullanarak yaptığımız örneği **channel** kullanarak yapalım;

[Örnek](../../src/15/channels/basic-goroutine-with-channel)

https://go.dev/play/p/E8EzcHasNxE

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool) // channel holds bool type!

	go func() {
		fmt.Println("hello from goroutine!")
		ch <- true // writing to ch channel
	}()

	<-ch // channel'dan veri gelene kadar blok!, reading from ch channel
	fmt.Println("exit!")
}

// hello from goroutine!
// exit!
```

![Goroutine With Channel](diagrams/goroutine-with-channel.gif)

---

## Unbuffered Channels

@wip
