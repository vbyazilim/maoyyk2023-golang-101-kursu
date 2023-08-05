# Bölüm 04/03: Veri Tipleri

## Numerics

Gündelik hayatta kullandığımız tam sayılar yani **Integers** ve ondalıklı
sayılar yani **Floats** olmak üzere iki ana grup mevcut. Sayının kaç bitlik
alan kapladığını sonundaki ekten anlayabiliriz. `int8` 8-bit, `int64` 64-bit
anlamındadır.

Başına eklenen `u` o sayının **unsigned integer** yani sadece pozitif tam sayı
olabileceğini söyler bize.

| Tip          | Açıklama                                                                                                 |
|:-------------|:---------------------------------------------------------------------------------------------------------|
| `int8`       | `-128` ile `127` arasında değer taşır.                                                                   |
| `int16`      | `-32768` ile `32767` arasında değer taşır.                                                               |
| `int32`      | `-2147483648` ile `2147483647` arasında değer taşır.                                                     |
| `int64`      | `-9223372036854775808` ile `9223372036854775807` arasında değer taşır.                                   |
| `uint8`      | `0` ile `255` arasında değer taşır.                                                                      |
| `uint16`     | `0` ile `65535` arasında değer taşır.                                                                    |
| `uint32`     | `0` ile `4294967295` arasında değer taşır.                                                               |
| `uint64`     | `0` ile `18446744073709551615` arasında değer taşır.                                                     |
| `float32`    | 32-bit ondalık sayılar, `-3.4E+38` ile `+3.4E+38` arası                                                  |
| `float64`    | 64-bit ondalık sayılar, `-1.7E+308` ile `+1.7E+308` arası                                                |
| `complex64`  | `float32` tipinde gerçek sayı ve hayali sayı: `1.0 + 7i`                                                 |
| `complex128` | `float64` tipinde gerçek sayı ve hayali sayı: `1.0 + 7i`                                                 |
| `byte`       | `uint8` için takma ad                                                                                    |
| `rune`       | Karater ifade etmek için `int32`’ye takma ad                                                             |
| `int`        | En az 32-bit’lik (*64-bit de olabilir*) negatif/pozitif sayı ifade etmek için. Dikkat! bu `int32` değil! |
| `uint`       | En az 32-bit’lik (*64-bit de olabilir*) pozitif sayı ifade etmek için. Dikkat! bu `uint32` değil!        |
| `uintptr`    | Hafıza adres işaretçilerini saklamak için (*memory address pointers*)                                    |

`complex64` ve `complex128` tipleri aslında tümleşik gelen `complex` fonksiyonu
ile bu tür sayıları üretir. Bu fonksiyonun imzasına baktığımızda;

```go
func complex(r, i FloatType) ComplexType
```

2 tane `r` ve `i` değişkenine atanmış `FloatType` tipinde girdi alıp geriye
`ComplexType` döner:

https://go.dev/play/p/K7WFE2TwbYt

```go
package main

import "fmt"

func main() {
	c1 := complex(5, 7) // ister fonksiyon ile
	c2 := 1 + 3i        // ister direkt yazarak

	fmt.Printf("%v\n", c1) // (5+7i)
	fmt.Printf("%v\n", c2) // (1+3i)
}
```

Go, tanımlanan değişkenlerin tipi konusunda çok katıdır. Yani sayısal olduğunu
düşündüğünüz iki değeri kafanıza göre işleyemezsiniz. `int` tipindeki bir sayı
ile `float32` tipindeki sayıyı toplamak için tip dönüştürmesi (Type
Conversion) yapmak gerekiyor ve bu iki farklı türün toplamının hangi tipte
sonuç vermesi gerekiyorsa o türden işlem yapmak gerekiyor:

https://go.dev/play/p/RiUrmoOWbwA

```go
package main

import "fmt"

func main() {
	a := 32

	// argument reuse tekniği!
	fmt.Printf("a: %v (%[1]T)\n", a) // a: 32 (int)

	b := 1.1
	fmt.Printf("b: %v (%[1]T)\n", b) // b: 1.1 (float64)

	sum1 := a + int(b)
	fmt.Printf("sum1: %v (%[1]T)\n", sum1) // sum1: 33 (int)

	sum2 := float64(a) + b
	fmt.Printf("sum2: %v (%[1]T)\n", sum2) // sum2: 33.1 (float64)
}
```
