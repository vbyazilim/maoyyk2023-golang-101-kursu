# Bölüm 03/01: Dil Kuralları

## Unicode

Go’nun varsayılan karakter kodlaması (encoding) `UTF-8`. Bu ne anlama geliyor?
Siz kod yazarken değişken/sabit ya da fonksiyon adı olarak **TÜRKÇE** karakterler
bile kullanabilirsiniz:

https://go.dev/play/p/ulIuTtHW0tf

```go
package main

import "fmt"

func main() {
	kullanıcıAdı := "vigo"
	fmt.Println(kullanıcıAdı) // vigo
}
```

Bu sadece **proof-of-concept** yani "evet gerçekten de çalışıyor"u göstermek
için yazılmış bir kod. Anadiliniz ne olursa olsun genel yazılım prensiplerini
takip edip, genelin anlayabileceği türden deklarasyonlar yapmanızı tavsiye
ediyoruz.

Unicode kapsamındaki tüm karakterler ve `_` karakteri go için bir **KARAKTER**
tanımı. `0`’dan `9`’a kadar sayılar, **Octal** yani 8’lik sayı sistemi için
`0`’dan `7`’ye sayılar, **Hexadecimal** yani 16’lık sayı sistemi için `0`’dan
`9`’a kadar sayılar ve `A`’dan `F`’e kadar harfler kullanılıyor.

---

## Anahtar Kelimeler

Toplamda **25** tane anahtar kelime bulunur:

    break        default      func         interface    select
    case         defer        go           map          struct
    chan         else         goto         package      switch
    const        fallthrough  if           range        type
    continue     for          import       return       var

Bunlara ek olarak;

**Öntanımlı sabitler ailesi**  

- `true`
- `false`
- `iota`
- `nil`


**Sayısal (numeric types) tipler ailesi**  

- `uint8` : (0-255)
- `uint16` : (0-65535)
- `uint32` : 32-bit’lik pozitif tamsayılar (0 ile 4294967295 arası)
- `uint64` : 64-bit’lik pozitif tamsayılar (0 ile 18446744073709551615 arası)
- `int8` : 8-bit’lik (signed) tamsayılar (-128 ile 127 arası)
- `int16` : 16-bit’lik (signed) tamsayılar (-32768 ile 32767 arası)
- `int32` : 32-bit’lik (signed) tamsayılar (-2147483648 ile 2147483647 arası)
- `int64` : 64-bit’lik (signed) tamsayılar (-9223372036854775808 ile 9223372036854775807 arası)
- `float32` : [IEEE-754][01] uyumlu 32-bit’lik ondalık sayılar
- `float64` : [IEEE-754][01] uyumlu 64-bit’lik ondalık sayılar
- `complex64` : float32 gerçek ve sanal kısımlara sahip tüm karmaşık sayıların kümesi
- `complex128` : float64 gerçek ve sanal kısımlara sahip tüm karmaşık sayıların kümesi
- `byte` : `uint8` için takma isim (alias)
- `rune` : `int32` için takma isim (alias)
- `any` : `interface{}` (empty interface) için takma isim (alias)

**Gömülü gelen fonksiyonlar**  

- `append`
- `cap`
- `close`
- `complex`
- `copy`
- `delete`, 
- `imag`
- `len`
- `make`
- `new`
- `panic`
- `print`
- `println`
- `real`
- `recover`

**Diğer**  

- `string`
- `error`

---

## Operatörler ve İşaretçiler

@wip

---

## Built-in Veri Tipleri

@wip

---

## Kod Stili

@wip

---

[01]: https://en.wikipedia.org/wiki/IEEE_754