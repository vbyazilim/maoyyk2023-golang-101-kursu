# Bölüm 11/01: Generics

Generic’ler, programlamada aynı kod yapısını farklı veri türleriyle
kullanmamıza olanak tanıyan bir yaklaşımdır.

Bu, kodun yeniden kullanılabilirliğini artırırken, tür güvenliğini de korur.
Örneğin, `array` veya `slice` gibi veri yapılarını farklı veri tipleri için de
kullanabiliriz, böylece kod tekrarını önleriz ve hataları azaltırız.

Genel olarak, generic’ler programlama dillerinde daha esnek ve verimli kod
yazmamıza olanak sağlar.

Go programlama dilinde, başlangıçta generic’ler desteklenmiyordu. Ancak `go
1.18` sürümüyle birlikte, **generic programlamayı** destekleyen bir özellik
olan **type parameters** (tür parametreleri) tanıtıldı. Bu, go dilinde
**generic** kod yazmayı mümkün kıldı.

---

## Fonksiyonlarda Generic’ler

Go’da generic fonksiyonlar, fonksiyonun parametrelerinde ve dönüş değerlerinde
tür parametreleri kullanılarak tanımlanır.

Öncelikle **non-generic** bir fonksiyonu ele alalım:

```go
func sum(a int, b int) int {
    return a + b
}
```

Bu şekilde tanımlanan fonksiyon, sadece `int` türü için çalışır. Bu fonksiyonu
`float64` türü için kullanmak istediğimizde, aşağıdaki gibi bir hata alırız:

```go
func main() {
    fmt.Println(sum(1.2, 2.3))
}
// cannot use 1.2 (type float64) as type int in argument to add
```

Bunun için `float64` türü için ayrı bir fonksiyon tanımlamamız gerekir:

```go
func sumFloat(a float64, b float64) float64 {
    return a + b
}
```

Bu, kod tekrarına neden olur ve hata ayıklamayı zorlaştırır. Bu sorunu çözmek
için, fonksiyonu **generic** olarak tanımlayabiliriz:

```go
//       T type parameter, can be int or float64
func sum[T int | float64](a T, b T) T {
    return a + b
}
```

Bu şekilde tanımlanan fonksiyon, `int` veya `float64` türü için çalışır. 

https://go.dev/play/p/aEoTFQBln_Q

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/generics-functions)

```go
package main

import "fmt"

func sum[T int | float64](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(1, 2))     // 3
	fmt.Println(sum(1.2, 2.3)) // 3.5
}
```

Ama bu şekilde tanımlanan fonksiyonlarda `T` için bütün tipleri eklememiz
gerekir. Bu da fonksiyonun okunabilirliğini azaltır.

Örnek olarak, fonksiyon şu hale gelir:

```go
func sum[T int | int8 | int16 | float32 | float64](a T, b T) T {
    return a + b
}
```

Bu sorunu çözmek için bir `interface` kullanabiliriz:

https://go.dev/play/p/dMFFc30TJH_t

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/generics-functions-interface)

```go
package main

import "fmt"

type number interface {
	int | int8 | int16 | float32 | float64
}

func sum[T number](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(int8(10), int8(2)))   // 12
	fmt.Println(sum(int16(10), int16(2))) // 12
	fmt.Println(sum(1, 2))                // 3
	fmt.Println(sum(1.2, 2.3))            // 3.5
}
```

Fakat bu şekilde bir tanımalama yaptığımızda da bütün tipleri interface
içerisine eklememiz gerekir. Bunun yerine `constraints.Ordered`
`interface`’ini kullanabiliriz, öncelikle paketi projeye ekleyelim:

```bash
$ go get -u golang.org/x/exp/constraints
```

sonra;

https://go.dev/play/p/a-AQTFUfZ3T

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/generics-constraints)

```go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func sum[T constraints.Ordered](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(int8(10), int8(2)))   // 12
	fmt.Println(sum(int16(10), int16(2))) // 12
	fmt.Println(sum(1, 2))                // 3
	fmt.Println(sum(1.2, 2.3))            // 3.5
}
```

`constraints.Ordered` gerekli tüm `Integer | Float | ~string` tipleri içerir.

---

## Custom tiplerde Generic’ler

