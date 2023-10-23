![Version](https://img.shields.io/badge/version-0.1.8-orange.svg?style=for-the-badge)
![Go Version](https://img.shields.io/badge/go-1.20.6-orange.svg?style=for-the-badge)
![Powered by Rake](https://img.shields.io/badge/powered_by-rake-blue?logo=ruby&style=for-the-badge)

# Mustafa Akgül Özgür Yazılım Yaz Kampı 2023

## Golang 101 Kursu

[Uğur Özyılmazel][vigo] ve [Erhan Akpınar][erhan] tarafından Ağustos 2023 tarihinde verilen
Golang programlama dili kursu.

---

## Bölüm 01: Golang Hakkında Genel Bilgiler

1. [Tanıtım, Öne Çıkan Kısımları](docs/01/01-tanitim.md)
1. [Go Proverbs](docs/01/02-proverbs.md)
1. [Kurulum](docs/01/03-kurulum.md)
1. [VSCode Kurulumu](docs/01/04-vscode-kurulumu.md)

## Bölüm 02: Golang Uygulamasına Genel Bakış

1. [Paket Kavramı ve `main` Paketi](docs/02/01-paket-kavrami.md)
1. [Executable, Library ve Golang Uygulamasını Çalıştırmak](docs/02/02-executable-lib-run.md)

## Bölüm 03: Dil Kuralları

1. [Encoding Nedir?](docs/03/01-dil-kurallari.md#encoding-nedir)
1. [Unicode Desteği](docs/03/01-dil-kurallari.md#unicode-desteği)
1. [Anahtar Kelimeler](docs/03/01-dil-kurallari.md#anahtar-kelimeler)
1. [Operatörler ve İşaretçiler](docs/03/01-dil-kurallari.md#operatörler)
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
1. [Argüman / Parametre](docs/05/01-fonksiyonlar.md#argümanlar)
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
1. [Tip Kontrol Meknizması](docs/08/01-interface.md#tip-kontrol-mekanizması)
1. [Davranış Olarak `interface`](docs/08/01-interface.md#satisfying-interface)

## Bölüm 09: Error

1. [`error` Nedir?](docs/09/01-error.md)
1. [Custom Error Types](docs/09/01-error.md#custom-error-types) (özelleştirilmiş error tipi oluşturmak)
1. [Wrapping](docs/09/01-error.md#wrapping) (sarmalama)
1. [Unwrapping](docs/09/01-error.md#unwrapping) (sarmalı açma)
1. [`error` Tip Kontrolleri](docs/09/01-error.md#tip-kontrolleri): `errors.Is` ve `errors.As`
1. [`panic` ve `recover`](docs/09/01-error.md#panic-ve-recover)
1. [Yaygın Pratikler](docs/09/01-error.md#yaygın-pratikler)

## Bölüm 10: `nil`

1. [`nil` Nedir?](docs/10/01-nil.md)
1. [Nerelerde ve Ne İçin Kullanınır?](docs/10/01-nil.md#nerelerde-kullanılır)

## Bölüm 11: Generics

1. [Nedir? Ne Amaçla Kullanılır](docs/11/01-generics.md)
1. [Fonksiyonlarda Genericler](docs/11/01-generics.md#fonksiyonlarda-genericler)
1. [Custom Tiplerde Generic](docs/11/01-generics.md#custom-tiplerde-genericler)
1. [Generic Fonksiyon Çağrıları](docs/11/01-generics.md#generic-fonksiyon-çağrıları)
1. [Generic Tipi struct’da Kullanmak](docs/11/01-generics.md#generic-tipi-structlarda-kullanmak)
1. [Generic Tipi map’lerde Kullanmak](docs/11/01-generics.md#generic-tipleri-maplerde-kullanmak)
1. [Generic Gerçek Hayat Örneği](docs/11/01-generics.md#generic-gerçek-hayat-örneği)

## Bölüm 12: Reflection

1. [Ne İşe Yarar? Faydaları ve Zararları](docs/12/01-reflection.md)

## Bölüm 13: JSON İle Çalışmak

1. [Genel Bilgi](docs/13/01-json-ile-calismak.md)
1. [Encoding (Marshal)](docs/13/01-json-ile-calismak.md#encodingjson-marshal)
1. [Decoding (Unmarshal)](docs/13/01-json-ile-calismak.md#encodingjson-unmarshal)
1. [`json:"TAG"`](docs/13/01-json-ile-calismak.md#jsonfield-tagi)
1. [Custom Decoding](docs/13/01-json-ile-calismak.md#custom-decoding)
1. [Custom Encoding](docs/13/01-json-ile-calismak.md#custom-encoding)
1. [Generic Interface](docs/13/01-json-ile-calismak.md#generic-interface)
1. [Streaming Encoders ve Decoders](docs/13/01-json-ile-calismak.md#streaming)

## Bölüm 14: Test

1. [Test Nedir? Neden Yazılır?](docs/14/01-test.md)
1. [Test Nasıl Çalıştırılır](docs/14/01-test.md#test-nasıl-çalıştırılır)
1. [Examples ve `godoc` Nedir?](docs/14/01-test.md#examples-ve-godoc)
1. [Race Detection Nedir?](docs/14/01-test.md#data-race-detection)
1. [Table Driven Test Nedir?](docs/14/02-table-driven-test.md)
1. [Sub Tests](docs/14/02-table-driven-test.md#sub-tests)
1. [SetUp ve TearDown Nedir?](docs/14/02-table-driven-test.md#setup-ve-teardown)
1. [Testlerin Paralel Çalıştırılması?](docs/14/02-table-driven-test.md#paralel-test)
1. [Code Coverage Nedir?](docs/14/03-test-coverage.md)
1. [Benchmarking Nedir?](docs/14/04-profiling.md#benchmarking)
1. [Escape Analysis](docs/14/04-profiling.md#escape-analysis)
1. [Memory ve CPU Profiling Temelleri](docs/14/04-profiling.md#memory-ve-cpu-profiling-temelleri)

## Bölüm 15: Concurrency

1. [Nedir? Golang’in Concurrency Stratejisi Nedir?](docs/15/01-concurrency.md)
1. [Goroutine Nedir?](docs/15/01-concurrency.md#goroutine)
1. [`go` Kelimesiyle Başlayan Anonim Fonksiyonlar](docs/15/01-concurrency.md#go-anahtar-kelimesi)
1. [WaitGroup Nedir?](docs/15/01-concurrency.md#waitgroup)
1. [Channels](docs/15/02-channels.md)
1. [`done` Pattern](docs/15/02-channels.md#done-pattern)
1. [Deadlock](docs/15/02-channels.md#deadlock)
1. [Range Over Channels](docs/15/02-channels.md#range-over-channels)
1. [Buffered Channels](docs/15/02-channels.md#buffered-channels)
1. [Semaphore Pattern](docs/15/02-channels.md#semaphore-pattern)
1. [Fan Out Pattern](docs/15/02-channels.md#fan-out-pattern)
1. [`select`](docs/15/02-channels.md#select)
1. [Ticker](docs/15/02-channels.md#ticker)
1. [Worker Pattern](docs/15/02-channels.md#worker-pattern)
1. [Mutex Nedir?](docs/15/03-mutex.md)
1. [Context](docs/15/04-context.md)

## Bölüm 16: İsimlendirme Kuralları

1. [Naming Conventions](docs/16/01-naming-conventions.md)
1. [Değişken İsimlendirmesi](docs/16/01-naming-conventions.md#değişken-adları)
1. [Fonksiyon İsimlendirmesi](docs/16/01-naming-conventions.md#fonksiyon-adları)
1. [Method’ların İsimlendirmesi](docs/16/01-naming-conventions.md#metot-adları)
1. [Interface’lerin İsimlendirmesi](docs/16/01-naming-conventions.md#interface-adları)
1. [Paketlerin İsimlendirmesi](docs/16/01-naming-conventions.md#paket-adları)

## Bölüm 17: Golang Paketi Geliştirmek

1. [`golangci-linter` Kurulumu ve Konfigürasyonu](docs/17/01-kurulumlar.md)
1. [Go Modülü Anatomisi](docs/17/02-go-modul-anotomisi.md)
1. [`stringutils` Paketi](docs/17/03-ornek-paket.md)
1. [Paketi Dağıtmak / Paylaşmak](docs/17/04-paketi-dagitmak.md)

## Bölüm 18: In-Memory Key-Value Store

1. [`http` Paketini Kullanarak Rest-API Tasarlamak](docs/18/01-http-paketi.md)
1. [Domain Driven Design prensibini Kullanmak](docs/18/02-ddd-basics.md)
1. [Geliştirme](docs/18/03-development.md)
1. [http server’ın Unit Testleri](docs/18/04-testing.md)
1. [Uygulamanın Docker Container’ından Çalıştırılması](docs/18/05-docker.md)
1. [GitHub Actions ile Linter/Checker Kullanımı](docs/18/06-github-actions.md)
1. [Açık Kaynak Haline Getirmek](docs/18/07-acik-kaynak.md)

## Bonus

1. [Faydalı Linkler](docs/bonus/01-links.md)
1. [kvstore](https://github.com/vbyazilim/kvstore)

## Quiz

1. [Quiz 1](docs/quiz/01-go-cash-register.md)

---

## mkdocs

Otomatik doküman oluşturmak için;

```bash
pip install -r requirements.txt

rake -T

rake mkdocs:build       # build docs
rake mkdocs:deploy      # deploy to GitHub
rake mkdocs:serve       # run docs server
rake release[revision]  # release new version major,minor,patch, default: patch
```

---

## Katkı

Hata raporları ve katkı istekleri,
https://github.com/vbyazilim/maoyyk2023-golang-101-kursu adresindeki GitHub
ortamında herkese açıktır. Bu projenin, işbirliği için güvenli ve davetkar bir
alan olması amaçlanmıştır ve katkıda bulunanların [Katkıcı Ahdi Topluluk
Sözleşmesi][COC] davranış kurallarına uyması beklenir.

---

## Lisans

Bu projede [MIT](https://opensource.org/licenses/MIT) lisansı kullanılmıştır.

---

## Katkıcı Ahdi Topluluk Sözleşmesi

Bu projenin kaynak kodunda, sorun izleyicilerinde, sohbet odalarında ve posta
listelerinde etkileşimde bulunan herkesin [davranış kurallarına][COC] uyması
beklenir.

---

[COC]:   https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/blob/main/CODE_OF_CONDUCT.md
[vigo]:  https://github.com/vigo
[erhan]: https://github.com/erhanakp
