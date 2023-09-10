# Bölüm 10/01: `nil`

## `nil`

`nil` aslında ön tanımlı bir anahtar kelimde değil (25 anahtar kelimeden biri
değil, `break`, `func` vs...) Kendine ait özel bir değeri var. Aslında `nil`
şu tipler için **uninitialized** değeridir:

- Pointer
- Interface
- Map
- Slice
- Channel
- Function

Eğer **identifier** (değişken, sabit gibi) değeri `nil` ise `nil` o zaman
varolur. `nil`’in tipi yoktur, go ile beraber gelen ön tanımlı **bir şey**’dir.

---

## Nerelerde Kullanılır

- pointer alan metot, fonksiyon’da argüman geçilmiş mi?
- struct `nil` mi?
- slice `nil` mi?
- map `nil` mi?
- error `nil` mi?
- channel `nil` mi?

Ne zaman `nil` olunur?

- pointer eğer hiçbir yere point etmiyorsa
- slice’ın altında bir Array yoksa (underlying array)
- map, channel ya da fonksiyon henüz **initialize** edilmemişse (uninitialized ise)
- Interface’ler; eğer boş durumdaysalar ya da **zero-value** durumundaysalar
  ya da herhangi bir tipi karşılamıyorlarsa (milyon dolarlık cevap)

### `nil` Pointer vs Pointer

```go
var s *string        // (nil) pointer, not allocated or initialized! points to a string value
var ss = new(string) // pointer but initialized, memory allocated, has a zero-value
```

### `nil` as Value Receiver

https://go.dev/play/p/2NwTWYjudLz

```go
package main

import "fmt"

// Numeros is a custom type
type Numeros struct {
	nums []int
}

func sum(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	return sum(n[1:]...) + n[0]
}

func (n *Numeros) Sum() int {
	if n == nil {
		return 0
	}
	return sum(n.nums...)
}

func main() {
	n1 := Numeros{nums: []int{1, 2, 3}}
	fmt.Println(n1.Sum()) // 6

	n2 := Numeros{}
	fmt.Println(n2.Sum()) // 0
}
```
