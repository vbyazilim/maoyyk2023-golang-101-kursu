# Bölüm 13/01: JSON İle Çalışmak

Go ile çok yüksek ihtimal ağ servisleri geliştireceksiniz. Dış dünyadan veri
almak ya da dış dünyadan istek atan bir istemciye cevap vermek için bir kaç
farklı yöntem mevcut. Bunlardan en sık kullanacaklarınız arasında `json` yani
hafif siklet veri değişim formatını kullancaksınız. Bu bakımdan;

1. Ya, `go` daki bir tipi, dış dünyanın anlayacağı `json` hale (`Marshal`)
1. Ya da dış dünyadan `json` formatında gelen veriyi `go`’nun anlayacağı hale (`Unmarshal`)

getirmek çok sık yaptığımız bir iş olacak. Bu dönüşüm esnasında 

---

## `encoding/json` Marshal

Elimizdeki `go` tipini `[]byte` haline dönüştürme işlemidir:

https://go.dev/play/p/nOzvhUsD40p

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// UserLevel is a type alias for defining user level.
type UserLevel string

// UserLevels holds collection of UserLevel types.
type UserLevels []UserLevel

func main() {
	UserLevelAdmin := UserLevel("admin")
	UserLevelModerator := UserLevel("moderator")
	UserLevelAnonymous := UserLevel("anonymous")

	userLevels := UserLevels{
		UserLevelAdmin,
		UserLevelModerator,
		UserLevelAnonymous,
	}

	j, err := json.Marshal(userLevels)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [91 34 97 100 109 105 110 34 44 34 109 111 100 101 114 97 116 111 114 34 44 34 97 110 111 110 121 109 111 117 115 34 93]

	fmt.Println(string(j))
	// ["admin","moderator","anonymous"]
}
```

`j`’nin string görüntüsü: `["admin","moderator","anonymous"]` şeklindedir. Bu
aslında JavaScript’teki `array` tipidir.

---

## `encoding/json` Unmarshal

Şimdi, tam ters işlemi yapalım; dış dünyadan bize `["admin","moderator","anonymous"]`
gelsin:

https://go.dev/play/p/7l7sIFzUh1S

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// UserLevel is a type alias for defining user level.
type UserLevel string

// UserLevels holds collection of UserLevel types.
type UserLevels []UserLevel

func main() {
	// byte slice from raw string.
	input := []byte(`["admin","moderator","anonymous"]`)

	// target data type for serialization.
	var userLevels UserLevels

	if err := json.Unmarshal(input, &userLevels); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", userLevels)
	// main.UserLevels{"admin", "moderator", "anonymous"}
}
```

---

## `json:"FIELD"` Tag’i

Genelde bu çevirme işlerini `struct` tipi üzerinde yaparız. Dışarıdan gelen ya
da dışarıya gidecek veride daha yapısal (structural) veri kullanmak isteriz:

https://go.dev/play/p/BB5D_VrLpLa

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents user model.
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	u := User{
		Name:  "Uğur Özyılmazel",
		Email: "vigo@example.com",
		Age:   51,
	}

	j, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [123 34 78 97 109 101 34 58 34 85 196 159 117 114 32 195 150 122 121 196 177 108 109 97 122 101 108 34 44 34 69 109 97 105 108 34 58 34 118 105 103 111 64 101 120 97 109 112 108 101 46 99 111 109 34 44 34 65 103 101 34 58 53 49 125]
	
	fmt.Println(string(j))
	// {"Name":"Uğur Özyılmazel","Email":"vigo@example.com","Age":51}
}
```

Dikkat ettiyseniz, `[]byte`’a baktığımızda; `{"Name":"Uğur
Özyılmazel","Email":"vigo@example.com","Age":51}` alan adlarının `struct`
field name’leri ile aynı olduğunu görürüz. `json` convention’a baktığımızda
alan adlarının bir kuralı var. Ana kural `key`’lerin (Name, Email ...) mutlaka
**Unicode** karakterlerden oluşmasıdır. Kimileri `camelCase` kimileri
`snake_case` de kullanabilir.

Kuralına uygun yazılmış bir rest-api servisi, alan adı kuralı olarak
`snake_case` kullanmalıdır.

Peki, biz `Name` yerine `name` nasıl döneceğiz? Bu durumda `json:"FIELD"` tag
devreye girer:

https://go.dev/play/p/EY_4u3pniOf

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents user model.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	u := User{
		Name:  "Uğur Özyılmazel",
		Email: "vigo@example.com",
		Age:   51,
	}

	j, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [123 34 110 97 109 101 34 58 34 85 196 159 117 114 32 195 150 122 121 196 177 108 109 97 122 101 108 34 44 34 101 109 97 105 108 34 58 34 118 105 103 111 64 101 120 97 109 112 108 101 46 99 111 109 34 44 34 97 103 101 34 58 53 49 125]

	fmt.Println(string(j))
	// {"name":"Uğur Özyılmazel","email":"vigo@example.com","age":51}
}
```

`json` tag’i (aslında bu bir struct annotation) sadece bu işe yaramaz;

```go
// JSON çıktıda key "myName" olarak görünür.
Field int `json:"myName"`

// JSON çıktıda key "myName" olarak görünür fakat değer "empty" (zero-value) ise
// omit edilir (yani dışarıda bırakılır) ve bu field çıktıda yer almaz!
Field int `json:"myName,omitempty"`

// JSON çıktıda key "Field" olarak görünür fakat değer "empty" (zero-value) ise
// omit edilir (yani dışarıda bırakılır) ve bu field çıktıda yer almaz!
Field int `json:",omitempty"`

// Field komple görmezden gelinir (ignore) ve çıktıda yer almaz!
Field int `json:"-"`

// JSON çıktıda key "-" olarak görünür.
Field int `json:"-,"`
```

Ek olarak, dış dünyadan gelen veri `string` görünümünde ama değer olarak
`integer`;

```json
{"age": "51"}
```

`age`’in tip olarak değeri `51` yani `string` ama biz bunu içeride `int`
olarak kullanmak istiyoruz (bazı kötü tasarlanmış api’larda bu tür durumlar oluyor):

```json
Age int `json:"age,string"`
```

şeklinde de kullanabiliyoruz:

https://go.dev/play/p/Qp0lnPOPBRQ

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents user model.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Age      int    `json:"age,string,omitempty"`
	Password string `json:"-"`
}

func main() {
	u := User{
		Name:  "Uğur Özyılmazel",
		Email: "vigo@example.com",
		Age:   51,
	}

	j, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [123 34 110 97 109 101 34 58 34 85 196 159 117 114 32 195 150 122 121 196 177 108 109 97 122 101 108 34 44 34 101 109 97 105 108 34 58 34 118 105 103 111 64 101 120 97 109 112 108 101 46 99 111 109 34 44 34 97 103 101 34 58 34 53 49 34 125]

	fmt.Println(string(j))
	// {"name":"Uğur Özyılmazel","email":"vigo@example.com","age":"51"}

	u2 := User{
		Name: "Uğur Özyılmazel",
	}

	j, err = json.Marshal(u2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [123 34 110 97 109 101 34 58 34 85 196 159 117 114 32 195 150 122 121 196 177 108 109 97 122 101 108 34 125]

	fmt.Println(string(j))
	// {"name":"Uğur Özyılmazel"}
}
```