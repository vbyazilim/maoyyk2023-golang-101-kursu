# Bölüm 04/04: Veri Tipleri / Koleksiyonlar

https://go.dev/blog/slices

**Composite Types** yani birleşik tipler olarak adlandırılan ailedendir.
**Array** ve **Slice** birbirleriyle kardeş iki kavramdır, aralarında ufak ama
önemli bir fark bulunur!

---

## Array

İçinde aynı tipte elemanların olduğu, boyutunun belli olduğu koleksiyona
**array** yani **dizi** deniyor. Bir dizi içinde sayısal ya da sadece metinsel
elemanlar olabiliyor. `Python` ya da `Ruby`’de bir dizi içindeki elemanların
tipleri farklı olabiliyordu ama go’da sadece tek tip oluyor. (Generics
haricinde...)

Bir diziyi tanımlamak için `[adet]Tip` şeklinde bir ifade kullanıyoruz.

`var a [3]int` dediğimizde `a` bir dizi ve boyu `3`. İçindeki elemanların tipi
de `int` demiş oluyoruz;

https://go.dev/play/p/WpeGUrzosN4

```go
package main

import "fmt"

func main() {
	var a [3]int

	fmt.Printf("%v %[1]T\n", a) // [0 0 0] [3]int
	
}
```

Elemanların tipi `int` olduğu için ve `int`’in default initializer’ı yani
zero-value’su `0` olduğu için `%v` bize `[0 0 0]` döndü. Yani içinde **3
adet** `0` olan bir dizi.

Array’ler aynı diğer dillerdeki gibi **zero-index-based** yani sıfırdan
başlayan sırayla erişilebilir haldedirler:

https://go.dev/play/p/vqIXPaUwZz9

```go
package main

import "fmt"

func main() {
	var a [3]int

	fmt.Printf("%v %[1]T\n", a) // [0 0 0] [3]int

	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Printf("%v\n", a) // [1 2 3]
}
```

Dizinin boyunu İngilizce **length** kelimesinin kısaltması olan `len` ile
alıyoruz, hatta tüm koleksiyonlar için `len` kullanıyoruz;

https://go.dev/play/p/RX2okvf6U9c

```go
package main

import "fmt"

func main() {
    var x [5]float64            // [0  0  0  0  0]
 
    x[0] = 98                   // [98 0  0  0  0]
    x[1] = 22                   // [98 22 0  0  0]
    x[2] = 31                   // [98 22 31 0  0]
    x[3] = 91                   // [98 22 31 91 0]
    x[4] = 7                    // [98 22 31 91 7]
	
	fmt.Printf("%v\n", len(x)) // 5
}
```

Kısa-yol (short-hand declaration) ile array tanımlamak için;

https://go.dev/play/p/BLSDUBtiy6n

```go
package main

import "fmt"

func main() {
	a := [3]int{10, 20, 30} // short-hand declaration

	fmt.Printf("%T %[1]v\n", a) // [3]int [10 20 30]
}
```

Eğer tanımlayacağımız array’in boyunu go derleyicisinin bulmasını istiyorsak;
`[...]` kullanıyoruz:

https://go.dev/play/p/WvbDaRXYLaF

```go
package main

import "fmt"

func main() {
	a := [...]int{10, 20, 30}

	fmt.Printf("%T %[1]v\n", a) // [3]int [10 20 30]
}
```

**Array’ler kopyalanır!**

Örnekteki `fruits` ve `otherFruits` başka başka array’lerdir:

https://go.dev/play/p/tqRKm7-ygDe

```go
package main

import "fmt"

var fruits = [...]string{"apple", "melon"}

func main() {
	otherFruits := fruits

	otherFruits[0] = "banana"

	fmt.Println("otherFruits", otherFruits, len(otherFruits), cap(otherFruits))
	// otherFruits [banana melon] 2 2

	fmt.Println("fruits", fruits, len(fruits), cap(fruits))
	// fruits [apple melon] 2 2

	// & ile değişkenin hafızadaki işaret (point) ettiği yeri alırız.
	fmt.Printf("fruits: %p\n", &fruits)           // fruits: 0x10276d3e0
	fmt.Printf("otherFruits: %p\n", &otherFruits) // otherFruits: 0x14000060020
}
```

