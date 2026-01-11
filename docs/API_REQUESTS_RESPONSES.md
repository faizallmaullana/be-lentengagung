# KTP-OCR API — Requests & Responses

## Overview
- Service entry: `app.py`
- Response model: `ktpocr/form.py`

## GET /
- Description: Health/status check
- Request: none
- Response (200):

```json
{
  "service": "KTP-OCR",
  "status": "ok"
}
```

## POST /extract
- Description: Upload an image (KTP) to extract fields via OCR.
- Method: `POST`
- Content-Type: `multipart/form-data`
- Form field: `file` — image file (jpg, png, etc.)

### Curl example

```bash
curl -X POST http://localhost:5000/extract -F "file=@/path/to/ktp.jpg"
```

### Success response (200)
Returns a JSON object with the following string fields (may be empty when OCR fails to detect a value):

- `nik` — NIK / identification number
- `nama` — full name
- `tempat_lahir` — place of birth
- `tanggal_lahir` — birth date (e.g. `DD-MM-YYYY` when parsed)
- `jenis_kelamin` — gender
- `golongan_darah` — blood type (O/A/B/AB or `-`)
- `alamat` — street address
- `rt` — RT
- `rw` — RW
- `kelurahan_atau_desa` — village/subdistrict
- `kecamatan` — district
- `agama` — religion
- `status_perkawinan` — marital status
- `pekerjaan` — occupation
- `kewarganegaraan` — nationality

Example success payload:

```json
{
  "nik": "3203010101010001",
  "nama": "BUDI SUKAMTO",
  "tempat_lahir": "BANDUNG",
  "tanggal_lahir": "01-01-1990",
  "jenis_kelamin": "LAKI-LAKI",
  "golongan_darah": "O",
  "alamat": "JL. MERDEKA NO. 1",
  "rt": "001",
  "rw": "002",
  "kelurahan_atau_desa": "CIBIRU",
  "kecamatan": "CIMAHI",
  "agama": "ISLAM",
  "status_perkawinan": "BELUM KAWIN",
  "pekerjaan": "SWASTA",
  "kewarganegaraan": "WNI"
}
```

### Error responses
- `400` — missing file: `{"error": "no file provided"}`
- `400` — empty filename: `{"error": "empty filename"}`
- `500` — processing error: `{"error": "<exception message>"}`

## Notes
- Uploaded files are saved to a temporary directory and removed after processing.
- OCR accuracy depends on image quality; some fields may be partially extracted or empty.
- See `app.py` for exact status codes and error messages.
