# Bölüm 05/01: Fonksiyonlar

Konular içinde bu noktaya gelene kadar kabaca fonksiyonları kullandık. Hem
kendimiz yazdık hem de built-in paket’lerden gelen (`fmt.Println` gibi) pek
çok fonksiyonu da çağırdık.

Belirli bir işlevi yerine getiren ve çağrıldığında belirli bir işlemi
gerçekleştiren, sıklıkla çağırana geri sonuç/sonuçlar dönen kod bloklarıdır
fonksiyonlar. Go’daki fonksiyonların karakteristik özellikleri neler?

- **first class citizen** yani fonksiyon tip olabilir, başka bir fonksiyona
  argüman olarak geçilebilir.
- Anonim olabilirler (closures)
- Slice ya da Map’in elemanı, key’i value’su olabilirler
- Bir struct’ın alanı (field’ı) olabilirler
- Channel’larda send/receive parametresi olabilirler
- Fonksiyon içinde anonim fonksiyonlar olabilir
- Sadece paket kapsamında yaşarlar (package scope)

## Signature

Fonksiyon imzası (function signature) denen şey aşağıdakiler gibidir:

```go
func Do(a string, b int) string {}           // function signature
func Done(x string, y int) (a string) {}     // function signature
```

## Argümanlar

Fonksiyonlar argümanları (by default) **pass by value** ile alırlar. Eğer
fonksiyon argümanları pointer olarak alırsa bu durumda **pass by reference**
olurlar.

Fonksiyonun aldığı ve döndüğü parametreler, go’nun **type safety**
yaklaşımından dolayı, mutlaka tanımlı tipler olmalı. Yani diğer dinamik
dillerdeki (python, ruby, javascript) gibi fonksiyon kafasına göre tipi belli
olmayan bir argüman alamaz. Son yıllarda güvenli tip tanımı ruby, python,
javascript gibi dillerede gelmeye başladı.

Otomatik olarak **pass by value** olarak giden tipler:

- sayısallar (numerics)
- bool
- array’ler
- struct’lar

**pass by reference** olanlar;

- pointer
- string’ler (immutable)
- slice’lar
- map’ler
- channel’lar

**Variadics** ile, yani **N tane** argüman geçme/alma işleri `...` ile olur:

https://go.dev/play/p/YPbLB5nstXZ

```go
package main

import "fmt"

func greet(names ...string) {
	for _, name := range names {
		fmt.Println("hello", name, "!")
	}
}

func main() {
	greet("vigo") // hello vigo !

	greet("vigo", "erhan")
	// hello vigo !
	// hello erhan !

	users := []string{"turbo", "max", "move"}
	greet(users...)
	// hello turbo !
	// hello max !
	// hello move !
}
```

`func greet(names ...string)` N tane string alır, bu `names`’e atanır, `names`
artık bir string slice yani `[]string` olur. `greet(users...)` bu durumda da
sona eklenen `...` ile verilen slice fonksiyona `greet(users[0], users[1],
users[2], ...)` gibi pas edilir.

---

## Return Values

Fonksiyon duruma göre;

- hiçbir şey dönmeye bilir.
- bir sonuç dönebilir.
- N tane sonuç dönebilir.

[Error][01] konusunda da değineceğiz ama sırası gelmişken bahsedelim, go’da
fonksiyon genelde dönmesi gereken şeyi ve hatayı döner. Hata (error) go’da
önemli bir konudur, hatta;

> Errors are values

yani error de bi değerdir, bu bakımdan da işlenmesi gerekir. **Early exit**
yaklaşımıyla, fonksiyon hata döndüğü an ya **exit** (çıkış) yapılır ya da o
hata ciddi bir şekilde değerlendirilir. Kodun akışı hemen kesilmelidir.

Bu bakımdan, go’nun kaynak koduna da baktığınızda, neredeyse tüm paket
fonksiyonları geriye sonuç + error döner:

```go
func Print(a ...any) (n int, err error)
func Printf(format string, a ...any) (n int, err error)
func Println(a ...any) (n int, err error)

// ve dahası
```

Hemen bir örnek yapalım:

https://go.dev/play/p/O-5Cuz4k0Dp

```go
package main

import (
	"fmt"
	"log"
)

type users map[string]struct{}

func greetFromMap(u users, name string) (string, error) {
	if _, ok := u[name]; !ok {
		return "", fmt.Errorf("%s not found in map", name)
	}
	return "hello " + name, nil
}

func main() {
	u := users{
		"vigo": {},
	}

	g, err := greetFromMap(u, "vigo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(g) // hello vigo

	g, err = greetFromMap(u, "lego")
	if err != nil {
		log.Fatal(err)
		// 2023/08/06 13:15:10 lego not found in map
		// exit status 1
	}

	fmt.Println(g)
}
```

---

## Recursivity

@wip

---

## Closure

@wip

---

## Defer

@wip

---

[01]: ../91/01-error.md