---

## Slice

Slice aslında dinamik, ölçeklenebilir (scalable) array’dir. Array’in sabit bir
boyu (length’i) olurken, slice’ın olmayabilir. Bu da bize esneklik sağlar.
Aslında array’in bir parçasını tanımlar/ifade eder.

Profesyonel hayatta geliştirme yaparken neredeyse sadece Slice’ları kullanırız!

https://go.dev/play/p/o_Pw7VXiqA-

```go
package main

import "fmt"

func main() {
	var users []string // boş slice tanımı, len=0

	fmt.Println("users", users, len(users))                      // users [] 0
	fmt.Printf("users: %v || %#[1]v || %d\n", users, len(users)) // users: [] || []string(nil) || 0

	admins := []string{"vigo", "erhan"} // 2 elemanı olan slice

	fmt.Println("admins", admins, len(admins)) // admins [vigo erhan] 2
	fmt.Printf(
		"admins: %v || %#[1]v || %d\n",
		admins,
		len(admins),
	) // admins: [vigo erhan] || []string{"vigo", "erhan"} || 2
}
```

Slice’ın sonuna eleman eklemek için `append` kullanırız:

https://go.dev/play/p/tNxkiH_56ON

```go
package main

import (
	"fmt"
)

func main() {
	users := []string{"erhan", "vigo"}
	users = append(users, "turbo", "max", "move") // append multiple

	fmt.Printf("%T\n", users) // []string
	fmt.Printf("%v\n", users) // [erhan vigo turbo max move]
}
```

Slice içine özelleştirilmiş tip ekleyelim:

https://go.dev/play/p/K9-tLrBWkjt

```go
package main

import "fmt"

type user string // type definition, custom type

type users []user // type definition, custom type

func main() {
	var ourUsers users

	ourUsers = append(ourUsers, user("vigo"))
	ourUsers = append(ourUsers, user("erhan"))

	fmt.Printf("%#v\n", ourUsers) // main.users{"vigo", "erhan"}
	fmt.Println(len(ourUsers))    // 2

	ourUsers[0] = user("lego")
	fmt.Printf("%#v\n", ourUsers) // main.users{"lego", "erhan"}
}
```

Slice pointer (işaretçi) kullanır, array gibi kopya yapmaz ve içeride 3 şey saklar;

1. **pointer**: *element
1. **length**: `int`
1. **capacity**: `int`

`make` fonksiyonu ile slice tanımlar ve hafızaya yerleştiririz (*preallocate*).
`s := make([]byte, 5)` dediğimizde hafızada durum şu şekildedir;

    []byte
    
              +---+
    pointer   |   |---------> [5]byte
              +---+           +---+---+---+---+---+
    length    | 5 |           | 0 | 0 | 0 | 0 | 0 |
              +---+           +---+---+---+---+---+
    capacity  | 5 |
              +---+
 

Eğer bu `nil` slice olursa; yani `var s []byte`:

    []byte
    
              +-----+
    pointer   | nil |
              +-----+
    length    |  0  |
              +-----+
    capacity  |  0  |
              +-----+

İşaret (point) ettiği bir şey, bir adres yok, dolayısıyla `nil`:

```go
var s []byte
len(s) // 0
cap(s) // 0
for range s // iterates 0 times
s[i] // panic: index out of range
```

`var s []string` dediğimizde elimizdeki slice **nil slice**’dır ve `nil`
slice’a `append` yani ekleme yapılabilir.

Eğer slice initialize edilirken kapasite değeri verilmemişse, varsayılan
(default) kapasite **slice’ın length**’i kadar olur:

https://go.dev/play/p/hzNeqh1nxv8