Şimdi başa dönüp tekrar add fonksiyonunu ele alalım. Bu fonksiyonu, custom bir
tipte kullanmak istediğimizi varsayalım:

```go
type SchoolNumber int

func sum[T int](a T, b T) T {
    return a + b
}
```

Bu fonksiyonu `SchoolNumber` türü için kullanmak istediğimizde, aşağıdaki gibi
bir hata alırız:

```go
func main() {
    fmt.Println(add(SchoolNumber(1), SchoolNumber(2)))
}
// SchoolNumber does not satisfy int (possibly missing ~ for int in int)
```

Bu hatanın nedeni, `sum` fonksiyonunun parametrelerinin `int` türü için
tanımlanmış olmasıdır. Bu nedenle, `SchoolNumber` türü için kullanamayız.

Bunu çözmek için `sum` fonksiyonunu aşağıdaki gibi tanımlayabiliriz:

```go
func sum[T ~int](a T, b T) T {
	return a + b
}
```

Tilda (~) işareti, `T`’nin `int` türü veya `int` türünden bir alt tür olması
gerektiğini belirtir. `SchoolNumber` türü, `int` türünden bir alt tür olduğu için,
bu fonksiyonu `SchoolNumber` türü için kullanabiliriz:

https://go.dev/play/p/R2t6n6xx1EW

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/generics-custom-types)

```go
package main

import "fmt"

// SchoolNumber is a custom type definition uses int.
type SchoolNumber int

func sum[T ~int](a T, b T) T {
	return a + b
}

func main() {
	n1 := SchoolNumber(1)
	n2 := SchoolNumber(2)
	fmt.Println(sum(n1, n2))
}
```

---

## Generic Fonksiyon Çağrıları

Şimdi generic fonksiyon çağrılarını ele alalım. Generic fonksiyonları çağırmak
için, fonksiyonun parametrelerindeki **tür parametrelerini** belirtmemiz
gerekir.

Öncelikle **generic olmayan** bir fonksiyonu ele alalım, fonksiyon `[]int`
alıyor ve bu slice’ı işleyecek fonksiyonu alıyor;

    input: [1, 2, 3, 4, 5]
    fonksiyon: (n) => n * 2
    
    çıktı: [2, 4, 6, 8, 10]  olmalı


```go
func numMutator(values numbers, fn mapperFunc) numbers {
	result := make(numbers, len(values))
	for i, v := range values {
		result[i] = fn(v)
	}
	return result
}
```

Bu fonksiyonu aşağıdaki gibi çağırabiliriz:

https://go.dev/play/p/4BJEzwJ9s5E

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/generics-function-calls)

```go
package main

import "fmt"

type (
	mapperFunc func(int) int
	numbers    []int
)

func numMutator(values numbers, fn mapperFunc) numbers {
	result := make(numbers, len(values))
	for i, v := range values {
		result[i] = fn(v)
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	fn := func(n int) int {
		return n * 2
	}

	fmt.Println(numMutator(input, fn))
}
```

Bu fonksiyon sadece `int` türü için çalışır. Yine generic fonksiyonlar
kullanarak bu fonksiyonu farklı türler için kullanabiliriz:

https://go.dev/play/p/iE4yCG6yjy8

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/generics-function-calls-and-types)

```go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	mapperFunc[T any] func(T) T
	numbers[T any]    []T
)

func numMutator[T constraints.Ordered](values numbers[T], fn mapperFunc[T]) []T {
	result := make([]T, len(values))
	for i, v := range values {
		result[i] = fn(v)
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	fn := func(n int) int {
		return n * 2
	}

	fmt.Println(numMutator(input, fn))
}
```

ya da `float64` türü için:

https://go.dev/play/p/kH2Q2ilzqwA

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/generics-function-calls-and-types2)

```go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	mapperFunc[T any] func(T) T
	numbers[T any]    []T
)

func numMutator[T constraints.Ordered](values numbers[T], fn mapperFunc[T]) []T {
	result := make([]T, len(values))
	for i, v := range values {
		result[i] = fn(v)
	}
	return result
}

func main() {
	input := []float64{1.2, 2.3, 3.4, 4.5, 5.6}
	fn := func(n float64) float64 {
		return n * 2
	}

	fmt.Println(numMutator(input, fn))
	// [2.4 4.6 6.8 9 11.2]
}
```

