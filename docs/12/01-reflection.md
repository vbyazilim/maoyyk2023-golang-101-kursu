# Bölüm 12/01: Reflection

Go statik yazıma sahiptir, bu da her değişkenin derleme zamanında bilinen
sabit bir tipe sahip olduğu anlamına gelir, bu da onun **statik tipi** olarak
bilinir.

En basit haliyle, `reflection`, bir programın yürütme sırasında kendi
değerlerini ve değişkenlerini incelemesine / incelemesine ve türlerini
belirlemesine olanak tanır.

Bir benzer senaryo, `json` verilerini kodlamamıza izin veren ve bilinmeyen
yapıları incelememize / manipüle etmemize izin veren yansımadır.

**Reflection**, üç kavram etrafında inşa edilmiştir:

1. Tür (Kind): aslında `uint` constant; `reflect.Bool (1)`
1. Tip (Type)
1. Değer (Value)

Hemen örneğe bakalım:

https://go.dev/play/p/kjCf6NmJSwE


```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	kinds := []reflect.Kind{
		reflect.Invalid,
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.Array,
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice,
		reflect.String,
		reflect.Struct,
		reflect.UnsafePointer,
	}

	for _, k := range kinds {
		fmt.Printf("%-16s : %[1]d (%v / %v)\n", k, reflect.ValueOf(k).Type(), reflect.ValueOf(k).Type().Kind())
	}
}

// invalid          : 0 (reflect.Kind / uint)
// bool             : 1 (reflect.Kind / uint)
// int              : 2 (reflect.Kind / uint)
// int8             : 3 (reflect.Kind / uint)
// int16            : 4 (reflect.Kind / uint)
// int32            : 5 (reflect.Kind / uint)
// int64            : 6 (reflect.Kind / uint)
// uint             : 7 (reflect.Kind / uint)
// uint8            : 8 (reflect.Kind / uint)
// uint16           : 9 (reflect.Kind / uint)
// uint32           : 10 (reflect.Kind / uint)
// uint64           : 11 (reflect.Kind / uint)
// uintptr          : 12 (reflect.Kind / uint)
// float32          : 13 (reflect.Kind / uint)
// float64          : 14 (reflect.Kind / uint)
// complex64        : 15 (reflect.Kind / uint)
// complex128       : 16 (reflect.Kind / uint)
// array            : 17 (reflect.Kind / uint)
// chan             : 18 (reflect.Kind / uint)
// func             : 19 (reflect.Kind / uint)
// interface        : 20 (reflect.Kind / uint)
// map              : 21 (reflect.Kind / uint)
// ptr              : 22 (reflect.Kind / uint)
// slice            : 23 (reflect.Kind / uint)
// string           : 24 (reflect.Kind / uint)
// struct           : 25 (reflect.Kind / uint)
// unsafe.Pointer   : 26 (reflect.Kind / uint)
```

Tür (Kind); `struct`, `int`, `string`, `slice`, `map` veya diğer go’nun basit
(primitive) türlerinden herhangi biri olabilir.

Eğer `Foo` adında bir `struct` tanımlarsanız, `Kind` bir `struct`’tır ve `Type`’ı `Foo`’dur. 

https://go.dev/play/p/R6D7WD8eJdE

```go
package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int
	B string
}

func main() {
	data := Foo{A: 1, B: "hello"}

	r := reflect.TypeOf(data)
	fmt.Println(r.Kind(), r.Name()) // struct Foo

	v := reflect.ValueOf(data)
	fmt.Println(v, v.Kind(), v.Type())
	// {1 hello} struct main.Foo
}
```

Kod örneğinde, `TypeOf` ve `ValueOf` fonksiyonları, sırasıyla, türü ve değeri
döndürür. 

**Reflection** paketini kullanarak bir `struct`’ın değerini okumaya (read) ek
olarak, değerini değiştirmek (write) için de kullanabiliriz. Bir değeri
`reflect.ValueOf` fonksiyonunu kullanarak değiştirmek için, bir değişkenin
pointer’ını geçirmemiz ve `Elem` fonksiyonunu çağırarak işaretçi adresindeki
değeri döndürmemiz gerekir:

https://go.dev/play/p/UzhantalrDx

```go
package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int
	B string
}

func main() {
	data := Foo{A: 1, B: "hello"}

	v := reflect.ValueOf(&data).Elem()
	v.Field(0).SetInt(100)
	v.Field(1).SetString("world")

	fmt.Println(v, v.Kind(), v.Type())
	// {100 world} struct main.Foo
}
```

Yukarıdaki örnekte görüldüğü gibi, `reflect.ValueOf` fonksiyonu, bir
`struct`’ın değerini değiştirmek için kullanılabilir. Değerin türüne bağlı
olarak;

- `SetInt()`
- `SetString()`
- `SetBool()`
- `SetFloat()`
- `SetUint()`
- `SetArray()`
- `SetMap()`
- `SetSlice()`
- `SetPointer()`
- `SetStruct()`
- `Set()`

gibi işlevler kullanılabilir.

https://pkg.go.dev/reflect@go1.21.0#Value.Set

`ClearValue` adlı bir fonksiyon oluşturmak istediğinizi varsayalım. Bu
fonksiyon, aldığı `any` türündeki verilerin yine metot parametresi
olarak aldığı `string` türündeki alanın değerini temizlemelidir.

https://go.dev/play/p/e29cVR9F2r9

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/reflect-clearvalue)

