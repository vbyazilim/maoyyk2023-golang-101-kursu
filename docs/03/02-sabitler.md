# Bölüm 03/02: Dil Kuralları

## Sabitler

Sabit, adından da anlaşılacağı gibi değeri bir kez atanan ve değişmeyen /
değişemeyen demektir.

Sabit olarak tanımlanabilir tipler;

- `bool`
- `rune` (*aslında `int32` için takma ad*)
- `int` familyası
- `float` familyası
- `complex` familyası
- `string`

Sayısal familyadan tanımlanan sabitler **Numeric Constants** (Sayısal
Sabitler) denir. Sabiti tanımlarken `const` anahtar kelimesi kullanılır,
takiben işaretçisi yani **identifier**’ı, tipi ve son olarak değeri atanır ve
bu tür tanıma **Typed Constant** denir;

```go
// const <IDENTIFIER_NAME> = <VALUE>         // Untyped Constant
// const <IDENTIFIER_NAME> <TYPE> = <VALUE>  // Typed constant


const domain string = "example.com"  // "domain" identifier, "string" tipi
const pi float32 = 3.14
```

Tanım esnasında tipi belirtilmeyen sabitlere **Untyped Constant** denir. Bu tür
durumlarda go **Type Inference** yani tipi/türü tahmin etme, tip çıkarımı işlemi
yaparak bunu anlar;

```go
const a = 5          // untyped integer constant
const b = "vigo"     // untyped string constant
const pi = 3.14      // untyped floating-point constant
const foo = '1'      // 49  - untyped rune constant
const world = "世界"  // untyped unicode string constant
```

`len` ile **Iterable** yani içinde gezilebilen, yinelenebilir tiplerin boyunu
(length) alırız. `len("hello")` dediğimizde `"hello"` aslında go için:
**Untyped String Constant** olarak işlenir. `print(1 > 2)` ifadesinde `1` ve `2`
go için **Untyped Integer Constant** durumundadır.

---

## `iota`

Sadece sabitler için geçerli olup, “küçük parça” anlamındadır. Orijini Yunan
alfabesindeki [9. karakterin][01] adından gelir. Asıl amacı belli bir matıkta
artan/azalan/değişen sabit değerler üretmektir ve başlangıç değeri `0`’dır:

https://go.dev/play/p/0qOhdiyAr6V

```go
package main

import "fmt"

type Num int

const (
	Sıfır Num = iota * 2 // 0’la başla 2’şer arttır
	İki
	Dört
	Altı
	Sekiz
)

func main() {
	fmt.Println("Sıfır", Sıfır)
	fmt.Println("İki", İki)
	fmt.Println("Dört", Dört)
	fmt.Println("Altı", Altı)
	fmt.Println("Sekiz", Sekiz)
}
// Sıfır 0
// İki 2
// Dört 4
// Altı 6
// Sekiz 8
```


[01]: https://en.wikipedia.org/wiki/Iota