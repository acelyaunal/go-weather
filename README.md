
# Go Weather CLI 🌤️

Merhaba!  
Bu proje, Go dilinde geliştirdiğim basit bir **hava durumu komut satırı uygulamasıdır**.  
Şehir adını yazarak anlık sıcaklık, nem, rüzgar hızı ve önümüzdeki birkaç günün tahminini terminalde görebilirsiniz.

---

## 🚀 Özellikler
- Şehir adına göre anlık hava durumu bilgisi  
- Günlük tahmin (sıcaklık, yağış, gün doğumu/batımı)  
- JSON formatında çıktı desteği (`-json` parametresi)  
- Türkçe dil desteği  
- Çoklu ülke desteği (“Berlin, DE” veya “Berlin, Almanya” yazılabilir)

---

## 🧩 Kullanım
Projeyi çalıştırmak için Go 1.21+ kurulu olmalı.

```bash
git clone https://github.com/acelyaunal/go-weather.git
cd go-weather
go run . "Istanbul"
````

**Örnekler:**

```bash
go run . "Istanbul"
go run . "Sinop"
go run . -days 5 "Berlin, DE"
go run . -json "Ankara"
```

---

## 📦 Derleme

İkili dosya oluşturmak için:

```bash
go build -o weather
./weather "Istanbul"
```

---

## 🛠️ Yapılacaklar

* Mini HTTP servis ( `/weather?city=Istanbul` )
* Renkli terminal çıktısı
* “-save” parametresiyle geçmiş kaydı

---

## 👩🏻‍💻 Hakkımda

Ben **Açelya Ünal**, yazılım mühendisi olarak backend ve full-stack geliştirme üzerine çalışıyorum.
Go diliyle ilk projemi paylaşıyorum.

