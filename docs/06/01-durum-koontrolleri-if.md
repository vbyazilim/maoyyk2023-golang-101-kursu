# Bölüm 06/01: Durum Kontrolleri

Go kodu, çalışırken, kod yukarıdan aşağı (top to bottom) doğru çalıştırılır.
Buna **sequential execution** ya da **lineer execution** denir.

Belli durumlarda ise kontrol yapıları devreye girer ve durum kontrolleri
yapıldıktan sonra, ilgili durumun sonucuna göre hareket edilir.

Bazı durumlarda tekrar eden döngüler olur (for loops) yine döngünün bitişine
ya da sonsuzluğa gidişine göre çalıştırılacak kısımlar devreye alınır.

## `if`, `else`, `else if`

Eğer verilen `<CONDITION>` (durum) -> `true` ise; ya da `false` ise kod akışının gidişatı
belirlenir:

```go
// true
if <CONDITION> {
    ...
}

// false
if !<CONDITION> {
    ...
}
```

Eğer `<CONDITION>` -> `true` ise şunu yap, değilse bunu yap:

```go
if <CONDITION> {
    ...
} else {
    ....
}
```

Eğer `<CONDITION>` -> `A` ise şunu, eğer `<CONDITION>` -> `B` ise bunu,
hiçbiri ise şunu yap:

```go
if <CONDITION> == <A> {
    ...
} else if <CONDITION> == <B>  {
    ....
} else {
    ....
}
```

---

## Short If

**Short variable declaration** (kısa değişken tanımı) gibi **short if
declaration** yani kısa `if` tanımı da yapmak mümkündür;

```go
if <IDENTIFIER> := <FUNCTION-RESULT()>; <IDENTIFIER> != <COMPARE> {
    ...
}

// if err is not nil?
if err := result(); err != nil { // err sadece {} arasında yaşar
    log.Fatal(err.Error())
}
// fmt.Println(err) olamaz! hata alırız, err diye bir şey yok!
```

Kısa tanımda aslında **on-the-fly** yani o anda bir değişken tanımlaması da
yapmış oluyoruz. Ömrü sadece `{}` arasında geçerli olan, **initializer**
dediğimiz şey. Hatta;

```go
package main

import "fmt"

func main() {
	// x'i on-the-fly initialize et!
    if x := 100; x == 1001 {
		fmt.Println("bu mümkün değil!")
	}
}
```
