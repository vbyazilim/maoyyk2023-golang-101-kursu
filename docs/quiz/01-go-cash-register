# Go Cash Regist Assignment

## Amaç

Go dilinde basit bir kasa kayıt programı oluşturun. Bu program, bir liste içindeki ürünleri işleyebilmeli ve olası indirimleri uygulayarak toplam tutarı hesaplayabilmelidir.

## Gereklilikler

### Tipler ve Değişkenler

**Item** adında bir yapı (struct) tanımlayın. Bu yapıda **Name** (İsim), **Price** (Fiyat) ve **Discount** (İndirim) alanları olsun. İşlenecek ürünleri tutmak için **Item** türünden bir **slice** oluşturun.

### Fonksiyonlar

**calculatePrice(item Item) float64:** Bu fonksiyon, bir ürünü parametre olarak alır ve indirim uygulandıktan sonra fiyatını döndürür.

**totalPrice(items []Item) float64:** Bu fonksiyon, kesitteki tüm ürünlerin toplam fiyatını hesaplayarak döndürür.

### Döngüler ve Koşullu İfadeler

**calculatePrice** fonksiyonunda, bir ürünün indirimli olup olmadığını kontrol etmek için bir if ifadesi kullanın ve varsa indirimi uygulayın.

**totalPrice** fonksiyonunda, bir döngü kullanarak kesit içindeki ürünleri dolaşın ve fiyatları toplayın.

### Arayüzler

**Describable** adında bir **interface** tanımlayın ve içinde **Description() string** adında bir method olsun.
**Item** yapısı için Description fonksiyonunu **receiver** olarak ekleyin; bu yöntem, **"Ad - Fiyat (Eğer indirim varsa indirimli fiyat)"** formatında bir metin döndürmelidir.

### Örnek Çıktı

```text
Elma - 0.75 TL (10% indirimle 0.68 TL)
Portakal - 1.00 TL
Toplam Fiyat: 1.68 TL
```