```go
package main

import "fmt"

func main() {
	slice1 := make([]int, 5) // len: 5, default cap => 5
	slice2 := make([]int, 2) // len: 2, default cap => 2

	fmt.Printf("slice1, len: %d, cap: %d\n", len(slice1), cap(slice1)) // slice1, len: 5, cap: 5
	fmt.Printf("slice2, len: %d, cap: %d\n", len(slice2), cap(slice2)) // slice2, len: 2, cap: 2

	fmt.Println() // boş satır

	slice1 = append(slice1, 1)                                         // 1 eleman ekle, kapasiteyi aş, ek alan ekle
	fmt.Printf("slice1, len: %d, cap: %d\n", len(slice1), cap(slice1)) // slice1, len: 6, cap: 1

	slice2 = append(slice2, 1)                                         // 1 eleman ekle, kapasiteyi aş, ek alan ekle
	fmt.Printf("slice2, len: %d, cap: %d\n", len(slice2), cap(slice2)) // slice2, len: 3, cap: 4

	fmt.Println() // boş satır

	// kapasite varsayılan uzunluğa göre büyür

	slice3 := make([]string, 0, 4)                                     // len: 0, default cap => 4
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 0, cap: 4

	slice3 = append(slice3, "1 daha")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 1, cap: 4

	slice3 = append(slice3, "2 daha")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 2, cap: 4

	slice3 = append(slice3, "3 daha")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 3, cap: 4

	slice3 = append(slice3, "4 daha")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 4, cap: 4

	// güm! kapasiteyi arttır, taştık çünkü!
	slice3 = append(slice3, "more 5, overflow!")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 5, cap: 8

	// slice3[100] = "foo" // panic: runtime error: index out of range [100] with length 5
}
```

Kapasite, **altta yatan array**’e ya da slice’a (underlying) göre değişir.
Slice’dan başka bir slice çıkartabiliriz. Yeni slice’ın kapasiteside
kullandığı slice’ın kapasitesi ile, altaki slice’ın işaret (point) ettiği
yerdeki kapasiteden kalanla orantılı:

    users := []string{"erhan", "vigo", "turbo", "max", "move"}
    
              +---+
    pointer   |   |---------> [5]string
              +---+           +-------+------+-------+-----+------+
    length    | 5 |           | erhan | vigo | turbo | max | move |
              +---+           +-------+------+-------+-----+------+
    capacity  | 5 |
              +---+
    
    ---
              
    userSlice1 := users[1:2] // 0 başlangıç, 1’den al, 2’ye kadar, 2 hariç
    
              +---+
    pointer   |   |-----------------------+
              |   |                       ↓ 
              +---+           +-------+------+-------+-----+------+
    length    | 1 |           |       | vigo | ..... | ... | .... |
              +---+           +-------+------+-------+-----+------+
    capacity  | 4 |
              +---+
     
    ---
              
    userSlice2 := users[4:] // 0 başlangıç, 4’ten al, sona kadar, son dahil
    
              +---+
    pointer   |   |--------------------------------------------+
              |   |                                            ↓ 
              +---+           +-------+------+-------+-----+------+
    length    | 1 |           |       |      |       |     | move |
              +---+           +-------+------+-------+-----+------+
    capacity  | 1 |
              +---+

Kod olarak görelim:

https://go.dev/play/p/BN-Lr0i0180

```go
package main

import "fmt"

func main() {
	//  0       1        2       3      4
	// "erhan", "vigo", "turbo", "max", "move"
	users := []string{"erhan", "vigo", "turbo", "max", "move"}
	fmt.Printf("users - [:] - len: %v, cap: %v\n", len(users), cap(users))
	// users - len: 5, cap: 5

	userSlice1 := users[1:2]
	//  0       1        2       3      4
	// "erhan", *"vigo"*, "turbo", "max", "move"
	//             x    ,    x   ,   x  ,   x
	// 1'den itibaren 2'ye kadar, 2 hariç
	fmt.Printf("userSlice1 - [1:2] - len: %v, cap: %v\n", len(userSlice1), cap(userSlice1))
	// userSlice1 - [1:2] - len: 1, cap: 4

	userSlice2 := users[4:]
	//  0       1        2       3      4
	// "erhan", "vigo", "turbo", "max", *"move"*
	//                                ,    x
	// 4'ten itibaren sona kadar, son dahil
	fmt.Printf("userSlice2 - [4:] - len: %v, cap: %v\n", len(userSlice2), cap(userSlice2))
	// userSlice2 - [4:] - len: 1, cap: 1
}
```

