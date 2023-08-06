# Bölüm 04/08: Veri Tipleri

https://go.dev/blog/maps

## Map

`key=value` çiftleri şeklinde içinde veri tuttuğumuz, key ekleme, çıkartma ve
silme yapabildiğimiz bir veri tipidir. Python’daki `dict` ya da Ruby’deki
`Hash` ya da Php’deki `Associative Array` gibi düşünülebilir. Key üzerinden
value’ya ulaşmak `O(1)` karmaşıklığındadır (complexity) yani çok hızlıdır.

- [Hash table][01] algoritmasını [kullanır][02].
- Ekleme, çıkartma (silme), okuma `O(1)` karmaşıklığındadır. Sadece ekleme
  işlemi [amortized][03] karmaşıklık algoritması kullanır.
- Key/Value çiftleri sıralanmadan (un-ordered) şekilde tutulur
- Key’ler eşsizdir (unique), bir map’te aynı key’den sadece bir tane olur
- `make` fonksiyonu ile **map literal** (map kalıbı) üretilir, hafızaya
  yerleştirilir, initialize olur
- map’lerin initialize değeri `nil` olur, `nil` olan map’e **key eklenemez**!
- `len` bize key/value çiftinin uzunluğunu verir
- map `nil` mi diye bakılabilir
- map’leri birbiriyle kıyaslamak için `reflection.DeepEqual` fonksiyonu kullanılır.

https://go.dev/play/p/X7zOWsYTuyA

```go
package main

import "fmt"

var m map[string]int // nil map, key’i string, value’su int...

func main() {
	fmt.Println(m, len(m)) // map[] 0
	// m["foo"] = 5 // panic: assignment to entry in nil map

	m = make(map[string]int)
	m["foo"] = 5
	fmt.Println(m, len(m)) // map[foo:5] 1
}
```

`map` alan fonksiyona `nil` geçebiliriz, key var mı? yok mu ? bakabiliriz:

https://go.dev/play/p/difVbLWsdCr

```go
package main

import "fmt"

type myMap map[string]string // key: string, value: string

func printMap(m myMap) {
	fmt.Printf("%+v\n", m)
}

func main() {
	printMap(nil) // map[]

	m := myMap{
		"username": "vigo",
	}

	printMap(m)                // map[username:vigo]
	fmt.Println(m["username"]) // vigo
	fmt.Println(m["foo"])      //

	val, ok := m["foo"]
	fmt.Println("ok", ok)   // ok false
	fmt.Println("val", val) // val
}
```

Ekleme, çıkarma ve hatalı işlemler örneği:

https://go.dev/play/p/YhkzsWMBGBI

```go
package main

import "fmt"

func main() {
	m1 := map[string]int{
		"ocak":  1,
		"şubat": 2,
	}

	var m2 map[string]int
	m2 = make(map[string]int)
	m2["ocak"] = 1
	m2["şubat"] = 2

	fmt.Println(m1) // map[ocak:1 şubat:2]
	fmt.Println(m2) // map[ocak:1 şubat:2]

	// fmt.Println(m1 == m2)
	// error:
	// invalid operation: m1 == m2 (map can only be compared to nil)

	m1["mart"] = 3
	m2["mart"] = 3

	fmt.Println(m1) // map[mart:3 ocak:1 şubat:2]
	fmt.Println(m2) // map[mart:3 ocak:1 şubat:2]

	delete(m1, "mart") // mart key'ini sil
	fmt.Println(m1)    // map[ocak:1 şubat:2]

	for k, v := range m2 {
		fmt.Println("key", k, "->", v)
	}
	// key ocak -> 1
	// key şubat -> 2
	// key mart -> 3

	// m1["mart"] = "ok"
	// error
	// cannot use "ok" (untyped string constant) as int value in assignment

	// m1[1] = "ocak"
	// error
	// cannot use 1 (untyped int constant) as string value in map index
	// cannot use "ocak" (untyped string constant) as int value in assignment
}
```

Aynı Array ve Slice’daki gibi kapasite kavramı `map` içinde var;

https://go.dev/play/p/7azH9ymAbvS

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	m1 := make(map[string]int, 10) // Preallocate, 10 tane yer ayır
	fmt.Printf("%#v\n", m1)        // map[string]int{}

	fmt.Println(len(m1))           // 0
	fmt.Println(unsafe.Sizeof(m1)) // 8 byte

	m2 := make(map[string]int)
	fmt.Println(unsafe.Sizeof(m2)) // 8 byte
}
```

Karşılaştırma;

https://go.dev/play/p/_XIw23G6bPq

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	m1 := make(map[string]int, 10) // Preallocate, 10 tane yer ayır
	m2 := make(map[string]int)

	m1["foo"] = 1
	m2["foo"] = 1

	fmt.Println(m1 == nil) // false
	fmt.Println(m2 == nil) // false

	fmt.Println(reflect.DeepEqual(m1, m2)) // true
}
```

`map` otomatik olarak referans tipindedir (by ref) aynı pointer ve slice’lar
gibi:

https://go.dev/play/p/VZz6a3Sn_QE

```go
package main

import "fmt"

type myMap map[string]string

func modifyMap(m myMap) {
	m["foo"] = "modified"
}

func main() {
	var m myMap
	m = make(myMap)
	m["foo"] = "bar"

	fmt.Printf("initial: %v, memory: %[1]p\n", m)
	// initial: map[foo:bar], memory: 0x14000074180

	fmt.Println("foo:", m["foo"])
	// foo: bar
	modifyMap(m)

	fmt.Printf("modified: %v, memory: %[1]p\n", m)
	// modified: map[foo:modified], memory: 0x14000074180

	fmt.Println("foo:", m["foo"])
	// foo: modified
}
```

Bazen `map`’in sadece `key` kısmı bize lazım olur, `value` ile işimiz yoktur.
Bu durumda value yerine öyle bir şey koymalıyız ki `0 byte` yer kaplasın? Bu
durumda `empty struct` tam da aradığımız şeydir:

https://go.dev/play/p/EUmDBlchY5b

```go
package main

import "fmt"

func main() {
	// elimizde 100_000 tane isim var
	// acaba dışarıdan gelen isim bizimde var mı?
	m := map[string]struct{}{
		"uğur":  {},
		"erhan": {},
		"turbo": {},
		"vigo":  {},
	}

	fmt.Println(m)
	// map[erhan:{} turbo:{} uğur:{} vigo:{}]

	// uğur var mı?
	if _, ok := m["uğur"]; ok {
		fmt.Println("uğur tanıdığımız biri")
		// uğur tanıdığımız biri
	}
}
```

[01]: https://en.wikipedia.org/wiki/Hash_table
[02]: https://yourbasic.org/algorithms/hash-tables-explained/
[03]: https://yourbasic.org/algorithms/amortized-time-complexity-analysis/