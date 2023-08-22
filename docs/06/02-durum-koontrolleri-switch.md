# Bölüm 06/02: Durum Kontrolleri

## Switch / Case İfadeleri

Genelde birden fazla şeyi `if`, `else` ile kontrol etmek yerine `switch`,
`case` ifadelerini kullanırız:

https://go.dev/play/p/ghfE9nunfBE

```go
package main

import "fmt"

func main() {
	operatingSystem := "darwin"

	switch operatingSystem {
	case "darwin":
		fmt.Println("Mac OS Hipster")
		// otomatik olarak case'den çıkar,
		// durumların birbirine geçişi (fallthrough) varsayılan olarak kapalıdır.
	case "linux":
		fmt.Println("Linux Geek")
	default:
		// Windows, BSD, ...
		fmt.Println("Other")
	}
}

// Mac OS Hipster
```

aynı kodu `if` ile yazsak:

https://go.dev/play/p/yTltrjoqQ66

```go
package main

import "fmt"

func main() {
	operatingSystem := "darwin"

	if operatingSystem == "darwin" {
		fmt.Println("Mac OS Hipster")
	} else if operatingSystem == "linux" {
		fmt.Println("Linux Geek")
	} else {
		// Windows, BSD, ...
		fmt.Println("Other")
	}
}
```

`switch` deklarasyonu esnasında **identifier initialization** da yapılabilir;

https://go.dev/play/p/5Nja9VW1bBx

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS Hipster, your os is", os)
		// Mac OS Hipster, your os is darwin
	case "linux":
		fmt.Println("Linux Geek, your os is", os)
	default:
		fmt.Println("Other:", os)
	}
	// fmt.Println("os was", os)
	// undefined: os
}
```

`switch os := runtime.GOOS; os {` bu noktada `os` diye bir değişken
tanımladık, ve `switch`, `case` içinde kullandık, ömrü kısa; `switch` bitiminde
artık `os` diye bir şey yok. Kapsamı (scope) limitli.

,`case` içinde çoklu seçim ya da pas geçme işlemi de yapılabilir:

https://go.dev/play/p/2yFN0SshvNy

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS Hipster, your os is", os)
	case "commodore", "amiga":
		// not possible!
	case "linux":
		fmt.Println("Linux Geek, your os is", os)
	default:
		fmt.Println("Other:", os)
	}
}
```

`case` içinde ek kontrol de yapılabilir:

https://go.dev/play/p/naDSdGOF3kn

```go
package main

import "fmt"

func main() {
	number := 42
	switch {
	case number < 42:
		fmt.Println("küçük")
	case number == 42:
		fmt.Println("eşit")
	case number > 42:
		fmt.Println("büyük")
	}
}
// eşit
```