Keza slicelar referans tiplerdir:

https://go.dev/play/p/YK1bjGumMEL

```go
package main

import "fmt"

func addOne(s []int) {
	fmt.Printf("s hafıza adresi: %p\n", s) // s hafıza adresi: 0x14000020180
	for i := range s {
		s[i]++ // index i'deki elemanın değerini 1 arttır
	}
}

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println("numbers - başlangıç değeri", numbers) // numbers - başlangıç değeri [1 2 3]

	addOne(numbers)
	fmt.Println("numbers - addOne sonrası", numbers)     // numbers - addOne sonrası [2 3 4]
	fmt.Printf("numbers - hafıza adresi: %p\n", numbers) // numbers - hafıza adresi: 0x14000020180
}
```

Başka bir örnek:

https://go.dev/play/p/yCEF399Sx7q

```go
package main

import "fmt"

// setZeroByIndex sadece demo amaçlı bir fonksiyon.
// olmayan index kontrolü yapılmamıştır.
func setZeroByIndex(s []int, i int) {
	fmt.Printf("s hafıza adresi: %p\n", s) // s hafıza adresi: 0x14000020150
	s[i] = 0
}

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println("numbers - başlangıç değeri", numbers) // numbers - başlangıç değeri [1 2 3]

	setZeroByIndex(numbers, 0)
	fmt.Println("numbers - setZeroByIndex 0 sonrası", numbers) // numbers - setZeroByIndex 0 sonrası [0 2 3]
	fmt.Printf("numbers - hafıza adresi: %p\n", numbers)       // numbers - hafıza adresi: 0x14000020150
}
```

Eğer `append` kullanılırsa, slice değişir, kopya çıkarılır, artık yeni bir
slice olur:

https://go.dev/play/p/PNysfmeb1vy

```go
package main

import "fmt"

func appendOne(s []int) {
	fmt.Printf("s hafıza adresi: %p\n", s)          // s hafıza adresi: 0x140000ac000
	s = append(s, 1)                                // bu s artık başka bir slice
	fmt.Printf("s hafıza adresi (append): %p\n", s) // s hafıza adresi (append): 0x140000b2000
}

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println("numbers - başlangıç değeri", numbers) // numbers - başlangıç değeri [1 2 3]

	appendOne(numbers)
	fmt.Println("numbers - appendOne sonrası", numbers)  // numbers - appendOne sonrası [1 2 3]
	fmt.Printf("numbers - hafıza adresi: %p\n", numbers) // numbers - hafıza adresi: 0x140000ac000
}
```

**Three-index slicing** ile yeni bir slice çıkartırken `[start:end:capacity]`
belirtebiliriz:

https://go.dev/play/p/-_dH6Cyvp39

```go
package main

import (
	"fmt"
)

func main() {
	users := []string{"Foo", "Bar", "Baz"}
	userSlice := users[1:2]               // 1'den başla (Bar), 2'ye kadar, 2 hariç
	userSliceWithCapacity := users[0:1:1] // 0'dan başla (Foo) 1'e kadar, 1 hariç, kapasite 1 olsun

	fmt.Printf("len: %v cap: %v\n", len(users), cap(users))                                 // len: 3 cap: 3
	fmt.Printf("len: %v cap: %v\n", len(userSlice), cap(userSlice))                         // len: 1 cap: 2
	fmt.Printf("len: %v cap: %v\n", len(userSliceWithCapacity), cap(userSliceWithCapacity)) // len: 1 cap: 1
}
```

Son olarak unutmamamız gereken bir konu; fonksiyona argüman olarak slice
göndermek gerekirse;

- Pointer’mı yollamalı?
- Direk slice’ı mı yollamalı?

```go
func checkUsers(users []string)   // ?

func checkUsers(users *[]string)  // ?
```

Büyük çoğunlukla (örnekteki gibi ise) direk slice’ı yollarız. Bazı durumlarda
**pointer to slice** göndeririz (gob, json, xml decoding, unmarshal gibi.)

[01]: https://www.willem.dev/articles/should-you-use-pointers-to-slices/
