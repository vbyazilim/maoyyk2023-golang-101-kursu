# Bölüm 07/01: Döngüler

Döngü ve iterasyon işleri için go’da tek bir komut var o da `for`. Diğer
dillerdeki gibi `Do - While` gibi şeyler yok. Otomatik olarak tekrar edilecek
şeyler için `for` döngüsü kullanırız.

---

## `C`-Style

`JavaScript` ile uğraşanların yakından tanıdığı, `C`’deki döngü yapısı go’da
da var, buna **explicit control** (açık, belirgin) denir:

```go
// i sadece {} içinde yaşar, kapsam (scope)
for i := 0; i < 10; i++ {
   // yapılacak işler
}
```

Üç aşamalı işlem;

1. initialize; `i := 0`
1. check; `i < 10`
1. augmentation; `i++` (arttır, ya da eksilt)

---

## `range`

Önceki bölümlerde de görmüştük, `range` neredeyse en çok kullanacağımız döngü
şeklidir. Bu stile **implicit control** (dahili, gizli) denir. **Ranging
over** yani nelerin içinde bu iterasyonları yapabiliriz?

- Koleksiyonlar; (Array, Slice, String)
- `map`
- Channel

Örneğe bakalım, index, index + value, blank identifer+value kullanımları:

https://go.dev/play/p/BvDqmQzoECi

```go
package main

import "fmt"

func main() {
	users := []string{"vigo", "lego", "turbo"} // slice

	// ilk değer her zaman slice'ın index'i
	for i := range users {
		fmt.Println("index", i)
	}
	// index 0
	// index 1
	// index 2

	// index, element durumu;
	// ilk değer her zaman slice'ın index'i
	// diğer değer koleksiyondaki element'in o index'deki değeri
	for i, user := range users {
		fmt.Println(i, user)
	}
	// 0 vigo
	// 1 lego
	// 2 turbo

	// blank identifer kullanımı, ilk değeri yutuyoruz
	// diğer değer koleksiyondaki element'in o index'deki değeri
	for _, user := range users {
		fmt.Println(user)
	}
	// vigo
	// lego
	// turbo
}
```

`_` **blank identifier** ile değeri yutuyoruz, bunu **compile time**’de görme,
kullanma diyoruz. Keza `_` için **untyped, reusable variable placeholder** da
denir.

`map` konusunda değinmiştik; `map`’de de `range` iterasyon yapıyoruz:

https://go.dev/play/p/kgj16srGhGx

```go
package main

import "fmt"

func main() {
	users := map[string]string{
		"vigo":  "user",
		"turbo": "admin",
		"lego":  "superadmin",
	}

	// ilk değer key
	for user := range users {
		fmt.Println(user, users[user])
		// key, value
	}
	// vigo user
	// turbo admin
	// lego superadmin

	// ilk değer key
	// sonraki değer value
	for user, level := range users {
		fmt.Println(user, level)
	}
	// vigo user
	// turbo admin
	// lego superadmin
}
```

Sonsuz döngü için;

```go
// kontrol durumu yok (omit the condition) ~ while (true) durumu
for {
	// yapılacak iş
}
```


---

## `break` ve `continue`

Döngüyü kırıp çıkmak için `break` kullanırız, `Do-While` bezeri bir döngü;

https://go.dev/play/p/MyM4CYviApA

```go
package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println("i", i)
		i++
		if i > 5 {
			break
		}
	}
	fmt.Println("i'nin son değeri", i)
	// i for-loop'dan önce initialize edildiği için
	// bu noktada erişilebilir...
}

// i 0
// i 1
// i 2
// i 3
// i 4
// i 5
// i'nin son değeri 6
```

Döngü esnasında belli durumları pas geçmek için, bir sonraki iterasyona geçmek
için `continue` kullanırız:

https://go.dev/play/p/y3DkERb_t-n

```go
package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		// eğer i 2'ye eşitse sonraki iterasyona geç
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
// 0
// 1
// 3
```

---

## `for` ve Koşul

`for` tanımı yaparken koşul da verebiliriz:

https://go.dev/play/p/cDFY7r5b2rA

```go
package main

import "fmt"

func main() {
	sum := 1

	// bu döngü, sum 6'dan küçük olduğu sürece çalışır
	for sum < 6 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}
// 1
// 2 (1+1)
// 4 (2+2)
// 8 (4+4)
```

---

## Label Kullanımı

Aynı `switch`, `case`’de olduğu gibi, belli bir durum/koşulda çıkış
yapacağımız yeri gösterebiliriz:

https://go.dev/play/p/hUziRABXmit

```go
package main

import "fmt"

func main() {
outer: // label, en dış katman
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
			if j == 2 {
				break outer // loop'tan komple çık
			}
		}
	}
}

// 0 0
// 0 1
// 0 2
```