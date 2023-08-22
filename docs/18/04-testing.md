# Bölüm 18/04: In-Memory Key-Value Store

## Test

Go, **test first** bir dil. Yani testler hayati derecede önemli. Bir projenin
ne kadar iyi test coverage’ı olursa, o proje / kütüphane / servis o kadar
sağlam çalışır anlamına gelir.

Kabaca, yazdığımız her satır kod, fonksiyon, metot, yani her şey test
edilebilir şeylerdir. Eğer yazdığınız kodu test edemiyorsanız, o zaman bir
sıkıntı var demektir. Bir şeyleri hatalı yapmış ya da atlamışsınızdır.

Nelerin testlerini yapmamız iyi olur?

- Ek paket yaptık mı? (`kverror`)
- Storage katmanı (tüm metotları)
- Service katmanı (tüm metotları)
- HTTP Handler katmanı (tüm metotları)

İyi bir test coverage yüzdesi ~ `%80` civarındadır. Yani yazılan kodun en az
`%80`’i cover edilmişse bu iş **OK**’dir. (%80 - %20 yaklaşımı) Coverage ne
kadar yüksek olursa kendimizi o kadar güvende hissederiz.

---
