# Bölüm 11/01: Generics

Genericler, programlamada aynı kod yapısını farklı veri türleriyle kullanmamıza olanak tanıyan bir yaklaşımdır. 

Bu, kodun yeniden kullanılabilirliğini artırırken, tür güvenliğini de korur. Örneğin, liste veya dizi gibi veri yapılarını farklı veri tipleri için kullanabiliriz, böylece kod tekrarını önleriz ve hataları azaltırız. 

Genel olarak, genericler programlama dillerinde daha esnek ve verimli kod yazmamıza olanak sağlar.

Go programlama dilinde, başlangıçta genericler desteklenmiyordu. Ancak Go 1.18 sürümüyle birlikte, generic programlamayı destekleyen bir özellik olan "type parameters" (tür parametreleri) tanıtıldı. Bu, Go dilinde generic kod yazmayı mümkün kıldı.

---

# Fonksiyonlarda Genericler

Go'da generic fonksiyonlar, fonksiyonun parametrelerinde ve dönüş değerlerinde tür parametreleri kullanılarak tanımlanır.

Öncelikle non-generic bir fonksiyonu ele alalım:

```go
func add(a int, b int) int {
    return a + b
}
```

Bu şekilde tanımlanan fonksiyon, sadece `int` türü için çalışır. Bu fonksiyonu `float64` türü için kullanmak istediğimizde, aşağıdaki gibi bir hata alırız:

```go
func main() {
    fmt.Println(add(1.2, 2.3))
}
```

```
cannot use 1.2 (type float64) as type int in argument to add
```

Bunun için float64 türü için ayrı bir fonksiyon tanımlamamız gerekir:

```go
func addFloat(a float64, b float64) float64 {
    return a + b
}
```

Bu, kod tekrarına neden olur ve hata ayıklamayı zorlaştırır. Bu sorunu çözmek için, fonksiyonu generic olarak tanımlayabiliriz:

```go

func add[T int | float64](a T, b T) T {
    return a + b
}
```

Bu şekilde tanımlanan fonksiyon, `int` veya `float64` türü için çalışır. 

```go
func main() {
    fmt.Println(add(1, 2))
    fmt.Println(add(1.2, 2.3))
}
```

Ama bu şekilde tanımlanan fonksiyonlarda T için bütün tipleri eklememiz gerekir. Bu da fonksiyonun okunabilirliğini azaltır.

Örnek olarak, fonksiyon şu hale gelir:

```go
func add[T int | int8 | int16 | float32 | float64](a T, b T) T {
    return a + b
}
```

Bu sorunu çözmek için bir interface kullanabiliriz:

```go

type Number interface {
	int | int8 | int16 | float32 | float64
}

func add[T Number](a T, b T) T {
    return a + b
}
```

Fakat bu şekilde bir tanımalama yaptığımızda da bütün tipleri interface içerisine eklememiz gerekir. 

Bunun yerine constraints.Ordered interface'ini kullanabiliriz:

```go

func add[T constraints.Ordered](a T, b T) T {
    return a + b
}
```

constraints.Ordered interface'ini kullanmak için "golang.org/x/exp/constraints" paketini import etmemiz gerekir.
Bu gerekli tüm Integer | Float | ~string tipleri içerir.

---

# Custom tiplerde Genericler

Şimdi başa dönüp tekrar add fonksiyonunu ele alalım. Bu fonksiyonu, custom bir tipte kullanmak istediğimizi varsayalım:

```go
type SchoolNumber int

func add[T int](a T, b T) T {
    return a + b
}
```

Bu fonksiyonu SchoolNumber türü için kullanmak istediğimizde, aşağıdaki gibi bir hata alırız:

```go
func main() {
    fmt.Println(add(SchoolNumber(1), SchoolNumber(2)))
}
```

```
SchoolNumber does not satisfy int (possibly missing ~ for int in int)
```

Bu hatanın nedeni, add fonksiyonunun parametrelerinin int türü için tanımlanmış olmasıdır. Bu nedenle, SchoolNumber türü için kullanamayız. 

Bunu çözmek için add fonksiyonunu aşağıdaki gibi tanımlayabiliriz:

```go
func add[T ~int](a T, b T) T {
	return a + b
}
```

Tilda (~) işareti, T'nin int türü veya int türünden bir alt tür olması gerektiğini belirtir. SchoolNumber türü, int türünden bir alt tür olduğu için, bu fonksiyonu SchoolNumber türü için kullanabiliriz.

---

# Generic fonksiyon çağrıları

Şimdi generic fonksiyon çağrılarını ele alalım. Generic fonksiyonları çağırmak için, fonksiyonun parametrelerindeki tür parametrelerini belirtmemiz gerekir.

