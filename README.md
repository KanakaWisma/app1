# CRUD Mahasiswa

Aplikasi CRUD Mahasiswa menggunakan:

- Frontend: React + Vite
- Backend: Golang + Gin
- Database: PostgreSQL
- ORM: GORM

## Cara Menjalankan Project

### 1. Jalankan PostgreSQL

Pastikan PostgreSQL sudah aktif dan database sudah dibuat.

Contoh database:

CREATE DATABASE db_mahasiswa;

---

### 2. Jalankan Backend

Masuk ke folder backend:

cd backend

Install dependency:

go mod tidy

Jalankan server:

go run main.go

Backend akan berjalan di:

http://localhost:8080

---

### 3. Jalankan Frontend

Masuk ke folder frontend:

cd frontend

Install dependency:

npm install

Jalankan aplikasi:

npm run dev

Frontend akan berjalan di:

http://localhost:5173

---

## Fitur

- Menampilkan data mahasiswa
- Menambahkan data mahasiswa
- Mengubah data mahasiswa
- Menghapus data mahasiswa
- Menampilkan data jurusan
- Relasi Mahasiswa dan Jurusan menggunakan GORM Preload

---

## Struktur Database

### Tabel Jurusan

| Field | Tipe |
|-------|------|
| IDJurusan | Integer |
| NamaJurusan | Varchar |
| Fakultas | Varchar |
| Jenjang | Varchar |

### Tabel Mahasiswa

| Field | Tipe |
|-------|------|
| ID | Integer |
| Nama | Varchar |
| Umur | Integer |
| NIM | Varchar |
| TglLahir | Date |
| Alamat | Text |
| IDJurusan | Integer |

---

## API Endpoint

### Mahasiswa

GET /mahasiswa

POST /mahasiswa

PUT /mahasiswa/:id

DELETE /mahasiswa/:id

### Jurusan

GET /jurusan

---

## Author

Nama: Kanaka Wisma

Program Studi: Informatika