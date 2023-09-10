# Bölüm 17/04: Golang Paketi Geliştirmek

## Paketi Dağıtmak / Paylaşmak

Şimdi paketi GitHub’a göndermeden önce açık kaynak projelerin olmazsa
olmazlarını ekleyelim, hemen [örnek projeden][01] gereken dosyaları alalım:

- `README.md`
- `LICENSE`
- `CODE_OF_CONDUCT.md`

## `README.md`

İyi bir `README` dosyasında;

- Projenin adı ve kısa tanımı bulunmalı
- Nasıl kurulumu yapılır?
- Nasıl kullanılır? Örnek kod parçaları
- İlave komutlar (Makefile, Rakefile) varsa açıklaması
- Katkı sağlayanların listesi
- Nasıl katkı yapılacağı bilgisi
- Lisans
- `CODE_OF_CONDUCT` yani **Katkıcı Ahdi Topluluk Sözleşmesi**

olsa tadından yenmez.

Ek olarak bu bir go projesi olduğu için, GitHub action’ları kullanarak,
linter/checker ve build işlemlerini otomatize edebiliriz. **img.shields.io**
kullanarak README dosyasına **badge**’ler (version, ci/cd bilgileri gibi...)
ekleyebiliriz.

Tüm bu işlemleri yaptıktan sonra projemizin GitHub linkini internet ortamında
yayabiliriz. Repo’muz **public** yani herkese açık olduğu için `go get` ile
kolayca kurulup kullanılabilir durumda.

Eğer repo **private** olsaydı, sadece repo’ya erişebilenler `go get`
yapabilecekti.

[01]: https://github.com/vigo/stringutils-demo
