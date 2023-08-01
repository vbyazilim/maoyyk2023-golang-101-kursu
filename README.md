![Version](https://img.shields.io/badge/version-0.0.0-orange.svg?style=for-the-badge)

# Mustafa Akgül Özgür Yazılım Yaz Kampı 2023

## Golang 101 Kursu

[Uğur Özyılmazel][vigo] ve [Erhan Akpınar][erhan] tarafından Ağustos 2023 tarihinde verilen
Golang programlama dili kursu.

---

## Bölüm 01: Golang Hakkında Genel Bilgiler

1. Tanıtım, Öne Çıkan Kısımları
1. Go Proverbs
1. Kurulum

## Bölüm 02: Golang Uygulamasına Genel Bakış

1. Paket Kavramı
1. `main` Paketi
1. Executable (çalıştırılabilir uygulama) ve Library (yardımcı kütüphane)
1. Golang Uygulamasını Çalıştırmak

## Bölüm 03: Dil Kuralları

1. Unicode Kavramı
1. Anahtar Kelimeler
1. Operatörler ve İşaretçiler
1. Built-in (gömülü gelen) Veri Tipleri
1. Kod Stili
1. Sabitler
1. Değişkenler

## Bölüm 04: Veri Tipleri

1. String (metinseller)
1. Boolean (mantıksallar)
1. Numerics (sayısallar)
1. Arrays (diziler)
1. Slices (dizi kesitleri)
1. Structs (yapılar)
1. Pointer (işaretçi) Kavramı
1. Struct Methods ve Receivers
1. Maps
1. Tip Dönüştürmek

## Bölüm 05: Fonksiyonlar

1. Function Signature (fonksiyon imzası)
1. Argüman / Parametre
1. Return Values (fonksiyodan geriye dönen değerler)
1. Recursivity (özyineleme)
1. Closure / Anonim Fonksiyonlar ve Function Scope (kapsama alanı)
1. Deferring (erteleme)

## Bölüm 06: Durum Kontrolleri

1. `if`, `else`, `else if`
1. Short `if` declaration (kısa if tanımı)
1. `switch` ve `case` İfadeleri
1. Label, `break` ve `goto` İşlemleri

## Bölüm 07: Döngüler

1. `C` stili döngü
1. `range`
1. `break` ve `continue`
1. `for` Kullanımı
1. Label Kullanımı

## Bölüm 08: Interface

1. Tanımı
1. Tip Olarak **empty interface** ya da `any`
1. Tip Kontrol Meknizması
1. Davranış Olarak `interface`

## Bölüm 09: Error

1. `error` Nedir?
1. Custom Error Types (özelleştirilmiş error tipi oluşturmak)
1. Wrapping (sarmalama)
1. Unwrapping (sarmalı açma)
1. Error Tip Kontrolleri: `errors.Is` ve `errors.As`
1. Yaygın Pratikler
1. `panic` ve `recover`

## Bölüm 10: `nil`

1. `nil` Nedir?
1. Nerelerde ve Ne İçin Kullanınır?

## Bölüm 11: Generics

1. Nedir? Ne Amaçla Kullanılır
1. Generic Types
1. Generic Functions
1. Generic Interfaces

## Bölüm 12: Reflection

1. Ne İşe Yarar? Faydaları ve Zararları
1. Tip Kontrolleri
1. @wip

## Bölüm 12: JSON İle Çalışmak

1. `json:"TAG"`
1. Encoding (Marshal)
1. Decoding (Unmarshal)
1. Generic Interface
1. Reference Types
1. Streaming Encoders ve Decoders

## Bölüm 13: Test

1. Test Nedir? Neden Yazılır?
1. Test Nasıl Çalıştırılır
1. Examples ve `godoc` Nedir?
1. Race Detection Nedir?
1. SetUp ve TearDown Nedir?
1. Table Driven Test Nedir?
1. Testlerin Paralel Çalıştırılması?
1. Code Coverage Nedir?
1. Benchmarking ve Profiling Nedir?
1. Escape Analysis
1. Memory ve CPU Profiling Temelleri

## Bölüm 14: Concurrency

1. Nedir? Golang’in Concurrency Stratejisi Nedir?
1. Goroutine Nedir?
1. `go` Kelimesiyle Başlayan Anonim Fonksiyonlar
1. WaitGroup Nedir?
1. Mutex Nedir?
1. Channels
1. Context
1. Concurrency Pratikleri (yaygın kullanılan desenler)

## Bölüm 16: İsimlendirme Kuralları

1. Naming Conventions
1. Değişken İsimlendirmesi
1. Fonksiyon İsimlendirmesi
1. Method’ların İsimlendirmesi
1. Interface’lerin İsimlendirmesi
1. Paketlerin İsimlendirmesi

## Bölüm 17: Golang Paketi Geliştirmek

1. `golangci-linter` Kurulumu ve Konfigürasyonu
1. Paketi Dağıtmak / Paylaşmak
1. Go Modülü Anatomisi

## Bölüm 18: In-Memory Key-Value Store

1. `http` Paketini Kullanarak Rest-API Tasarlamak
1. Domain Driven Design prensibini Kullanmak
1. http server’ın Unit Testleri
1. Uygulamanın Docker Container’ından Çalıştırılması
1. GitHub Actions ile Linter/Checker Kullanımı

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
