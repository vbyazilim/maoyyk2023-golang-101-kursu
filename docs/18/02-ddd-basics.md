# Bölüm 18/02: In-Memory Key-Value Store

## Domain Driven Design (DDD)

Domain Driven Design (DDD), yazılım geliştirme sürecinde karmaşık iş
domain’lerini anlamak ve bu domain’leri daha iyi modellemek için kullanılan
bir yaklaşımı ifade eder. 

DDD, iş dünyasının gereksinimlerini, kullanıcıların ihtiyaçlarını ve yazılımın
tasarımını bir araya getirerek daha iyi bir anlayış ve etkili bir kod
oluşturmayı amaçlar. İşte Domain Driven Design’ın temel kavramları ve ana
prensipleri:

### Domain

İşinizin odaklandığı konu veya iş kolu. DDD’de, bu domain’i anlamak ve
modellemek temel amaçtır.

### Model

DDD, gerçek dünyadaki nesneleri ve ilişkileri yazılım nesnelerine dönüştürmeyi
içerir. Bu, karmaşık iş domain’lerini daha iyi anlamak ve yönetmek için
kullanılır. Veritabanındaki tablo’nun karşılığı gibi düşünülebilir.

### Servis

DDD’de, bazı işlemler modele ait olmayabilir. Bu durumlarda servisler
kullanılır. Örneğin, bir ödeme işlemi bir servis tarafından yönetilebilir.

### Repository (Storage)

Veritabanı ile etkileşimde bulunan bir bileşen. Repository, veritabanı
işlemlerini yönetir ve sınıf nesnelerini saklar.

---

Bizim uygulamamızda ise; gelen http isteği sırasıyla;

1. HTTP Handler Katmanı: gelen isteği alıp, validate edecek.
1. Servis Katmanı: Yapacağı işe göre ilgili servisi, istekten aldığı verileri
kullanarak çağıracak.
1. Storage Katmanı: Servis, ilgili storage’ları kullanarak veritabanı ile
konuşacak

sonra; yine sırasıyla geri dönüş başlayacak;

1. Storage -> Servis
1. Servis -> HTTP Handler
1. HTTP Hander -> İsteği atan istemci (client)

Her zaman, HTTP handler input’u alacak, kontrollerini yapacak, gerektiği gibi
gelmiş mi diye. Sonra görevi burada bitip, ilgili servisi çağıracak. Servis
esas yapılacak işten sorumlu. Eğer veritabanı ile işi varsa Storage ile
iletişim kuracak. Her şey yolunda giderse, Storage’dan aldıklarını HTTP
handler’a iletecek. HTTP handler yine her şey yolunda giderse client’a cevap
(response) dönecek ve akış tamamlanmış olacak.

---

## Kaynaklar

- https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/
- https://dev.to/stevensunflash/using-domain-driven-design-ddd-in-golang-3ee5
- https://medium.com/@gsigety/domain-driven-design-golang-kata-1-d76d01459806
