# Dokumentasi: GET /api/form/

Endpoint: `GET /api/form/`

- Ringkasan: Mengambil daftar form/permintaan milik user yang sedang terautentikasi. Handler: `GetFormByUserID`.
- Proteksi: membutuhkan JWT. Sertakan header `Authorization: Bearer <token>`.

Request
- Method: `GET`
- URL: `/api/form/`
- Headers:
  - `Authorization: Bearer <JWT>` (wajib)

Response (sukses)
- Status: `200 OK`
- Body: JSON array dari objek form. Contoh struktur (field dapat berbeda sesuai DTO/models):

```json
[
  {
    "id": "f2f3e8a0-...",
    "id_user": "u1234-...",
    "status": "Pengisian Data",
    "created_at": "2026-01-12T08:00:00Z",
    "updated_at": "2026-01-12T08:00:00Z",
    "fields": {
      "nama": "Budi",
      "alamat": "Jl. Contoh 1"
    }
  }
]
```

Response (error umum)
- `401 Unauthorized`: token tidak ada/invalid.
- `403 Forbidden`: token tidak memenuhi syarat (jika ada middleware tambahan).
- `500 Internal Server Error`: kesalahan server/databse.

Contoh curl

```bash
TOKEN="$(cat /path/to/token.txt)"
curl -s -X GET http://localhost:8080/api/form/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Accept: application/json"
```

Catatan
- Gantilah `http://localhost:8080` dengan base URL server Anda.
- Untuk schema lengkap response lihat DTO/handler di `internal/domains/dto` dan `internal/domains/handler`.
