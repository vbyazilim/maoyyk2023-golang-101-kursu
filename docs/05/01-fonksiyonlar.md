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

### Naked Returns ya da Named Returns

İyi bir pratik olmamakla birlikte, bazen ismilendirilmiş (named) ya da çıplak
(naked) geri dönüş değerleri kullanılabilir (return values). Bu tam olarak ne
demek? Hemen örneğe bakalım:

https://go.dev/play/p/L1CVHUT19VY

```go
package main

import "fmt"

func sum(a, b int) (result int) {
	result = a + b // buradaki result, (result int)'deki result
	return         // geri dönen şey ne? aaa pardon (result int)'deki result
}

func main() {
	fmt.Println(sum(1, 2)) // 3
}
```

Fonksiyon imzasına bakınca `func sum(a, b int) (result int)` şunu anlıyoruz;
Go, fonksiyonu işlemeye başlarken `result` diye `int` tipinde bir değişken
ataması yapacak, sonra bu fonksiyonun içinde bir yerlerde (function body)
birisi `result`’ı set edecek (değer atayacak) en sonda da o atanan değer geri
dönecek (return).

İmza esnasında atanan `result` artık isimlendirilmiş yani **named** oluyor.
Fonksiyonun sonunda neyin döndüğü bilinmeyen `return` ifadesi de **naked**
oluyor. Fonksiyonun ne döndüğünü anlamak için sürekli imzaya bakıp takip etmek
gerekiyor.

Konu ile ilgili güzel bir [makale][02].

---

## Recursivity

Türkçeye çevirmeye çalışınca **Özyineleme**, **Özyinelemeli fonksiyonlar**
gibi bir çeviri buldum internette. Kendi kendini çağırabilme durumuna
**Recursivity** deniyor. Bu tür fonksiyonlar da doğal olarak **Recursive
Functions** oluyorlar.

Go bu durumu destekliyor; en klişe örnekle devam edelim; faktöriyel hesabı:

https://go.dev/play/p/l6pS9__0mXp

```go
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) // kendini çağırdı.
}

func main() {
	fmt.Println(fact(3)) // 6
}
```


---

## Closure

Bir fonksiyon içindeki **başka** bir fonksiyonun, dışarıdaki yerel
değişkenleri de kullanması, kendi tanımlandığı kapsamın dışındaki değişkenlere
erişebilme yeteneği yani **closes over** yapması durumudur. İçerideki
fonksiyon dışarıdaki değişkenleri **referans** olarak alıp kullanır.

https://go.dev/play/p/TT0gtxq6L7U

```go
package main

import "fmt"

func fib() func() int {
	a, b := 0, 1

	// bu fonksiyon a ve b yi kullanabilir
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib()

	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}
}

// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// 55
// 89
```

Başka bir örnek;

https://go.dev/play/p/gn6O1adOQF-

```go
package main

import "fmt"

func scope() func() int {
	outer_var := 2
	foo := func() int { return outer_var } // 2
	return foo                             // foo fonksiyon olarak döndü, birinin çağırması lazım, çağırana 2 döner
}

func main() {
	// aslında "add" isimli bir fonksiyon tanımı
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(3, 4)) // 7

	sc := scope()
	// geriye fonksiyon döner,
	// bu fonksiyon da geriye int döner

	fmt.Println(sc()) // 2
}
```

### Anonim Fonksiyonlar

Adı, imzası olmayan fonksiyonlar anonim fonksiyonlardır:

```go
// anonim fonksiyon
func() {
	fmt.Println("anonymous")
}()
```

Vur-kaç (fire and forget) durumlarında, hızlıca çalıştırıp çöpe atacağımız
fonksiyonlara ihtiyaç duyduğumuzda, go routine’lerle çalışırken sıklıkla bu
tür ifadelere kullanacağız.

### Tip ya da Argüman olarak Fonksiyon

Birinci sınıf vatandaş olduğu için, fonksiyonları **type definition**
mantığında kullanabiliyoruz:

https://go.dev/play/p/tiNguGxxvNW

```go
package main

import "fmt"

type GreeterFunc func(string) string

func greet(f GreeterFunc, name string) string {
	return f(name)
}

func main() {
	func1 := func(name string) string {
		return "func1 - " + name
	}

	func2 := func(name string) string {
		return "func2 - " + name
	}

	fmt.Println(greet(func1, "vigo"))  // func1 - vigo
	fmt.Println(greet(func2, "erhan")) // func2 - erhan
}
```

`greet` fonksiyonuna `string` alıp, `string` dönen herhangi bir fonksiyonu
parametre olarak geçebiliriz. `type GreeterFunc func(string) string` artık
`GreeterFunc` diye bir type’ımız var, tipi ne? bir fonksiyon, nasıl bir
fonksiyon? `string` alıp, `string` dönen bir fonksiyon.

---

## Defer

**Defer** kelimesi; tehir etmek geciktirmek anlamındadır. Go’da da aynı
mantıkla çalışır, Fonksiyon işini bitirip geri dönmeden önce (return etmeden
önce) `defer` edilenleri çalıştırır ve çıkar.

https://go.dev/play/p/kAYIa_7Qm0U

```go
package main

import "fmt"

func greet(n string) {
	defer func() {
		fmt.Println("exit - greet") // 2. exit - greet
	}()

	fmt.Println("hello ", n)
}

func main() {
	defer func() {
		fmt.Println("exit - main") // 4. exit - main
	}()

	greet("vigo") // 1. hello  vigo

	fmt.Println("after greet") // 3. after greet
}
```

`defer` bize pek çok durumda esneklik sağlar. Bir dosya açtık, sonra otomatik
kapatmak istiyoruz:

https://go.dev/play/p/rHwN9oMOBa-

```go
package main

import (
	"log"
	"os"
)

func createTempFile() {
	f, err := os.Create("/tmp/foo.txt") // dosyayı oluştur
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // fonksiyondan çıkarken dosyayı kapat!
}

func main() {
	createTempFile()
}
```

`createTempFile` exit (return) etmeden önce file’ı kapatacaktır. `defer`
çağırılana kadar `os.Create` işlemini bekler. 

`defer` kullanırken dikkat edilmesi gereken husus şu; kapsam içindeki
değişkenlerin değerleri kopyalanır, bu da bazen yanlış sonuç almamıza neden
olur:

https://go.dev/play/p/TUsycKzYeDP

```go
package main

import "fmt"

func main() {
	a := 1
	defer fmt.Println("defer a", a) // a'nın değeri 1 kopyalandı

	a = 100             // a artık 100
	fmt.Println("a", a) // a 1000
	// eğer defer a = 100'den sonra tanımlansaydı
	// defer fmt.Println("defer a", a) // 100
}

// a 100
// defer a 1
```

Fonksiyon `return` etmeden önce `inner-scope` yani fonksiyon kapsamı içinde
bir değişkeni de bozabilir:

https://go.dev/play/p/Hpbsy_WGRT0

```go
package main

import "fmt"

func do() (a int) {
	// a -> named return
	// şu an allocate edildi, zero-value aldı: 0

	defer func() { a = 100 }() // en son çalışır ve 100 döner

	a = 1  // a'nın değeri değişti; 1 oldu
	return // naked return;
	// defer en son çalıştığı için a’yı bozar...
}

func main() {
	fmt.Println(do()) // 100
}
```

---

[01]: https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/docs/91/01-error.md
[02]: https://www.ardanlabs.com/blog/2013/10/functions-and-naked-returns-in-go.html