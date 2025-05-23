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

## `/api/` dan `/api/random`
```json
{
  "meme": {
    "PostURL": "http://1cak.com/2998617",
    "Title": "SeFruit Pembicaraan(Vol.1100)",
    "ImageUrl": "https://1cak.com//posts/17b0e8a6d2e009a0fdfef58675ea836f_t.jpg",
    "Like": "1284",
    "Uploader": "bagbamwtwo",
    "UploadTime": "Thu, 24 Aug 2023 13:45:25 +0700"
  },
  "status": "Success"
}
```
## `/api/popular`

```json
{
  "memes": [
    {
      "PostURL": "test",
      "Title": "Gelar Tikar Di Air",
      "ImageUrl": "https://1cak.com/3047786",
      "Like": "1045",
      "Uploader": "aribloong",
      "UploadTime": "Tue, 15 Apr 2025 08:49:19 +0700"
    },
    {
      "PostURL": "test",
      "Title": "Saya Aslinya Dua Orang",
      "ImageUrl": "https://1cak.com/3047782",
      "Like": "1010",
      "Uploader": "panjias",
      "UploadTime": "Tue, 15 Apr 2025 08:05:30 +0700"
    },
    {
      "PostURL": "test",
      "Title": "Ivan Estrada Be Like",
      "ImageUrl": "https://1cak.com/3047781",
      "Like": "1091",
      "Uploader": "panjias",
      "UploadTime": "Tue, 15 Apr 2025 07:47:50 +0700"
    },
    {
      "PostURL": "test",
      "Title": "Walau Besi Suramadu Suka Diprintilin, Jembatan Ini Tetep Kokoh Dan Siap Menyediakan Jalannya Bagi Warga Surabaya Maupun Madura Yang Mau Melintas",
      "ImageUrl": "https://1cak.com/3047780",
      "Like": "1168",
      "Uploader": "dewa_kipas",
      "UploadTime": "Tue, 15 Apr 2025 07:42:31 +0700"
    }
  ],
  "status": "Success"
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