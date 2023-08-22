# Bölüm 17/03: Golang Paketi Geliştirmek

Sık kullandığımız yardımcı fonksiyonları bir kütüphane haline getirebiliriz.
Böylece başka projelerde de bu paketten faydalanabiliriz. **Reusable** yani
tekrar kullanılabilir kod parçaları oluşturabiliriz.

Paket aslında bir dizin (folder). Örneklerde sık kullandığımız `fmt` paketi.
Hemen kaynak koda bakalım:

```bash
$ ls -al "$(go env GOROOT)/src/fmt"
total 236K
drwxr-xr-x 16 vigo admin  512 Aug  4 23:14 .
drwxr-xr-x 73 vigo admin 2.3K Aug  4 23:14 ..
-rw-r--r--  1 vigo admin  15K Aug  4 23:14 doc.go
-rw-r--r--  1 vigo admin 1.7K Aug  4 23:14 errors.go
-rw-r--r--  1 vigo admin 3.7K Aug  4 23:14 errors_test.go
-rw-r--r--  1 vigo admin  12K Aug  4 23:14 example_test.go
-rw-r--r--  1 vigo admin  219 Aug  4 23:14 export_test.go
-rw-r--r--  1 vigo admin  59K Aug  4 23:14 fmt_test.go
-rw-r--r--  1 vigo admin  14K Aug  4 23:14 format.go
-rw-r--r--  1 vigo admin 1.6K Aug  4 23:14 gostringer_example_test.go
-rw-r--r--  1 vigo admin  32K Aug  4 23:14 print.go
-rw-r--r--  1 vigo admin  32K Aug  4 23:14 scan.go
-rw-r--r--  1 vigo admin  40K Aug  4 23:14 scan_test.go
-rw-r--r--  1 vigo admin 1.5K Aug  4 23:14 state_test.go
-rw-r--r--  1 vigo admin  551 Aug  4 23:14 stringer_example_test.go
-rw-r--r--  1 vigo admin 2.2K Aug  4 23:14 stringer_test.go
```

`doc.go`’ya bakalım:

```bash
$ cat "$(go env GOROOT)/src/fmt/doc.go"
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package fmt implements formatted I/O with functions analogous
to C's printf and scanf.  The format 'verbs' are derived from C's but
are simpler.

# Printing

The verbs:
.
.
.
Note: Fscan etc. can read one character (rune) past the input
they return, which means that a loop calling a scan routine
may skip some of the input.  This is usually a problem only
when there is no space between input values.  If the reader
provided to Fscan implements ReadRune, that method will be used
to read characters.  If the reader also implements UnreadRune,
that method will be used to save the character and successive
calls will not lose data.  To attach ReadRune and UnreadRune
methods to a reader without that capability, use
bufio.NewReader.
*/
package fmt
```

Kocaaaaman bir **comment** ve son satırda paketin adı `package fmt` yazıyor.
Bu dizindeki tüm dosyaların `package` deklarasyon kısmına baksak hepsinde de
`package fmt` yazdığını görürüz. Mantık şu;

    paket1/
        paket2/
            paket2.go
        paket1.go

Dikkat ettiyseniz `fmt` paketini `import` ederken;

```go
import "fmt"
```

şeklinde kullanıyoruz. yani;

```go
import foo/bar/baz/go/1.21.0/libexec/src/fmt
```

gibi bir tanım yok, çünkü `go` kurulunca otomatik olarak built-in paketlerin
nereye kurulduğunu biliyor;

```bash
$ go env GOROOT
/opt/homebrew/Cellar/go/1.21.0/libexec
```

ve paketlerin da `src/` altında olduğunu biliyor. Aslında biz `import "fmt"`
dediğimizde `go` otomatik olarak; `/opt/homebrew/Cellar/go/1.21.0/libexec/src/`
değerini ekliyor.

## `stringutils` Paketi

Evet, bir pakete ihtiyacımız var. İçinde `string`’lerle ilgili küçük küçük
fonksiyonlar olacak. `func Reverse(string)string` mesela. Bu paket için nasıl
isim vereceğiz?

Paket adını düşünürken hep nasıl import edeceğimi, paket içinden fonksiyonları
nasıl çağıracağımı düşünürüm. Eğer `python` ya da `ruby` kodu yazıyor olsak;
`utils.py` ya da `helper.rb` ya da `common.py` gibi bir dosya yapar geçerdik.

Go’da isimlendirme kurallarında bahsetmiştik; 

- Paketin amacına uygun bir isim olmalı
- Olası başka paket adlarıyla çakışmamalı. Ben olsam `uuid` diye bir paket
  yapmak durumunda kalsam, ayrıştırıcı bir isim düşünürüm: `simpleuuid` mesela...
- Paket adı mümkünse tek kelime olsun.

Dedik ya, `string`’ler için küçük yardımcı fonksiyonlar. Bu bakımdan adına;
`stringutils` diyorum ve bu paketi ben dahil herkes kullanabilsin diye GitHub’a
koymayı planlıyorum:

