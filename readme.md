# 1cak API

![Go Version](https://img.shields.io/github/go-mod/go-version/Satr10/1cak-api)

API tidak resmi untuk mengambil data postingan dari [1cak.com](https://1cak.com). API ini melakukan scraping pada halaman web 1cak untuk menyediakan data dalam format JSON.

Dibangun menggunakan Go (Golang) dan dirancang untuk mudah di-deploy ke platform seperti Vercel.

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https://github.com/Satr10/1cak-api)

## ✨ Fitur

* Mengambil postingan dari 1cak.com
* Mendapatkan postingan populer dan acak
* Menyediakan data dalam format JSON yang terstruktur
* Endpoint API yang simpel dan mudah digunakan
* Konfigurasi siap untuk deployment di Vercel

## 🚀 Endpoint API

Ganti `<YOUR_DEPLOYED_URL>` dengan URL deployment Anda.

* **`GET /`**
  * Menampilkan informasi API dan cara penggunaan
* **`GET /api/`**
  * Mengambil meme acak (sama dengan `/api/random`)
* **`GET /api/random`**
  * Mengambil meme acak dari situs
* **`GET /api/popular`**
  * Mengambil meme populer/trending

### Struktur Respons Sukses (JSON)

```json
{
  "success": true,
  "posts": [
    {
      "id": "...", // ID Postingan di 1cak
      "url": "...", // URL lengkap ke postingan
      "title": "...", // Judul postingan
      "imageUrl": "...", // URL gambar postingan
      "votes": "...", // Jumlah vote (string)
      "nsfw": false // Boolean, true jika NSFW
    },
    // ... post lainnya
  ]
}
```

### Struktur Respons Gagal (JSON)

```json
{
  "success": false,
  "error": "Pesan error di sini"
}
```

## 🛠️ Setup Lokal

Untuk menjalankan API ini di mesin lokal Anda:

1. **Prasyarat:** Pastikan Anda sudah menginstal [Go](https://golang.org/doc/install)
2. **Clone Repositori:**
   ```bash
   git clone https://github.com/Satr10/1cak-api.git
   cd 1cak-api
   ```
3. **Jalankan Server:**
   ```bash
   go run main.go
   ```
   Secara default, server akan berjalan di port `5001`
4. **Akses API:** Buka browser atau gunakan `curl` untuk mengakses `http://localhost:5001/`

## ☁️ Deployment

Proyek ini sudah dikonfigurasi untuk deployment mudah ke [Vercel](https://vercel.com/).

* File `api/index.go` berfungsi sebagai entrypoint untuk serverless function Vercel
* File `vercel.json` berisi konfigurasi build dan routing untuk Vercel

Cukup impor repositori GitHub Anda ke Vercel, dan Vercel seharusnya secara otomatis mendeteksi konfigurasi Go dan men-deploy-nya.

## 💡 Contoh Penggunaan (dengan `curl`)

Ganti `<YOUR_DEPLOYED_URL>` dengan URL deployment Anda (atau `http://localhost:8080` jika berjalan lokal).

```bash
# Mendapatkan info API
curl <YOUR_DEPLOYED_URL>/

# Mendapatkan meme acak
curl <YOUR_DEPLOYED_URL>/api/random

# Mendapatkan meme populer
curl <YOUR_DEPLOYED_URL>/api/popular
```

## ⚠️ Disclaimer

* Ini adalah API tidak resmi dan bergantung sepenuhnya pada struktur HTML dari website `1cak.com`. Jika 1cak mengubah tata letak situsnya, API ini kemungkinan besar akan **berhenti berfungsi** sampai diperbarui.
* Gunakan API ini secara bertanggung jawab. Hormati ketentuan layanan `1cak.com`. Penggunaan berlebihan dapat menyebabkan IP Anda diblokir oleh 1cak.
* Pengembang tidak bertanggung jawab atas penggunaan API ini.

## 📜 Lisensi

Proyek ini dilisensikan di bawah [Lisensi MIT](LICENSE).