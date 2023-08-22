# Bölüm 06/03: Durum Kontrolleri

## Label Kullanımı

Kodun akışı içinde, aynı makine dilindeki `jmp` (jump) gibi, bir yerden bir
yere zıplamak mümkün:

https://go.dev/play/p/LKOJu0advHS

```go
package main

import "fmt"

func main() {
switchStatement:
	switch 1 {
	case 1:
		fmt.Println("1") // 1
		for i := 0; i < 5; i++ {
			break switchStatement // daha ilk harekette switchStatement'dan çıkar ve fmt.Println("3") kısmına gider
		}
		fmt.Println("2")
	case 2:
	default:
		fmt.Println("default case...")
	}
	fmt.Println("3") // 3
}

// 1
// 3
```

---

## `goto` Kullanımı

Bilgisayar programlama dillerinin atası olan [B.A.S.I.C][01]’de olduğu gibi,
belli bir durum olduğunda kodun içinde başka bir yere gitmeyi sağlar:

https://go.dev/play/p/fPIMNugFU-S

```go
package main

import "fmt"

func main() {
	i := 0
Start:
	fmt.Println(i)
	if i > 2 {
		goto End
	} else {
		i += 1
		goto Start
	}
End:
}

// 0
// 1
// 2
// 3
```

[01]: https://en.wikipedia.org/wiki/BASIC