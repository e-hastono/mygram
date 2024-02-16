# MyGram - Emmanoel Pratama Putra Hastono

BFLP Bootcamp IT - Final Project Golang

MyGram adalah sebuah aplikasi untuk penyimpanan foto serta komentar foto orang lain.

## Pengantar

### Direktori

```
├── controllers
    ├── comment_c_ontroller.go
    ├── photo_c_ontroller.go
    ├── socialmedia_controller.go
    └── user_controller.go
├── database
    └── database.go
├── entities
    └── entities.go
├── helpers
    ├── token.go
    └── validation.go
├── middlewares
    ├── authentication.go
    └── authorization.go
├── models
    ├── comment.go
    ├── photo.go
    ├── socialmedia.go
    └── user.go
├── routers
    └── router.go
├── .env
├── go.mod
├── go.sum
└── main.go
```

### Daftar Endpoint

1. User : - Register [POST] - Login [POST]
2. Photo : - GetAll [GET] - GetOne [GET] - CreatePhoto [POST] - UpdatePhoto [PUT] - DeletePhoto [DELETE]
3. Comment : - GetAll [GET] - GetOne [GET] - CreateComment [POST] - UpdateComment [PUT] - DeleteComment [DELETE]
4. Social Media : - GetAll [GET] - GetOne [GET] - CreateSocialMedia [POST] - UpdateSocialMedia [PUT] - DeleteSocialMedia [DELETE]

### Instalasi

Pastikan Anda memiliki [Go](https://go.dev/) dan PostgreSQL terpasang pada perangkat Anda.

1. Clone repository ini

```bash
git clone https://github.com/e-hastono/mygram
```

2. Masuk ke dalam folder repository melalui terminal

```bash
cd mygram
```

3. Jalankan perintan untuk menginstalasi dependency repository

```bash
go get .
```

4. Salin file `.env.example` dan ganti namanya menjadi `.env`

```
cp .env.example .env
```

5. Ubah variabel dalam `.env` sesuai dengan pengaturan Anda:

- DB_HOST: alamat host database Postgres
- DB_USER: nama user database Postgres
- DB_PASS: password database Postgres
- DB_NAME: nama database Postgres
- DB_PORT: alamat port database Postgres
- API_SECRET: frasa yang digunakan untuk encoding token JWT
- TOKEN_HOUR_LIFESPAN: masa berlaku token JWT dalam jam

6. Jalankan repository

```bash
go run main.go
```

API dapat diakses di http://localhost:8080

## Demonstrasi

Demonstrasi repository ini dapat diakses di https://mygram-production-282f.up.railway.app/