Öncelikle generic olmayan bir fonksiyonu ele alalım:

```go
 
input : [1, 2, 3, 4, 5], (n) => n * 2
output: [2, 4, 6, 8, 10] 
```

```go
func Mapper(values []int, fn func(int) int) []int {
    result := make([]int, len(values))
    for i, v := range values {
        result[i] = fn(v)
    }
    return result
}
```

Bu fonksiyonu aşağıdaki gibi çağırabiliriz:

```go
func main() {
    values := []int{1, 2, 3, 4, 5}
    result := Mapper(values, func(n int) int {
        return n * 2
    })
    fmt.Println(result) --> [2, 4, 6, 8, 10]
}
```

Bu fonksiyon sadece int türü için çalışır. Yine generic fonksiyonlar kullanarak bu fonksiyonu farklı türler için kullanabiliriz:

```go
func Mapper[T constraints.Ordered](values []T, fn func(T) T) []T {
    result := make([]T, len(values))
    for i, v := range values {
        result[i] = fn(v)
    }
    return result
}
```

Bu fonksiyonu aşağıdaki gibi çağırabiliriz:

```go
func main() {
    values := []int{1, 2, 3, 4, 5}
    result := Mapper(values, func(n int) int {
        return n * 2
    })
    fmt.Println(result) --> [2, 4, 6, 8, 10]
}
```
ya da float64 türü için:

```go
func main() {
    values := []float64{1.2, 2.3, 3.4, 4.5, 5.6}
    result := Mapper(values, func(n float64) float64 {
        return n * 2
    })
    fmt.Println(result) --> [2.4, 4.6, 6.8, 9.0, 11.2]
}
```
---

# Generic tipi structlarda kullanmak

Generic tipleri fonksiyonlarda olduğu gibi structlarda da kullanabiliriz. Örneğin, aşağıdaki gibi bir struct tanımlayabiliriz:

```go

type GradeType interface {
    constraints.Ordered
}

type AgeType interface {
    constraints.Ordered
}


type Student[gradeType GradeType, ageType AgeType] struct {
	Name  string
	Age   gradeType
	Grade ageType
}

func main() {
    student := Student[int, float64]{
        Name: "John",
        Age:  20,
        Grade: 10.21,
    }
    fmt.Println(student) --> {John 20 10.21}
}
```

---

# Generic tipleri maplerde kullanmak

Generic tipleri maplerde de kullanabiliriz. Öncelikle generic bir map tanımlayalım:

```go

type GenericMap[K comparable, V int | string] map[K]V

func main() {
    m := GenericMap[string, int]{
        "one": 1,
        "two": 2,
        "three": 3,
    }
    fmt.Println(m) --> map[one:1 two:2 three:3]
}
```

Not: comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types)

---

# Generic gerçek hayat örneği

Şimdi genericlerin gerçek hayatta nasıl kullanıldığına bakalım. Örneğin, bir veritabanı oluştuyorsunuz.

Bu veritabanında User ve UserGrade adında iki tablonuz var.

Veritabanına kayıtları eklemek için aşağıdaki gibi bir fonksiyon tanımlayabilirsiniz:

```go
func InsertUser(user User) {
    // insert user into database
}

func InsertUserGrade(userGrade UserGrade) {
    // insert user grade into database
}
```
Bu durumda her tablo için ayrı bir fonksiyon tanımlamamız gerekir. Bu da kod tekrarına neden olur ve hata ayıklamayı zorlaştırır.

Bu problemin çözümü için önce User ve UserGrade tipleri için standart bir interface tanımlayalım:

```go

type Base interface {
    TableName() string // User veya UserGrade insert edilecek tablonun adını döndürür.
}

User tablosunu değiştirelim:

type User struct {
    Base
    ID int
    Name string
}

func (u User) TableName() string {
	return "users"
}

UserGrade tablosunu değiştirelim:

type UserGrade struct {
    Base
    ID int
    Grade int
}

func (u UserGrade) TableName() string {
	return "user_grades"
}
```

Şimdi InsertUser ve InsertUserGrade fonksiyonlarını aşağıdaki gibi tanımlayabiliriz:

```go

func Insert[T Base](t T) {
    fmt.Println("Inserting into", t.TableName())
}
```

Bu fonksiyonu aşağıdaki gibi çağırabiliriz:

```go

func main() {
    user := User{
        ID : 1,
        Name: "John",
    }
    Insert(user)
    userGrade := UserGrade{
        ID : 1,
        Grade: 10,
    }
    Insert(userGrade)
}
```

---