```go
package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// Foo1 represents dummy type.
type Foo1 struct {
	A int
	B string
}

// Foo2 represents dummy type.
type Foo2 struct {
	C int
	D string
}

// Foo3 represents dummy type.
type Foo3 struct {
	X bool
}

var (
	errInvalidKind  = errors.New("invalid kind")
	errInvalidValue = errors.New("invalid value")
)

func resetStringValueOfGivenField(d any, f string) error {
	v := reflect.ValueOf(d) // value of d
	k := v.Kind()           // kind of d

	if k != reflect.Ptr {
		return fmt.Errorf("%w, must be a pointer to struct, not a %q", errInvalidKind, k)
	}

	structValue := v.Elem()
	fieldValue := structValue.FieldByName(f)
	if !fieldValue.IsValid() {
		return fmt.Errorf("%w", errInvalidValue)
	}
	if fieldValue.Kind() != reflect.String {
		return fmt.Errorf("%w, %s value must be a string, not a %q", errInvalidKind, f, fieldValue.Kind())
	}
	fieldValue.SetString("")
	return nil
}

func main() {
	foo1 := Foo1{1, "hello"}
	fmt.Printf("foo1: %+v\n", foo1) // foo1: {A:1 B:hello}

	if err := resetStringValueOfGivenField(&foo1, "B"); err != nil {
		log.Print(err)
	}
	fmt.Printf("foo1 after: %+v\n", foo1) // foo1 after: {A:1 B:}

	foo2 := Foo2{2, "world"}
	fmt.Printf("foo2: %+v\n", foo2) // foo2: {C:2 D:world}

	if err := resetStringValueOfGivenField(&foo2, "D"); err != nil {
		log.Print(err)
	}
	fmt.Printf("foo2 after: %+v\n", foo2) // foo2 after: {C:2 D:}

	foo3 := Foo3{}
	if err := resetStringValueOfGivenField(&foo3, "X"); err != nil {
		log.Print(err) // invalid kind, X value must be a string, not a "bool"
	}
	fmt.Printf("foo3 after: %+v\n", foo3)

	// passing non-pointer
	if err := resetStringValueOfGivenField(foo3, "X"); err != nil {
		log.Print(err) // invalid kind, must be a pointer to struct, not a "struct"
	}
}
```

Keza, yine `any` alan bir fonksiyon olsun, geçilen tipi (Kind) tespit edelim:

https://go.dev/play/p/RbRn7tVQHV-

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/11/reflect-typecheck)

```go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func checkType(d any) {
	switch v := reflect.ValueOf(d); v.Kind() {
	case reflect.Struct:
		s := fmt.Sprintf("%+v", v)
		fmt.Printf("%-24v (%v)\n", s, v.Type())
	case reflect.String:
		fmt.Printf("%-24v (%v)\n", v, v.Type())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v -> %d\n", v, t, v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v -> %d\n", v, t, v.Uint())
	case reflect.Bool:
		fmt.Printf("%-24v (%v)\n", v, v.Type())
	case reflect.Float32, reflect.Float64:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v -> %f\n", v, t, v.Float())
	case reflect.Complex64, reflect.Complex128:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Func:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v\n", v, t)
	case reflect.Chan:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v\n", v, t)
	case reflect.Map:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Array, reflect.Slice:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Uintptr:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.UnsafePointer:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Pointer:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		el := reflect.TypeOf(d).Elem()
		if el.Kind() == reflect.Interface {
			t = "(" + el.Kind().String() + ")"
		}
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Interface:
		fmt.Println("this is not possible")
	case reflect.Invalid:
		fmt.Printf("%-24v\n", v)
	default:
		fmt.Println(v, "unknown")
	}
}

func main() {
	i := 1

	var err error

	data := []any{
		"hello",
		2023,
		int8(127),                    // upper limit
		int16(32767),                 // upper limit
		int32(2147483647),            // upper limit
		int64(9223372036854775807),   // upper limit
		uint(9223372036854775807),    // upper limit
		uint8(255),                   // upper limit
		uint16(65535),                // upper limit
		uint32(4294967295),           // upper limit
		uint64(18446744073709551615), // upper limit
		true,
		false,
		float32(1.0),
		float64(1.0),
		complex(float32(3.0), float32(4.0)),
		complex(5, 7),
		func() {},
		make(chan int),
		map[string]string{"key": "value"},
		[1]string{"array"},
		[]string{"slice"},
		uintptr(unsafe.Pointer(&i)), // nolint
		&i,
		unsafe.Pointer(&i), // nolint
		&err,
		struct{}{},
		nil,
	}

	for _, v := range data {
		checkType(v)
	}
}

// hello                    (string)
// 2023                     (int)      -> 2023
// 127                      (int8)     -> 127
// 32767                    (int16)    -> 32767
// 2147483647               (int32)    -> 2147483647
// 9223372036854775807      (int64)    -> 9223372036854775807
// 9223372036854775807      (uint)     -> 9223372036854775807
// 255                      (uint8)    -> 255
// 65535                    (uint16)   -> 65535
// 4294967295               (uint32)   -> 4294967295
// 18446744073709551615     (uint64)   -> 18446744073709551615
// true                     (bool)
// false                    (bool)
// 1                        (float32)  -> 1.000000
// 1                        (float64)  -> 1.000000
// (3+4i)                   (complex64)
// (5+7i)                   (complex128)
// 0x102b45b50              (func())
// 0x140001000c0            (chan int)
// map[key:value]           (map[string]string)
// [array]                  ([1]string)
// [slice]                  ([]string)
// 1374390648856            (uintptr)
// 0x14000110018            (*int)
// 0x14000110018            (unsafe.Pointer)
// 0x14000102020            (interface)
// {}                       (struct {})
// <invalid reflect.Value>
```
