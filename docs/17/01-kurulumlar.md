# Bölüm 17/01: Golang Paketi Geliştirmek

Geliştirme yapmaya başlamadan önce bazı yardımcı araçlara ihtiyacımız var.
Bunların başında da `golangci-linter` geliyor.

https://golangci-lint.run/

Kurulumu;

https://golangci-lint.run/usage/install/#local-installation

adresinden takip ederek yapabilirsiniz. `golangci-linter` çalışırken config
yaml dosyası arar, bu dosya belli lokasyonlarda olabilir. Genelde `go.mod`
dosyasının bulunduğu, projenin **ROOT** dizinine koyarız bu dosyayı.

Komut satırında, ROOT dizindeyken;

```bash
$ golangci-lint run
```

diyerek tüm kontrolleri çalıştırabilirsiniz. Eğer hata yoksa geriye hiçbir şey
dönmez, bu durum işlerin yolunda olduğu anlamındadır.

Go, `gofmt` komutu ile beraber gelir;

```bash
$ command -v gofmt
/opt/homebrew/opt/go/libexec/bin/gofmt
```

Bu araç sayesinde yazdığınız kodun otomatik olarak formatlanması, yani düzgün
şekilde görünmesini sağlar:

```bash
$ gofmt -h
$ gofmt <dosya.go>      # düzeltilmiş kodu stdout’a yazar
$ gofmt -w <dosya.go>   # düzeltilmiş kodu <dosya.go> üzerine yazar
```

`gofumpt` ise `gofmt`’un daha da katı/kuralcı halidir ve aynı şekilde çalışır;

```bash
$ go install mvdan.cc/gofumpt@latest
$ gofumpt -h
$ gofumpt <dosya.go>      # düzeltilmiş kodu stdout’a yazar
$ gofumpt -w <dosya.go>   # düzeltilmiş kodu <dosya.go> üzerine yazar
```

Otomatik olarak `import` ifadelerinin tamamlanması işini `goimports` yapar;

```bash
$ go install golang.org/x/tools/cmd/goimports@latest
$ goimports -h

$ nano /tmp/main.go
```

Şimdi özellikle yamuk yumuk, import’u eksik bir go kodu:

```go
package main

func main(){
    fmt.Println("ok")
    }
```

Şimdi;

```bah
$ gofmt /tmp/main.go # sadece stdout’a çıktı
package main

func main() {
	fmt.Println("ok")
}

$ gofmt -w /tmp/main.go
$ gofumpt -w /tmp/main.go
$ goimports -w /tmp/main.go

$ cat /tmp/main.go 
package main

import "fmt"

func main() {
	fmt.Println("ok")
}
```

Go otomatik olarak girintileme için (indentation) `TAB` kullanır. Gördüğünüz
gibi kod otomatik olarak düzeltildi.

`golines` ile otomatik olarak uzun satırları daha okunur hale getirebiliriz.

```bash
$ go install github.com/segmentio/golines@latest
$ golines --help
$ golines -m 120 -w <dosya.go>
```

`go vet` aslında `go` ile built-in gelen, yine kodu analiz edip belli
düzeltmeleri bize söyler.

```bash
$ go vet ./...         # tüm paketleri vet’le
$ go vet <PAKET>       # <PAKET> vet’le
$ go vet github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/15/mutex/kvstore # gibi
```

`vet` yaparken yan araçlar da kullanırız: `shadow`;

```bash
$ go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
$ shadow -h
```

`shadow` bize istemeden yaptığımız **Variable Shadowing**’leri gösterir.
Özellikle `if val, ok := function; !ok{}` gibi ifadelerde istemeden değişken
gölgeleme yapmış olabiliriz.

```bash
$ go vet -vettool "$(command -v shadow)" ./...
$ go vet -vettool "$(command -v shadow)" <PAKET>
```

Kullandığınız kod editörleri genelde tüm bu linter/checker işlerini otomatik
olarak yapmanızı sağlar. Gerekli kurulumları yaptıktan sonra kod editörünüzü
de ayarlamanız gerekir.

`goimports` bazen hata yapar. Mesela `uuid` paketi. Eğer başka projelerde;
`github.com/gofrs/uuid` kullanmışsanız ve yeni projede
`github.com/google/uuid` kulanırsanız, `goimports` ilk bulduğu paketi otomatik
olarak takar :) Yani siz `github.com/google/uuid` bulmasını beklerken diğerini
görürseniz elle düzeltme yapmak gerekir.

Kurulan tüm paketler;

```bash
$ cd $(go env GOMODCACHE)
```

altına **clone** (`git clone`) edilir.
