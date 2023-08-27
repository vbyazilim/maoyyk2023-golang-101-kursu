# Bölüm 03/01: Dil Kuralları

## Encoding Nedir?

**Encoding**, metin karakterlerini bilgisayarların anlayabileceği sayısal
değerlere çevirme işlemidir. Bu sayısal değerler, metinleri depolamak, iletmek
ve işlemek için kullanılır.

**UTF-8** (Unicode Transformation Format - 8bit), Unicode karakter setini
temsil etmek için kullanılan bir metin kodlama standardıdır. UTF-8, farklı
karakterleri değişen uzunluklarda bayt dizileri olarak temsil eder.

Temel olarak, İngilizce harfleri ve sembolleri 1 byte’la temsil ederken, diğer
dillerdeki karakterleri daha fazla byte’la ifade eder. Bu özelliği sayesinde
hem İngilizce metinleri hem de farklı dillerdeki karakterleri aynı kodlama
altında destekler. Bu nedenle, günümüzde genellikle tercih edilen bir metin
kodlama standardıdır.

---

## Unicode Desteği

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
- `uint` : üzerinde çalıştığı mimariye göre 32 ya da 64bit’lik pozitif tamsayılar
- `int` : üzerinde çalıştığı mimariye göre 32 ya da 64bit’lik (signed) tamsayılar
- `uintptr` : işaretçi (pointer) değerinin yorumlanmamış bitlerini saklamak için yeterince büyük pozitif tamsayılar


**Gömülü gelen fonksiyonlar**  

- `append`
- `cap`
- `close`
- `complex`
- `copy`
- `delete`
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

## Operatörler

    +    &     +=    &=     &&    ==    !=    (    )
    -    |     -=    |=     ||    <     <=    [    ]
    *    ^     *=    ^=     <-    >     >=    {    }
    /    <<    /=    <<=    ++    =     :=    ,    ;
    %    >>    %=    >>=    --    !     ...   .    :
         &^          &^=          ~

Aritmetik işlemler, mantık işlemleri, büyük/küçük kontrolleri, tek sefer
(unary) işlemleri, bit kaydırma, kısa değişken tanımlama, diziler ve kesitler
olmak üzere bir kısım operatör karakterleri bulunur.

## İşaretçiler (Identifiers)

Değişken, sabit ya da tip tanımlamarı yaptığımız şeylere **identifier**
diyoruz. Yani `a := 5` dediğimizde `a` bir işaretçi yani identifier oluyor.
Go’daki geçerli (valid) işaretçi tanımlamarına bazı örnekler;

```go
a 
_x1
BuExportEdilebilir
uğur
```

Sayısal tanımlamalarda da;

```go
42        // 10’luk sayı
0600      // 8’lik sayı
0xFF      // 16’lık sayı
0.        // Kesirli ondalıklı
1.2       // Kesirli ondalıklı
072.40    // == 72.40
1.e+0
011i      // == 11i
170141183460469231731687303715884105727 // çılgın
```

şeklinde kullanılabiliyor. Unicode yani 32-bit’lik karakterler için `rune`
kullanıyoruz, bu tür ifade şekline **Rune Literal** (rune kalıbı) deniyor:

```go
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
```

---

## Built-in Veri Tipleri

Standart kütüphane bir kısım hazır veri tipi ile birlikte geliyor, kabaca;

- **Strings** : Metinsel tipler
- **Booleans** : `true` / `false` mantıksal veri tipleri
- **Numerics** : `int` / `float` ve `complex` familyası
- **Composite (Unnamed) Types** (Bileşik İsimsiz Tipler) : Array, Slice, Struct, Map

---

## Kod Stili

2 tür yorum (comment) yazma stili var;

1. **Line Comment** : `// bu bir yorum satırı` şeklinde
1. **General Comment** : `/* bu bir yorum satırı */` şeklinde

Line delimeter yani kod ifadesi satırları (code statements) `C`, `JavaScript`
ya da `PHP` dilindeki gibi `;` ile bitmiyor, go bunu compile time (derleme
anında) kendisi ekliyor. `;` sadece **inner-scope** yani sadece iç kapsam
durumlarında kullanılıyor;

```go
// short-if declaration - kısa if bildirimi - inner-scope
// v değişkeni sadece {} içinde yaşar
if v := math.Pow(x, n); v < lim {
	return v
}
print(v) // error

// i sadece {} içinde yaşar
for i := 5; i< 9; i++  {
  fmt.Println(i)
}
print(i) // error
```


---

[01]: https://en.wikipedia.org/wiki/IEEE_754