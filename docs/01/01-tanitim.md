# Bölüm 01/01: Tanıtım

[Go][01]; statik tipler kullanan, derlenen yüksek seviyeli bir programlama
dilidir. Google çalışanlarından [Robert Griesemer][02], [Rob Pike][03], ve 
[Ken Thompson][04] tarafından geliştirilmiştir.

Bir rivayete göre; google aramalarında `go` çok jenerik kaldığı için; `golang`
olarak da isimlendirilmiştir. `Go` ya da `Golang` aslında aynı anlamda
kullanılmıştır.

Yazılım stili olarak (syntax) `C` diline benzese bile, güvenli hafıza
yönetimi ve kullanımına, işi biten hafızanın geri bırakılmasına (garbage collection),
yapısal tiplerine ve kendine özgü [CSP-style][05] (concurrency) eş-zamanlılık
yapısına kadar çok büyük farklar ve avantajlar sağlar.

YouTube [videosunda][06] tüm go ekibini görebilirsiniz.

**2007** yılında duyurulmuş, **2009** yılında da tüm dünyaya açılmıştır. Açık
Kaynak (open source) şeklinde geliştirilmeye devam ediyor.

Dili geliştirenlerin şöyle bir sözü var;

> The language is designed to build software services.

Yani dilin asıl amacı yazılım servisleri, hatta internet servisleri geliştirmek.
**Cloud Native** tanımıyla örtüşüyor, yani cloud’ın dili: Go!

Ana presibleri;

- Basitlik
- Açıkça tanımlanmış talimatlar dizisi
- Statik tipler (tanımlı tipler)
- Üzerinde derlendiği işletim sisteminin doğal dili ne ise o dile derlenme,
  Bu sayede Java ve benzeri dillerdeki sanal makine (virtual machine) ihtiyacı yok
- Nesne yönelimli değil (oop) ama tipler birbir içine geçebiliyor, miras kavramı yok
- Interface mantığı
- Fonksiyonlar hem tip hem de argüman olarak kullanılabilir
- [Orthogonality][07]; yani bir fonksiyon ya da işlem başka bir şeyi bozmadan değişebiliyor
- Gömülü olarak gelen eşzamanlılık ilkelleri: Goroutines ve Channels
- Doküman ve test öncelikli yaklaşım


[01]: https://en.wikipedia.org/wiki/Go_(programming_language)
[02]: https://en.wikipedia.org/wiki/Robert_Griesemer
[03]: https://en.wikipedia.org/wiki/Rob_Pike
[04]: https://en.wikipedia.org/wiki/Ken_Thompson
[05]: https://en.wikipedia.org/wiki/Communicating_sequential_processes
[06]: https://www.youtube.com/watch?v=sln-gJaURzk
[07]: https://en.wikipedia.org/wiki/Orthogonality_(programming)
