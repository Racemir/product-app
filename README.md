# product-app
# Product Management API

Bu proje, **Go (Golang)** ve **Echo framework** kullanılarak geliştirilmiş, katmanlı mimariye (Layered Architecture) sahip bir ürün yönetim API'sidir. Proje, veritabanı olarak **PostgreSQL** kullanır ve geliştirme süreçlerini hızlandırmak için **Docker** ile entegre çalışır.

## 🚀 Teknolojiler ve Araçlar

* **Backend:** Go (Golang)
* **Web Framework:** [Echo](https://echo.labstack.com/)
* **Database:** PostgreSQL
* **Containerization:** Docker & Docker Compose
* **Testing:** Testify (Unit & Integration Testing)

## 🏗️ Mimari Yapı

Proje, bağımlılıkların yönetimi ve test edilebilirliği artırmak için katmanlı bir yapıda kurgulanmıştır:

* **Controller:** HTTP isteklerini karşılar, veriyi valide eder ve Response döner.
* **Service:** İş mantığının (Business Logic) bulunduğu ana katman. (Örn: %70 indirim sınırı kontrolü).
* **Persistence:** Veritabanı (Repository) işlemlerinin yürütüldüğü katman.
* **Domain:** Projenin temel nesneleri (Entity) ve arayüzleri (Interface).
* **Common:** Veritabanı bağlantı ayarları ve konfigürasyon yöneticisi.

## 📁 Dosya Yapısı

```text
product-app/
├── common/             # Konfigürasyon ve DB bağlantı yönetimi
├── controller/         # Request/Response modelleri ve Handler'lar
├── domain/             # Core Entity (Product) modelleri
├── persistence/        # Veritabanı (SQL) katmanı
├── service/            # İş mantığı ve Servis modelleri
├── test/
│   ├── infrastructure/ # Gerçek DB ile yapılan Entegrasyon testleri
│   ├── scripts/        # Docker/Postgres kurulum script'leri (test_db.sh)
│   └── service/        # Fake Repo ile yapılan hızlı Birim (Unit) testleri
├── main.go             # Uygulama giriş noktası
└── go.mod              # Bağımlılık yönetimi

🛠️ Başlangıç
1. Veritabanını Ayağa Kaldırma (Docker)
Proje içinde bulunan bash script'i ile PostgreSQL konteynırını ve gerekli tabloları otomatik olarak oluşturabilirsiniz:

Bash
bash test/scripts/test_db.sh

2. Uygulamayı Çalıştırma
Bağımlılıkları yükleyin ve sunucuyu başlatın:

Bash
go mod tidy
go run main.go

Sunucu varsayılan olarak http://localhost:8080 adresinde çalışmaya başlayacaktır.
