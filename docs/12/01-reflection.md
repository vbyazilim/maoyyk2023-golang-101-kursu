# Bölüm 12/01: Reflection 

Go statik yazıma sahiptir, bu da her değişkenin derleme zamanında bilinen sabit bir tipe sahip olduğu anlamına gelir, bu da onun statik tipi olarak bilinir.

En basit haliyle, reflection, bir programın yürütme sırasında kendi değerlerini ve değişkenlerini incelemesine / incelemesine ve türlerini belirlemesine olanak tanır.

Bir benzer senaryo, JSON verilerini kodlamamıza izin veren ve bilinmeyen yapıları incelememize / manipüle etmemize izin veren yansımadır.  
 
Reflection, üç kavram etrafında inşa edilmiştir: Türler, Türler ve Değerler. Tür, struct, int, string, slice, map veya diğer Go ilkel türlerinden herhangi biri olabilir.

Eğer Foo adında bir struct tanımlarsanız, Kind bir struct'tır ve Type Foo'dur. 

```go
type Foo struct {
    A int
    B string
}

var data = Foo{A: 1, B: "hello"}
```

```go

func main() {
    r := reflect.TypeOf(data)
    fmt.Println(r.Kind(), r.Name()) // struct main.Foo

    v := reflect.ValueOf(data)
    fmt.Println(v, v.Kind(), v.Type()) // {1 hello} struct main.Foo
}

```

Kod örneğinde, TypeOf ve ValueOf fonksiyonları, sırasıyla, bir türün ve değerini döndürür.


In addition to reading, we can also use the reflection package to modify/write the value of a structure. To modify a value using the reflect.ValueOf function, we need to pass a pointer to the variable instead and call the Elem function, which will return the value at the pointer address.


Reflection paketini kullanarak bir struct'in değerini okumaya ek olarak, değerini değiştirmek / yazmak için kullanabiliriz. Bir değeri reflect.ValueOf fonksiyonunu kullanarak değiştirmek için, bir değişkenin pointer'ini geçirmemiz ve Elem fonksiyonunu çağırarak işaretçi adresindeki değeri döndürmemiz gerekir.

```go

type Foo struct {
    A int
    B string
}

func main() {
    data := Foo{A: 1, B: "hello"}
    v := reflect.ValueOf(&data).Elem()
    v.Field(0).SetInt(100)
    v.Field(1).SetString("world")
    fmt.Println(data) // {100 world}
}

```

Yukarıdaki örnekte görüldüğü gibi, reflect.ValueOf fonksiyonu, bir struct'in değerini değiştirmek için kullanılabilir. Değerin türüne bağlı olarak, SetInt, SetString, SetBool, SetFloat, SetUint, SetArray, SetMap, SetSlice, SetPointer, SetStruct ve Set gibi işlevler kullanılabilir.

# Gerçek hayat örneği

ClearValue adlı bir fonksiyon oluşturmak istediğinizi varsayalım. Bu fonksiyon, aldığı interface{} türündeki verilerin yine metot parametresi olarak aldığı string türündeki alanın değerini temizlemelidir.

```go


func ClearValue(data any, field string) {
	reflecType := reflect.ValueOf(data).Elem()
	fieldType := reflecType.FieldByName(field)
	fieldType.SetString("")
}

type Foo struct {
	A int
	B string
}

type Foo2 struct {
	C int
	B string
}

func main() {
	foo := Foo{A: 1, B: "2"}
	ClearValue(&foo, "B") // pass pointer to foo 

	foo2 := Foo2{C: 1, B: "2"}
	ClearValue(&foo2, "B") // pass pointer to foo2

	fmt.Println(foo)  // {"A":1,"B":""}
	fmt.Println(foo2) // {"C":1,"B":""}
}

```

Yukarıdaki örnekte görüldüğü gibi, ClearValue fonksiyonu, interface{} türündeki verilerin yine metot parametresi olarak aldığı string türündeki alanın değerini temizlemek için kullanılabilir fakat eğer temizlemek için verdiğimiz alanın o struct içerisinde olmadığını varsayarsak, bu durumda programımız hata verecektir. Bunun önüne geçmek için, field type'ın zero value olup olmadığını kontrol etmemiz gerekmektedir.

```go

func ClearValue(data any, field string) {
	reflecType := reflect.ValueOf(data).Elem()
	fieldType := reflecType.FieldByName(field)
	if fieldType != (reflect.Value{}) {
		fieldType.SetString("")
	}
}

```

