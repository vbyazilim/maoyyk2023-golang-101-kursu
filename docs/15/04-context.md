# Bölüm 15/04: Concurrency

## Context

Goroutine’leri kullanarak başka goroutine’lere, network üzerindeki bir servise
(database, rpc) ya da başka bir backend servisine erişmeye çalıştığı zaman,
goroutine’in ulaşmaya çalıştığı şeye ulaşıp ulaşamadığını ya da belirli bir
süre sonrasında bu operasyonu iptal edilmesini sağlayan mekanizmanın adına
`context` deniyor.

Genelde **Deadline**, **Cancellation** ve diğer request kapsamındaki
sinyalları taşımak için kullanılır.

En basit tanımıyla, belli bir süre sonra goroutine’i durdurmak ya da bir http
isteğine eğer **5sn**’de (çünkü bu isteği yapan yine goroutine kullanıyor)
cevap gelemezse iptal etmemizi sağlan araç olarak anlayabiliriz.

- `WithTimeout`
- `WithCancel`
- `WithDeadline`
- `WithValue`

gibi fonksiyonları bulunur. Unutulmaması gereken şey; eğer context
kullanılacaksa **her işlemin kendi context**’i olmalı, yani, şu hatalı bir
kullanış:

```go
// kesinlikle olmaz!!!
type Foo struct {
	ctx context.Context
}
```

bu şekilde **share** edilebilir bir değer değildir. Eğer context alan bir
fonksiyon olacaksa mutlaka ilk parametre olarak `ctx` almalı ve her seferinde
sıfır bir `ctx` instance’ı verilmelidir.

```go
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}

func main(){
    duration := 150 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
    
    DoSomething(ctx, ...)
}
```

`Context` bir ağaç (tree) yapısı şeklindedir. Mutlaka **Parent / Root Context**
olur. Parent / Root Context asla **cancel** olmaz ama bundan türeyen **child**
Context’lerde bu tür cancel operasyonları yapılabilir.

Genelde parent / root olarak `context.Background()` kullanırız. Bazen
`context.TODO()` da kullanılabilir. `context.TODO()` bize `nil` olmayan boş
bir `context` döner. Uygulama içinde bir context ihtiyacı olduğunu ama nasıl
kullanacağımıza tam karar vermediğimiz durumlarda `context.TODO()` kullanırız.

---

## WithTimeout

[örnek](../../src/15/context-with-timeout)

https://go.dev/play/p/biXpvephej1

```bash
$ go run src/15/context-with-timeout/main.go 
timeout!!! context deadline exceeded
```

kod:

```go
package main

import (
	"context"
	"fmt"
	"time"
)

const timeout = 1 * time.Millisecond // 1 mili saniye

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-time.After(1 * time.Second): // time.After geriye channel döner
		fmt.Println("1 saniye sonra...")
	case <-ctx.Done():
		fmt.Println("timeout!!!", ctx.Err()) // context deadline exceeded
	}
}

// timeout!!! context deadline exceeded
```

---

## WithCancel

[örnek](../../src/15/context-with-cancel)

https://go.dev/play/p/vZ2FD2wyDjN

```bash
$ go run src/15/context-with-cancel/main.go 
```

kod:

```go
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
```

---

## WithDeadline

[örnek](../../src/15/context-with-dead-line)

https://go.dev/play/p/3_c7JD2sKJG

```bash
$ go run src/15/context-with-dead-line/main.go 
```

kod:

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := time.Now().Add(3 * 1000 * time.Millisecond) // 3sn

	ctx, cancel := context.WithDeadline(context.Background(), timeout)
	defer cancel()

LOOP:
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("1sn!")
		case <-ctx.Done():
			fmt.Println("WithDeadline", ctx.Err())
			break LOOP
		}
	}

	fmt.Println("exit")
}

// 1sn!
// 1sn!
// 1sn!
// WithDeadline context deadline exceeded
// exit
```

---

## WithValue

[örnek](../../src/15/context-with-value)

https://go.dev/play/p/wFyT4vdYsey

```bash
$ go run src/15/context-with-value/main.go 
```

kod:

```go
package main

import (
	"context"
	"fmt"
)

type ck string // custom key

func hasKey(ctx context.Context, key ck) bool {
	if v := ctx.Value(key); v != nil {
		return true
	}
	return false
}

func main() {
	idKey := ck("id")
	emailKey := ck("email")
	secretKey := ck("secret")

	// parent context
	ctx := context.Background()

	// child context
	ctx = context.WithValue(ctx, idKey, 1)

	// child context
	ctx = context.WithValue(ctx, emailKey, "vigo@foo.com")

	fmt.Println("idKey", hasKey(ctx, idKey))
	fmt.Println("emailKey", hasKey(ctx, emailKey))
	fmt.Println("secretKey", hasKey(ctx, secretKey))

	if hasKey(ctx, idKey) {
		fmt.Println("value of id", ctx.Value(idKey))
	}
	if hasKey(ctx, emailKey) {
		fmt.Println("value of email", ctx.Value(emailKey))
	}
}
```

---

## Kaynaklar

- https://go.dev/blog/context
- https://talks.golang.org/2014/gotham-context.slide#1
- https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
