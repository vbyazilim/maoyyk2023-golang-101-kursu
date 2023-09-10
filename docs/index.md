# İçindekiler

## Bölüm 01: Golang Hakkında Genel Bilgiler

1. [Tanıtım, Öne Çıkan Kısımları](01/01-tanitim.md)
1. [Go Proverbs](01/02-proverbs.md)
1. [Kurulum](01/03-kurulum.md)
1. [VSCode Kurulumu](01/04-vscode-kurulumu.md)

## Bölüm 02: Golang Uygulamasına Genel Bakış

1. [Paket Kavramı ve `main` Paketi](02/01-paket-kavrami.md)
1. [Executable, Library ve Golang Uygulamasını Çalıştırmak](02/02-executable-lib-run.md)

## Bölüm 03: Dil Kuralları

1. [Encoding Nedir?](03/01-dil-kurallari.md#encoding-nedir)
1. [Unicode Desteği](03/01-dil-kurallari.md#unicode-desteği)
1. [Anahtar Kelimeler](03/01-dil-kurallari.md#anahtar-kelimeler)
1. [Operatörler ve İşaretçiler](03/01-dil-kurallari.md#operatörler)
1. [Built-in (gömülü gelen) Veri Tipleri](03/01-dil-kurallari.md#built-in-veri-tipleri)
1. [Kod Stili](03/01-dil-kurallari.md#kod-stili)
1. [Sabitler](03/02-sabitler.md)
1. [Değişkenler](03/03-degiskenler.md)

## Bölüm 04: Veri Tipleri

1. [Strings (metinseller)](04/01-string.md)
1. [Booleans (mantıksallar)](04/02-bool.md)
1. [Numerics (sayısallar)](04/03-numerics.md)
1. [Arrays (diziler)](04/04-collections.md#array)
1. [Slices (dizi kesitleri)](04/04-collections.md#slice)
1. [Structs (yapılar)](04/05-struct.md)
1. [Struct Annotations](04/05-struct-annotations.md)
1. [Pointer (işaretçi) Kavramı](04/06-pointer.md)
1. [Struct Methods ve Receivers](04/07-struct-methods-receivers.md)
1. [Maps](04/08-map.md)
1. [Tip Dönüştürmek](04/09-tip-donusturmek.md)

## Bölüm 05: Fonksiyonlar

1. [Function Signature (fonksiyon imzası)](05/01-fonksiyonlar.md#signature)
1. [Argüman / Parametre](05/01-fonksiyonlar.md#argümanlar)
1. [Return Values (fonksiyodan geriye dönen değerler)](05/01-fonksiyonlar.md#return-values)
1. [Recursivity (özyineleme)](05/01-fonksiyonlar.md#recursivity)
1. [Closure / Anonim Fonksiyonlar ve Function Scope (kapsama alanı)](05/01-fonksiyonlar.md#closure)
1. [Deferring (erteleme)](05/01-fonksiyonlar.md#defer)

## Bölüm 06: Durum Kontrolleri

1. [`if`, `else`, `else if`](06/01-durum-koontrolleri-if.md)
1. [Short `if` declaration (kısa if tanımı)](06/01-durum-koontrolleri-if.md#short-if)
1. [`switch` ve `case` İfadeleri](06/02-durum-koontrolleri-switch.md)
1. [Label, `break` ve `goto` İşlemleri](06/03-label-break-goto.md)

## Bölüm 07: Döngüler

1. [`C` stili döngü](07/01-dongu.md#c-style)
1. [`range`](07/01-dongu.md#range)
1. [`break` ve `continue`](07/01-dongu.md#break-ve-continue)
1. [`for` ve Koşul Kullanımı](07/01-dongu.md#for-ve-kosul)
1. [Label Kullanımı](07/01-dongu.md#label-kullanimi)

## Bölüm 08: Interface

1. [Tanımı](08/01-interface.md)
1. [Tip Olarak **empty interface** ya da `any`](08/01-interface.md#empty-interface)
1. [Tip Kontrol Meknizması](08/01-interface.md#tip-kontrol-mekanizması)
1. [Davranış Olarak `interface`](08/01-interface.md#satisfying-interface)

## Bölüm 09: Error

1. [`error` Nedir?](09/01-error.md)
1. [Custom Error Types](09/01-error.md#custom-error-types) (özelleştirilmiş error tipi oluşturmak)
1. [Wrapping](09/01-error.md#wrapping) (sarmalama)
1. [Unwrapping](09/01-error.md#unwrapping) (sarmalı açma)
1. [`error` Tip Kontrolleri](09/01-error.md#tip-kontrolleri): `errors.Is` ve `errors.As`
1. [`panic` ve `recover`](09/01-error.md#panic-ve-recover)
1. [Yaygın Pratikler](09/01-error.md#yaygın-pratikler)

## Bölüm 10: `nil`

1. [`nil` Nedir?](10/01-nil.md)
1. [Nerelerde ve Ne İçin Kullanınır?](10/01-nil.md#nerelerde-kullanılır)

## Bölüm 11: Generics

1. [Nedir? Ne Amaçla Kullanılır](11/01-generics.md)
1. [Fonksiyonlarda Genericler](11/01-generics.md#fonksiyonlarda-genericler)
1. [Custom Tiplerde Generic](11/01-generics.md#custom-tiplerde-genericler)
1. [Generic Fonksiyon Çağrıları](11/01-generics.md#generic-fonksiyon-çağrıları)
1. [Generic Tipi struct’da Kullanmak](11/01-generics.md#generic-tipi-structlarda-kullanmak)
1. [Generic Tipi map’lerde Kullanmak](11/01-generics.md#generic-tipleri-maplerde-kullanmak)
1. [Generic Gerçek Hayat Örneği](11/01-generics.md#generic-gerçek-hayat-örneği)

## Bölüm 12: Reflection

1. [Ne İşe Yarar? Faydaları ve Zararları](12/01-reflection.md)

## Bölüm 13: JSON İle Çalışmak

1. [Genel Bilgi](13/01-json-ile-calismak.md)
1. [Encoding (Marshal)](13/01-json-ile-calismak.md#encodingjson-marshal)
1. [Decoding (Unmarshal)](13/01-json-ile-calismak.md#encodingjson-unmarshal)
1. [`json:"TAG"`](13/01-json-ile-calismak.md#jsonfield-tagi)
1. [Custom Decoding](13/01-json-ile-calismak.md#custom-decoding)
1. [Custom Encoding](13/01-json-ile-calismak.md#custom-encoding)
1. [Generic Interface](13/01-json-ile-calismak.md#generic-interface)
1. [Streaming Encoders ve Decoders](13/01-json-ile-calismak.md#streaming)

## Bölüm 14: Test

1. [Test Nedir? Neden Yazılır?](14/01-test.md)
1. [Test Nasıl Çalıştırılır](14/01-test.md#test-nasıl-çalıştırılır)
1. [Examples ve `godoc` Nedir?](14/01-test.md#examples-ve-godoc)
1. [Race Detection Nedir?](14/01-test.md#data-race-detection)
1. [Table Driven Test Nedir?](14/02-table-driven-test.md)
1. [Sub Tests](14/02-table-driven-test.md#sub-tests)
1. [SetUp ve TearDown Nedir?](14/02-table-driven-test.md#setup-ve-teardown)
1. [Testlerin Paralel Çalıştırılması?](14/02-table-driven-test.md#paralel-test)
1. [Code Coverage Nedir?](14/03-test-coverage.md)
1. [Benchmarking Nedir?](14/04-profiling.md#benchmarking)
1. [Escape Analysis](14/04-profiling.md#escape-analysis)
1. [Memory ve CPU Profiling Temelleri](14/04-profiling.md#memory-ve-cpu-profiling-temelleri)

## Bölüm 15: Concurrency

1. [Nedir? Golang’in Concurrency Stratejisi Nedir?](15/01-concurrency.md)
1. [Goroutine Nedir?](15/01-concurrency.md#goroutine)
1. [`go` Kelimesiyle Başlayan Anonim Fonksiyonlar](15/01-concurrency.md#go-anahtar-kelimesi)
1. [WaitGroup Nedir?](15/01-concurrency.md#waitgroup)
1. [Channels](15/02-channels.md)
1. [`done` Pattern](15/02-channels.md#done-pattern)
1. [Deadlock](15/02-channels.md#deadlock)
1. [Range Over Channels](15/02-channels.md#range-over-channels)
1. [Buffered Channels](15/02-channels.md#buffered-channels)
1. [Semaphore Pattern](15/02-channels.md#semaphore-pattern)
1. [Fan Out Pattern](15/02-channels.md#fan-out-pattern)
1. [`select`](15/02-channels.md#select)
1. [Ticker](15/02-channels.md#ticker)
1. [Worker Pattern](15/02-channels.md#worker-pattern)
1. [Mutex Nedir?](15/03-mutex.md)
1. [Context](15/04-context.md)

## Bölüm 16: İsimlendirme Kuralları

1. [Naming Conventions](16/01-naming-conventions.md)
1. [Değişken İsimlendirmesi](16/01-naming-conventions.md#değişken-adları)
1. [Fonksiyon İsimlendirmesi](16/01-naming-conventions.md#fonksiyon-adları)
1. [Method’ların İsimlendirmesi](16/01-naming-conventions.md#metot-adları)
1. [Interface’lerin İsimlendirmesi](16/01-naming-conventions.md#interface-adları)
1. [Paketlerin İsimlendirmesi](16/01-naming-conventions.md#paket-adları)

## Bölüm 17: Golang Paketi Geliştirmek

1. [`golangci-linter` Kurulumu ve Konfigürasyonu](17/01-kurulumlar.md)
1. [Go Modülü Anatomisi](17/02-go-modul-anotomisi.md)
1. [`stringutils` Paketi](17/03-ornek-paket.md)
1. [Paketi Dağıtmak / Paylaşmak](17/04-paketi-dagitmak.md)

## Bölüm 18: In-Memory Key-Value Store

1. [`http` Paketini Kullanarak Rest-API Tasarlamak](18/01-http-paketi.md)
1. [Domain Driven Design prensibini Kullanmak](18/02-ddd-basics.md)
1. [Geliştirme](18/03-development.md)
1. [http server’ın Unit Testleri](18/04-testing.md)
1. [Uygulamanın Docker Container’ından Çalıştırılması](18/05-docker.md)
1. [GitHub Actions ile Linter/Checker Kullanımı](18/06-github-actions.md)
1. [Açık Kaynak Haline Getirmek](18/07-acik-kaynak.md)

## Bonus

1. [Faydalı Linkler](bonus/01-links.md)
1. [kvstore](https://github.com/vbyazilim/kvstore)

## Quiz

1. [Quiz 1](quiz/01-go-cash-register.md)