---

## Generic tipi `struct`larda Kullanmak

Generic tipleri fonksiyonlarda olduğu gibi `struct`’larda da kullanabiliriz.
Örneğin, aşağıdaki gibi bir struct tanımlayabiliriz:

https://go.dev/play/p/xpzfCN2fDST

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/generics-in-structs)

```go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// GradeType defines generic grade type.
type GradeType interface {
	constraints.Ordered
}

// AgeType defines generic age type.
type AgeType interface {
	constraints.Ordered
}

// Student represents generic student type model.
type Student[gradeType GradeType, ageType AgeType] struct {
	Name  string
	Age   gradeType
	Grade ageType
}

func main() {
	student := Student[int, float64]{
		Name:  "John",
		Age:   20,
		Grade: 10.21,
	}

	fmt.Printf("%+v\n", student) // {Name:John Age:20 Grade:10.21}
}
```

---

## Generic tipleri `map`’lerde kullanmak

Generic tipleri `map`’lerde de kullanabiliriz. Öncelikle generic bir map
tanımlayalım:

https://go.dev/play/p/daZVdBfA-xz

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/generics-in-maps)

```go
package main

import "fmt"

// GenericMap represents generic map type.
type GenericMap[K comparable, V int | string] map[K]V

func main() {
	m := GenericMap[string, int]{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Printf("%v\n", m) // map[one:1 three:3 two:2]
}
```

`comparable` aslında bir `interface`, tüm `comparable` (karşılaştırılabilir)
tipler bu `interface`’i implemente eder;

- boolean’ler
- nümerikler
- strings’ler
- pointer’lar
- channels
- karşılaştırılabilir elementlerden oluşan diziler (arrays of comparable types)
- alanı karşılaştırılabilir elementlerden oluşan struct’lar (whose fields are all comparable types)

---

## Generic Gerçek Hayat Örneği

Genericlerin gerçek hayatta nasıl kullanıldığına bakalım. Örneğin, bir
veritabanı oluştuyorsunuz. Bu veritabanında `User` ve `UserGrade` adında iki
tablonuz var. Veritabanına kayıtları eklemek için aşağıdaki gibi bir fonksiyon
tanımlayabilirsiniz:

```go
func InsertUser(user User) {
    // insert user into database
}

func InsertUserGrade(userGrade UserGrade) {
    // insert user grade into database
}
```

Bu durumda her tablo için ayrı bir fonksiyon tanımlamamız gerekir. Bu da kod
tekrarına neden olur ve hata ayıklamayı zorlaştırır.

Bu problemin çözümü için önce `User` ve `UserGrade` tipleri için standart bir
`interface` tanımlayalım:

```go
type Modelizer interface {
	TableName() string
}

type User struct {
	ID   int
	Name string
}

func (u User) TableName() string {
	return "user"
}

type UserGrade struct {
	ID     int
	UserID int
	Grade  int
}

func (u UserGrade) TableName() string {
	return "user_grade"
}
```

Şimdi `InsertUser` ve `InsertUserGrade` fonksiyonlarını aşağıdaki gibi tanımlayabiliriz:

```go
func Insert[T Modelizer](t T) {
	fmt.Println("Inserting into:", t.TableName())
	// do real insert operation
}
```

Bu fonksiyonu aşağıdaki gibi çağırabiliriz:

https://go.dev/play/p/kPpj-QPx2oJ

```go
package main

import "fmt"

type Modelizer interface {
	TableName() string
}

type User struct {
	ID   int
	Name string
}

func (u User) TableName() string {
	return "user"
}

type UserGrade struct {
	ID     int
	UserID int
	Grade  int
}

func (u UserGrade) TableName() string {
	return "user_grade"
}

func Insert[T Modelizer](t T) {
	fmt.Println("Inserting into:", t.TableName())
	// do real insert operation
}

func main() {
	user := User{
		ID:   1,
		Name: "Uğur Özyılmazel",
	}

	Insert(user)
	// Inserting into: user

	userGrade := UserGrade{
		ID:     1,
		UserID: 1,
		Grade:  100,
	}
	Insert(userGrade)
	// Inserting into: user_grade
}
```
