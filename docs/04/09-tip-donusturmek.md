# Bölüm 04/09: Veri Tipleri

## Tip Dönüştürmek

**Type Conversion** yani tipleri birbirine dönüştürmek. Aynı familyadan
türeyen tipleri çevirmek daha kolayken, alakasız tipleri arasında çevirmek
bazen mümkün değil. `int`’i `float64`’e çevirmek kolay;

```go
var i int = 42
var f float64 = float64(i)
```

Dikkat ettiyseniz aslında tip işaretçisi (type identifier) olan `float64` aynı
bir fonksiyon gibi çağırılabiliyor. Örnekteki `f`’i `uint`’e de çevirsek;

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// aynı işi bu şekilde de yapabiliriz:
i := 42
f := float64(i)
u := uint(f)
```

Peki, elimizde sayısal bir değer var bunu metinsel (string) değere çevirmek
istiyoruz:

https://go.dev/play/p/gZ3Xjujw8UN

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 5

	var s string

	s = strconv.Itoa(i)          // Integer to Ascii
	fmt.Printf("%v, %[1]T\n", s) // 5, string

	// aslında strconv.Itoa bir kısa yol
	s = strconv.FormatInt(int64(i), 10)
	fmt.Printf("%v, %[1]T\n", s) // 5, string
}
```

Peki string’i sayıya nasıl çevireceğiz?

https://go.dev/play/p/-1r2gsSTf1i

```go
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	s := "5"
	var i int

	n, err := strconv.Atoi(s) // Ascii to integer
	if err != nil {
		log.Fatal(err) // error'ü ekrana yaz va os.Exit(1)
	}

	i = n
	fmt.Printf("%v, %[1]T\n", i) // 5, int
	fmt.Printf("%d, %[1]T\n", i) // 5, int

	// aslında strconv.Atoi bir kısa yol

	ii, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		log.Fatal(err) // error'ü ekrana yaz va os.Exit(1)
	}
	i = int(ii)
	fmt.Printf("%v, %[1]T\n", i) // 5, int
	fmt.Printf("%d, %[1]T\n", i) // 5, int
}
```

## Type Alias

Kendi oluşturduğumuz tipler aslında varolan (built-in gelen) tiplere bir tür
kısa yol. Kod içinde uygun tip kontrolleri yapmak, daha az kod yazmak için
kullanılır:

https://go.dev/play/p/pkVgdra5tQQ

```go
package main

import "fmt"

// myString'in altında yatan tip built-in string
type myString string

func greet(s myString) {
	fmt.Println("greet:", s)
}

func main() {
	s := myString("hello")
	ss := "hello"

	fmt.Printf("%s, %[1]T\n", s)  // hello, main.myString
	fmt.Printf("%s, %[1]T\n", ss) // hello, string

	// fmt.Println(s == ss)
	// error
	// invalid operation: s == ss (mismatched types myString and string)

	greet(s) // greet: hello

	// greet(ss)
	// error
	// annot use ss (variable of type string) as myString value in argument to greet

	// ss'i yani düz string'i myString tipine çevirdik
	greet(myString(ss)) // greet: hello

	fmt.Println(string(s) == ss)   // true
	fmt.Println(s == myString(ss)) // true
}
```
