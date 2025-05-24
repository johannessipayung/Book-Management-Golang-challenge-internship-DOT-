# 📚 Book API Golang

## 📝 Deskripsi
Aplikasi backend sederhana untuk manajemen buku dan user authentication berbasis JWT.

## 🧱 Project Structure & Pattern Explanation

Project ini menggunakan **Layering Pattern** (Layered Architecture) yang membagi aplikasi menjadi beberapa lapisan dengan tanggung jawab berbeda. Tujuan penggunaan pattern ini adalah untuk:

- ⚙️ Memisahkan concerns (separation of concerns) agar tiap lapisan fokus pada fungsinya
- 🛠️ Meningkatkan maintainability sehingga kode mudah dipahami, dikembangkan, dan diperbaiki
- 🧩 Mendukung modularitas agar tiap layer bisa dikembangkan dan diuji secara independen
- 🧪 Memudahkan testing karena boundary antar layer jelas
- 👥 Mempermudah kolaborasi dalam tim
- 🚀 Mendukung skalabilitas aplikasi di masa depan

### 🗂️ Struktur Layer:
- `handler/` → 📥 Menerima HTTP request dan memberikan response (Controller)
- `service/` → 🧠 Menangani business logic dan validasi
- `repository/` → 🗃️ Berinteraksi langsung dengan database (CRUD operation)
- `model/` → 🧾 Struktur data (entities/DTO)
- `middleware/` → 🛡️ Middleware seperti JWT authentication, logging, dll
- `config/` → ⚙️ Konfigurasi dan setup koneksi database
- `router/` → 🌐 Setup routing endpoint menggunakan Gin

### ✅ Alasan Penggunaan Pattern:
- 🔍 Memudahkan testing karena tiap layer dapat diuji secara terpisah
- 🧱 Code lebih modular dan mudah dikelola
- 🏗️ Cocok untuk proyek skala menengah hingga besar
- 🧼 Penerapan praktik clean architecture

## 🧪 Testing

- ✅ Menggunakan `testify` untuk assertion dan mocking
- 🧾 Disediakan E2E test di `e2e/login_test.go` untuk memastikan alur login berjalan sesuai harapan

## 🧭 Arsitektur

Menggunakan **Layered Architecture Pattern** agar kode terstruktur dengan baik dan tanggung jawab dipisah antar lapisan:

- `handler/` → 📥 Mengelola input/output data dari dan ke client (sebagai controller)
- `service/` → 🧠 Menangani logika bisnis utama
- `repository/` → 💾 Bertanggung jawab untuk komunikasi langsung dengan database menggunakan GORM
- `model/` → 🧾 Representasi struktur data dan entitas (seperti User dan Book)
- `config/` → ⚙️ Konfigurasi aplikasi seperti koneksi ke database dan environment setup
- `middleware/` → 🛡️ Middleware khusus seperti autentikasi JWT yang memfilter request sebelum sampai ke handler
- `e2e/` → 🧪 Pengujian End-to-End untuk memastikan alur sistem bekerja sebagaimana mestinya
- `test/` → 🔬 Unit test untuk masing-masing komponen seperti service atau repository, agar fungsionalitas dapat diuji secara terisolasi

## 🔧 Teknologi
- 🚀 Gin (Web Framework)
- 🛠️ GORM (ORM)
- 🐘 PostgreSQL
- 🔐 JWT untuk autentikasi
- 📬 Postman untuk dokumentasi dan testing API

## ▶️ Jalankan
```bash
go run cmd/main.go
```

## 🧪 E2E Test
```bash
go test ./...
```

## 📬 Dokumentasi
Import file Postman Collection `postman/book-api.postman_collection.json`

## 🌱 Pengembangan Selanjutnya (Opsional)
Penggunaan DTO (Data Transfer Object) secara konsisten untuk memisahkan model domain dengan data yang dikirim dan diterima lewat API, sehingga menjaga keamanan dan fleksibilitas data yang diproses di tiap layer
## 👨‍💻 Author
Johannes Bastian Jasa Sipayung