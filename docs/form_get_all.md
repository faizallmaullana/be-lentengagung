# Dokumentasi: GET /api/form/all

Endpoint: `GET /api/form/all`

- Ringkasan: Mengambil daftar form/permintaan. Jika pemanggil memiliki role `admin` atau `superadmin`, endpoint mengembalikan semua entri. Jika bukan admin, mengembalikan hanya entri yang dimiliki oleh user pemanggil (filtered by `id_user`).
- Proteksi: membutuhkan JWT. Sertakan header `Authorization: Bearer <token>`.

Request
- Method: `GET`
- URL: `/api/form/all`
- Headers:
  - `Authorization: Bearer <JWT>` (wajib)

Behavior
- Jika token valid dan user.role == `admin` || `superadmin` → kembalikan semua form.
- Jika token valid dan role lain → kembalikan hanya form dengan `id_user` sesuai user pemanggil.

Success Response (admin — contoh)
- Status: `200 OK`
- Body: JSON array — semua form yang tersimpan.

```json
[
  {
    "id": "f2f3e8a0-...",
    "id_user": "u-admin-...",
    "status": "Selesai",
    "timestamp": "2026-01-12T08:00:00Z",
    "fields": {"nama":"Admin","keterangan":"..."}
  },
  {
    "id": "a7d9b1c3-...",
    "id_user": "u1234-...",
    "status": "Pengisian Data",
    "timestamp": "2026-01-11T10:00:00Z",
    "fields": {"nama":"Budi","keterangan":"..."}
  }
]
```

Success Response (non-admin — contoh)
- Status: `200 OK`
- Body: JSON array — hanya entri milik user.

```json
[
  {
    "id": "a7d9b1c3-...",
    "id_user": "u1234-...",
    "status": "Pengisian Data",
    "timestamp": "2026-01-11T10:00:00Z",
    "fields": {"nama":"Budi","keterangan":"..."}
  }
]
```

Error responses
- `401 Unauthorized`: token tidak disertakan atau tidak valid.
- `403 Forbidden`: (opsional) middleware tambahan menolak akses.
- `500 Internal Server Error`: kesalahan server atau database.

Contoh curl

```bash
TOKEN="$TOKEN"
# Admin (mengambil semua)
curl -X GET http://localhost:8080/api/form/all \
  -H "Authorization: Bearer $TOKEN" \
  -H "Accept: application/json"

# Non-admin (mengambil hanya milik user)
curl -X GET http://localhost:8080/api/form/all \
  -H "Authorization: Bearer $TOKEN" \
  -H "Accept: application/json"
```

Notes
- Gantilah base URL (`http://localhost:8080`) sesuai konfigurasi server.
- Untuk schema field yang tepat, lihat DTO/handler di `internal/domains/dto` dan `internal/domains/handler`.
