# Bölüm 15/03: Concurrency

## Mutex

Mutex, **mutual exclusion** yani verilen ortak bir kararla dışlama işlemi
anlamındadır. `sync` paketindeki `Mutex` type’ı da bu tür durumlardaki
senkronizasyonu sağlar.

Örneğin hafızada bir sayı var. **10 tane goroutine** ateşleyerek bu sayıyı
arttırıyoruz. Peki o esnada okumak istesek ne olacak? Bazı goroutine’ler
değeri değiştirirken, bazıları da okumaya çalışacak ve bu esnada **DATA RACE**
oluşacak!

Şimdi test konusunda işlediğimiz [örneğe](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-datarace) geri
dönelim ve DATA RACE’i çözelim:

```bash
$ go test -v -race github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/15/mutex/kvstore
=== RUN   TestDataRace
--- PASS: TestDataRace (0.00s)
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/15/mutex/kvstore	1.637s
```

kod:

```go
package kvstore

import (
	"errors"
	"sync"
)

var errKeyNotFound = errors.New("key not found")

// Store is key-value store!
type Store struct {
	mu sync.RWMutex
	db map[string]string
}

// Set new key to store.
func (s *Store) Set(k, v string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.db[k] = v
	return nil
}

// Get accepts key, returns value and error.
func (s *Store) Get(k string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.db[k]
	if !ok {
		return "", errKeyNotFound
	}
	return v, nil
}

// New returns new Store instance.
func New(db map[string]string) Store {
	return Store{db: db}
}
```

`Store` içinde `sync.RWMutex` gömdük (embed, composition). Read ve Write
işlemlerinde kullanacağımız için `RWMutex` kullandık. Eğer sadece okuma
yapsak; `Mutex` kullanmak yeterli olurdu. Okuma ve yazma işlemlerinden önce
lock ederek bir tür değeri sahipleniyoruz ve o an biz unlock edene kadar kimse
işlem yapamıyor. İş bitiminde kilidi açıyoruz ve akış devam ediyor.

Günün sonunda biz bu işi `map`’ten okuma, `map`’e yazma için kullanıyoruz,
go’da bu iş için hazır bir tip var; `sync.Map`. **Concurrent Safe Map** yani
eş zamanlı işlerde güvenle kullanabileceğimiz bir `map`. `map`’in tipi:
`map[string]any`

İki özel durum için optimize edilmiştir:

1. key’in değeri sadece bir kez yazıldığında ama çok kez okunduğunda
   cache’leme yapar 
2. Birden fazla goroutine okuyabilir, yazabilir ve varolan key’in değeri üzerine
   değişiklik yapabilir

Bu tür kullanımlar olduğunda performans olarak `Mutex` ve `RWMutex`’e göre
lock etme işlerinde **gözle görülür** derecede performanslı çalışır.

[örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/15/mutex/sync-map)

https://go.dev/play/p/k974sMo66ZD

```bash
$ go run -race src/15/mutex/sync-map/main.go   # DATA RACE varsa çıksın! -race
```

kod:

```go
package main

import (
	"fmt"
	"sync"
)

var (
	m  sync.Map
	wg sync.WaitGroup
)

func main() {
	// 10 tane goroutine kullanarak key:i, value: i
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}

	wg.Wait() // goroutine'lerin işini bitirmesini bekle

	m.Store("foo", "bar") // manual olarak key ekle

	// value, ok syntactic sugar, eklediğin key'i oku

	if v, ok := m.Load("foo"); ok {
		fmt.Println("foo ->", v)
	}

	// goroutine'lerle için doldurduğun map'ten değerleri geri oku.
	for i := 0; i < 10; i++ {
		if v, ok := m.Load(i); ok {
			fmt.Printf("%d -> %v\n", i, v)
		}
	}

	fmt.Println("bitti")
}
```

---

## Channel Kendi içinde Mutex Kullanır

Hemen örneğe bakalım; basit bir webserver. Her istek geldiğinde hafızadaki
değeri **1** arttırıyor (sanki??)!

[örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/15/mutex/in-channel)

```bash
$ go run -race src/15/mutex/in-channel/main.go   # DATA RACE varsa çıksın! -race

# şimdi ayrı bir shell session açıp:
$ hey "http://localhost:9000"   # 200 tane get isteği atacak.
$ http "http://localhost:9000"  # bakalım counter kaç oldu?
```

kod:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = make(chan int)

func main() {
	go generator()

	http.HandleFunc("/", handler)

	fmt.Println("listening on :9000")
	log.Fatal(http.ListenAndServe(":9000", nil)) // nolint
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.String())
	fmt.Fprintf(w, "number %d", <-counter)
}

func generator() {
	for i := 0; ; i++ {
		counter <- i
	}
}
```

---

## `sync/atomic`

Bazı durumlarda **mutex** işini **cpu instruction**’larını kullanarak da
çözebiliriz. Örneğimizde toplamda **10 goroutine** ile 0’dan-100’e kadar
sayarak, `counter` değerini atomic olarak arttırıyoruz:

[örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/15/mutex/atomic-waitgroup)

https://go.dev/play/p/6AtJqNlGx18

```bash
$ go run -race src/15/mutex/atomic-waitgroup/main.go   # DATA RACE varsa çıksın! -race
```

kod:

```go
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
```

Aynı işi **done channel pattern**’i kullanarak yapalım:

[örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/15/mutex/atomic-done-channel)

https://go.dev/play/p/G_ZM6by6Dph

```bash
$ go run -race src/15/mutex/atomic-done-channel/main.go   # DATA RACE varsa çıksın! -race
```

kod:

```go
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
```

---

## Kaynaklar

- https://gobyexample.com/mutexes
- https://yourbasic.org/golang/mutex-explained/