```bash
$ cd /path/to/development/
$ mkdir stringutils && cd stringutils

# go mod init github.com/<GITHUB-KULLANICI-ADINIZ>/stringutils
$ go mod init github.com/vigo/stringutils
go: creating new go.mod: module github.com/vigo/stringutils

$ ls
go.mod

$ cat go.mod 
module github.com/vigo/stringutils

go 1.21.0
```

Modül olarak **initialize** edilince, `go.mod` adında bir dosya oluşur. Bu
dosya içinde; paketin adı ve paketin bağımlılıkları yazar. Uygulamaya
katılacak her bağımlılık sonrası, yani kod içinde kullandığınız her ilave
paketin bu dosyada yer alması için, paket kurulumundan sonra mutlaka 
`go mod tidy` yapmak gerekiyor!.

Otomatik olarak, paket ekleyip çıkarttıkça `go.sum` dosyası da güncelleniyor.
Bu dosyada geriye dönük uyumluluk adına, kurulan tüm paketlerin ve hatta o
paketlerin de bağımlı olduğu paketlerin bir **hash-checksum** listesi duruyor.

Baze bu `go.mod` ve `go.sum` dosyaları canımızı sıkabiliyor. Hatta bazen bu
**module summary** olayını komple kapatıyoruz. `GONOSUMDB` environment
variable’ı sayesinde;

```bash
GONOSUMDB="github.com/vigo"
```

kurulan paketlerden başı `github.com/vigo` başlayanların `go.sum` kısmını
işleme katma diyoruz.

Bazen kütüphanelerimiz GitHub’a koyarız ama repo sadece bize erişilebilir
olur, yani repo **PRIVATE** olur. Bu tür durumlarda `GOPRIVATE` environment
variable’ı ile build mekanizmalarının private repo’lara ulaşmalarını da
sağlıyoruz;

```bash
$ GOPRIVATE="github.com/vbyazilim,*.vbyazilim,vigo.io/private"
```

şeklinde birden fazla domain ve **regex** kullanımı ile bu ayarlamayı yapıyoruz.

Artık bu kütüphane birileri tarafından kod içinde kullanılacağı zaman
`github.com/vigo/stringutils` üzerinden fonksiyonlara ulaşacaklar.

Şimdi dosyaları oluşturalım:

```bash
$ touch stringutils.go stringutils_test.go
```

Şimdi `stringutils.go` için;

```go
/*
Package stringutils implements basic string utility functions for demo
purposes only!
*/
package stringutils

// Reverse reverses given string!
func Reverse(s string) string {
	r := []rune(s)
	lr := len(r)
	ss := make([]rune, lr)

	for i := 0; i < lr; i++ {
		ss[lr-1-i] = r[i]
	}

	return string(ss)
}
```

ve `stringutils_test.go` için;

```go
package stringutils_test

import (
	"fmt"
	"testing"

	"github.com/vigo/stringutils"
)

func TestReverse(t *testing.T) {
	tcs := map[string]struct {
		input []string
		want  []string
	}{
		"none Turkish letters": {
			input: []string{"hello", "this is vigo"},
			want:  []string{"olleh", "ogiv si siht"},
		},
		"with Turkish letters": {
			input: []string{"uğur", "kırmızı şapka ve ÖĞRENCİ"},
			want:  []string{"ruğu", "İCNERĞÖ ev akpaş ızımrık"},
		},
		"with German letters": {
			input: []string{"Präzisionsmeßgerät"},
			want:  []string{"täregßemsnoisizärP"},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			for i, in := range tc.input {
				got := stringutils.Reverse(in)

				if got != tc.want[i] {
					fmt.Println(len(got), len(tc.want[i]))
					t.Errorf("want: %v; got: %v", tc.want[i], got)
				}
			}
		})
	}
}

var gs string

func BenchmarkReverse(b *testing.B) {
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = stringutils.Reverse("merhaba dünya!")
	}

	gs = s
}

func ExampleReverse() {
	fmt.Println(stringutils.Reverse("vigo"))
	// Output: ogiv
}
```

Hemen testleri çalıştıralım:

```bash
$ go test -v # ./... ya da paket adı vermedik!
=== RUN   TestReverse
=== RUN   TestReverse/none_Turkish_letters
=== RUN   TestReverse/with_Turkish_letters
=== RUN   TestReverse/with_German_letters
--- PASS: TestReverse (0.00s)
    --- PASS: TestReverse/none_Turkish_letters (0.00s)
    --- PASS: TestReverse/with_Turkish_letters (0.00s)
    --- PASS: TestReverse/with_German_letters (0.00s)
=== RUN   ExampleReverse
--- PASS: ExampleReverse (0.00s)
PASS
ok  	github.com/vigo/stringutils	0.911s
```

Paket artık dağıtıma hazır?
