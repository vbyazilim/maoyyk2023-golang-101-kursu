# Bölüm 04/07: Veri Tipleri

## Struct Methods ve Receivers

Şimdi daireyi ifade eden bir `struct`’ımız olsun, alanları da `x` ve `y`
koordinatları ile yarıçapı `r` olsun;

```go
type Circle struct {
	x, y, r float64
}
```

Dairenin alanını bulmak için `𝞹 * r * r` yapmamız gerekiyor. Bunun için bir
fonksiyon yazacağız, argüman olarak `Circle` tipinde bir girdi alacak, bize
sonucu dönecek;

https://go.dev/play/p/tYc_W5HX2ei

```go
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func area(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	circle := Circle{0, 0, 5}
	result := area(circle)

	fmt.Printf("%v\n", result) // 78.53981633974483
}
```

Halbuki bu işlemi, `Circle` struct’ına bir method olarak takabilseydik?

https://go.dev/play/p/LKuNk71sl6v

```go
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	circle := Circle{0, 0, 5}
	result := circle.area()

	fmt.Printf("%v\n", result) // 78.53981633974483
}
```

`area()`, `Circle` struct’ına ait metot (method) oluyor. `(c Circle)` ise
alıcı (receiver) oluyor. `Circle`’ın kendisine `c` üzerinden erişiyoruz.
`area()` sadece içeriden veri okuduğu için, yani **read-only** olduğu için
**value receiver** oluyor.

Eğer `Circle`’ın alanlarının değerlerinde değişiklik yapmak gerekseydi,
örneğin yarı çapı 2 katına çıkaran bir iş gerekseydi, o zaman yazma işlemi de
yapmamız gerekecekti, bu durumda **pointer receiver** kullanmamız gerekiyor:

https://go.dev/play/p/YKFDt2wJnJX

```go
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) rDoubler() {
	c.r *= 2
}

func main() {
	circle := Circle{0, 0, 5}
	result := circle.area()
	fmt.Printf("%v\n", result) // 78.53981633974483

	circle.rDoubler()
	fmt.Println(circle.area()) // 314.1592653589793
}
```

- Eğer struct’a ait alanlarda (field/property) değişiklik yapacaksak **pointer receiver**
- Sadece değer okuyup işlem yapacaksak **value receiever**
- Eğerki sadece bir metot pointer receiver bile olacaksa tüm metotları pointer receiver olarak
  tanımlamak

iyi pratiklerdendir. İyi bir go struct metotlarında tek tip receiver olur.
