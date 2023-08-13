![Version](https://img.shields.io/badge/version-0.0.0-orange.svg?style=for-the-badge)
![Go Version](https://img.shields.io/badge/go-1.20.6-orange.svg?style=for-the-badge)

# Mustafa Akgül Özgür Yazılım Yaz Kampı 2023

## Golang 101 Kursu

[Uğur Özyılmazel][vigo] ve [Erhan Akpınar][erhan] tarafından Ağustos 2023 tarihinde verilen
Golang programlama dili kursu.

---

## Bölüm 01: Golang Hakkında Genel Bilgiler

1. [Tanıtım, Öne Çıkan Kısımları](docs/01/01-tanitim.md)
1. [Go Proverbs](docs/01/02-proverbs.md)
1. [Kurulum](docs/01/03-kurulum.md)

## Bölüm 02: Golang Uygulamasına Genel Bakış

1. [Paket Kavramı ve `main` Paketi](docs/02/01-paket-kavrami.md)
1. [Executable, Library ve Golang Uygulamasını Çalıştırmak](docs/02/02-executable-lib-run.md)

## Bölüm 03: Dil Kuralları

1. [Unicode Kavramı](docs/03/01-dil-kurallari.md#unicode)
1. [Anahtar Kelimeler](docs/03/01-dil-kurallari.md#anahtar-kelimeler)
1. [Operatörler ve İşaretçiler](docs/03/01-dil-kurallari.md#operat%C3%B6rler-ve-i%CC%87%C5%9Faret%C3%A7iler)
1. [Built-in (gömülü gelen) Veri Tipleri](docs/03/01-dil-kurallari.md#built-in-veri-tipleri)
1. [Kod Stili](docs/03/01-dil-kurallari.md#kod-stili)
1. [Sabitler](docs/03/02-sabitler.md)
1. [Değişkenler](docs/03/03-degiskenler.md)

## Bölüm 04: Veri Tipleri

1. [Strings (metinseller)](docs/04/01-string.md)
1. [Booleans (mantıksallar)](docs/04/02-bool.md)
1. [Numerics (sayısallar)](docs/04/03-numerics.md)
1. [Arrays (diziler)](docs/04/04-collections.md#array)
1. [Slices (dizi kesitleri)](docs/04/04-collections.md#slice)
1. [Structs (yapılar)](docs/04/05-struct.md)
1. [Struct Annotations](docs/04/05-struct-annotations.md)
1. [Pointer (işaretçi) Kavramı](docs/04/06-pointer.md)
1. [Struct Methods ve Receivers](docs/04/07-struct-methods-receivers.md)
1. [Maps](docs/04/08-map.md)
1. [Tip Dönüştürmek](docs/04/09-tip-donusturmek.md)

## Bölüm 05: Fonksiyonlar

1. [Function Signature (fonksiyon imzası)](docs/05/01-fonksiyonlar.md#signature)
1. [Argüman / Parametre](docs/05/01-fonksiyonlar.md#argumanlar)
1. [Return Values (fonksiyodan geriye dönen değerler)](docs/05/01-fonksiyonlar.md#return-values)
1. [Recursivity (özyineleme)](docs/05/01-fonksiyonlar.md#recursivity)
1. [Closure / Anonim Fonksiyonlar ve Function Scope (kapsama alanı)](docs/05/01-fonksiyonlar.md#closure)
1. [Deferring (erteleme)](docs/05/01-fonksiyonlar.md#defer)

## Bölüm 06: Durum Kontrolleri

1. [`if`, `else`, `else if`](docs/06/01-durum-koontrolleri-if.md)
1. [Short `if` declaration (kısa if tanımı)](docs/06/01-durum-koontrolleri-if.md#short-if)
1. [`switch` ve `case` İfadeleri](docs/06/02-durum-koontrolleri-switch.md)
1. [Label, `break` ve `goto` İşlemleri](docs/06/03-label-break-goto.md)

## Bölüm 07: Döngüler

1. [`C` stili döngü](docs/07/01-dongu.md#c-style)
1. [`range`](docs/07/01-dongu.md#range)
1. [`break` ve `continue`](docs/07/01-dongu.md#break-ve-continue)
1. [`for` ve Koşul Kullanımı](docs/07/01-dongu.md#for-ve-kosul)
1. [Label Kullanımı](docs/07/01-dongu.md#label-kullanimi)

## Bölüm 08: Interface

1. [Tanımı](docs/08/01-interface.md)
1. [Tip Olarak **empty interface** ya da `any`](docs/08/01-interface.md#empty-interface)
1. [Tip Kontrol Meknizması](docs/08/01-interface.md#tip-kontrol-mekanizmasi)
1. [Davranış Olarak `interface`](docs/08/01-interface.md#satisfying interface)

## Bölüm 09: Error

1. [`error` Nedir?](docs/09/01-error.md)
1. [Custom Error Types](docs/09/01-error.md#custom-error-types) (özelleştirilmiş error tipi oluşturmak)
1. [Wrapping](docs/09/01-error.md#wrapping) (sarmalama)
1. [Unwrapping](docs/09/01-error.md#unwrapping) (sarmalı açma)
1. [`error` Tip Kontrolleri](docs/09/01-error.md#tip-kontrolleri): `errors.Is` ve `errors.As`
1. [`panic` ve `recover`](docs/09/01-error.md#panic-ve-recover)
1. [Yaygın Pratikler](docs/09/01-error.md#yaygin-pratikler)

## Bölüm 10: `nil`

1. [`nil` Nedir?](docs/10/01-nil.md)
1. [Nerelerde ve Ne İçin Kullanınır?](docs/10/01-nil.md#nerelerde-kullanilir)

## Bölüm 11: Generics

1. [Nedir? Ne Amaçla Kullanılır](docs/11/01-generics.md)
1. [Generic Types](docs/11/01-generics.md)
1. [Generic Functions](docs/11/01-generics.md)
1. [Generic Interfaces](docs/11/01-generics.md)

## Bölüm 12: Reflection

1. [Ne İşe Yarar? Faydaları ve Zararları](docs/12/01-reflection.md)
1. [Tip Kontrolleri](docs/12/01-reflection.md)
1. @wip

## Bölüm 13: JSON İle Çalışmak

1. [Genel Bilgi](docs/13/01-json-ile-calismak.md)
1. [Encoding (Marshal)](docs/13/01-json-ile-calismak.md#encodingjson-marshal)
1. [Decoding (Unmarshal)](docs/13/01-json-ile-calismak.md#encodingjson-unmarshal)
1. [`json:"TAG"`](docs/13/01-json-ile-calismak.md#jsonfield-tagi)
1. [Custom Decoding](docs/13/01-json-ile-calismak.md#custom-decoding)
1. [Generic Interface](docs/13/01-json-ile-calismak.md#generic-interface)
1. [Reference Types](docs/13/01-json-ile-calismak.md)
1. [Streaming Encoders ve Decoders](docs/13/01-json-ile-calismak.md)

## Bölüm 14: Test

1. [Test Nedir? Neden Yazılır?](docs/14/01-test.md)
1. [Test Nasıl Çalıştırılır](docs/14/01-test.md)
1. [Examples ve `godoc` Nedir?](docs/14/01-test.md)
1. [Race Detection Nedir?](docs/14/01-test.md)
1. [SetUp ve TearDown Nedir?](docs/14/01-test.md)
1. [Table Driven Test Nedir?](docs/14/02-table-driven-test.md)
1. [Testlerin Paralel Çalıştırılması?](docs/14/02-test.md)
1. [Code Coverage Nedir?](docs/14/03-test-coverage.md)
1. [Benchmarking ve Profiling Nedir?](docs/14/04-profiling.md)
1. [Escape Analysis](docs/14/04-profiling.md)
1. [Memory ve CPU Profiling Temelleri](docs/14/04-profiling.md)

## Bölüm 15: Concurrency

1. Nedir? Golang’in Concurrency Stratejisi Nedir?
1. Goroutine Nedir?
1. `go` Kelimesiyle Başlayan Anonim Fonksiyonlar
1. WaitGroup Nedir?
1. Mutex Nedir?
1. Channels
1. Context
1. Concurrency Pratikleri (yaygın kullanılan desenler)

## Bölüm 16: İsimlendirme Kuralları

1. [Naming Conventions](docs/16/01-naming-conventions.md)
1. [Değişken İsimlendirmesi](docs/16/01-naming-conventions.md#degisken-adlari)
1. [Fonksiyon İsimlendirmesi](docs/16/01-naming-conventions.md#fonksiyon-adlari)
1. [Method’ların İsimlendirmesi](docs/16/01-naming-conventions.md#metod-adlari)
1. [Interface’lerin İsimlendirmesi](docs/16/01-naming-conventions.md#interface-adlari)
1. [Paketlerin İsimlendirmesi](docs/16/01-naming-conventions.md#paket-adlari)

## Bölüm 17: Golang Paketi Geliştirmek

1. [`golangci-linter` Kurulumu ve Konfigürasyonu](docs/17/01-kurulumlar.md)
1. [Go Modülü Anatomisi](docs/17/02-go-modul-anotomisi.md)
1. [Geliştirilmiş Paket Örneği](docs/17/03-ornek-paket.md)
1. [Paketi Dağıtmak / Paylaşmak](docs/17/04-paketi-dagitmak.md)

## Bölüm 18: In-Memory Key-Value Store

1. [`http` Paketini Kullanarak Rest-API Tasarlamak](docs/18/01-http-paketi.md)
1. [Domain Driven Design prensibini Kullanmak](docs/18/02-ddd-basics.md)
1. [Geliştirme Ön Hazırlık](docs/18/03-development.md)
1. [GitHub Actions ile Linter/Checker Kullanımı](docs/18/03-development.md)
1. [http server’ın Unit Testleri](docs/18/04-testing.md)
1. [Uygulamanın Docker Container’ından Çalıştırılması](docs/18/05-docker.md)

---

## Katkı

Hata raporları ve katkı istekleri,
https://github.com/vbyazilim/maoyyk2023-golang-101-kursu adresindeki GitHub
ortamında herkese açıktır. Bu projenin, işbirliği için güvenli ve davetkar bir
alan olması amaçlanmıştır ve katkıda bulunanların [Katkıcı Ahdi Topluluk
Sözleşmesi][COC] davranış kurallarına uyması beklenir.

---

## Lisans

Bu projed [MIT](https://opensource.org/licenses/MIT) lisansı kullanılmıştır.

---

## Katkıcı Ahdi Topluluk Sözleşmesi

Bu projenin kaynak kodunda, sorun izleyicilerinde, sohbet odalarında ve posta
listelerinde etkileşimde bulunan herkesin [davranış kurallarına][COC] uyması
beklenir.

---

[COC]:   https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/blob/main/CODE_OF_CONDUCT.md
[vigo]:  https://github.com/vigo
[erhan]: https://github.com/erhanakp
