-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pewaris (
	id uuid PRIMARY KEY,
	name VARCHAR(255),
	nik VARCHAR(255),
	phone VARCHAR(255),
	rt VARCHAR(255),
	rw VARCHAR(255),
	kelurahan VARCHAR(255),
	kecamatan VARCHAR(255),
	kabupaten VARCHAR(255),
	province VARCHAR(255),
	address VARCHAR(255),
	religion VARCHAR(255),
	work VARCHAR(255),
	date_of_birth VARCHAR(255),
	gender VARCHAR(255),
	blood_type VARCHAR(255),
	id_lampiran VARCHAR(255),
	no_akta_kematian VARCHAR(255),
	keterangan_kematian VARCHAR(255),
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS pasangan_pewaris (
	id uuid PRIMARY KEY,
	name VARCHAR(255),
	nik VARCHAR(255),
	phone VARCHAR(255),
	rt VARCHAR(255),
	rw VARCHAR(255),
	kelurahan VARCHAR(255),
	kecamatan VARCHAR(255),
	kabupaten VARCHAR(255),
	province VARCHAR(255),
	address VARCHAR(255),
	religion VARCHAR(255),
	work VARCHAR(255),
	date_of_birth VARCHAR(255),
	gender VARCHAR(255),
	blood_type VARCHAR(255),
	id_lampiran VARCHAR(255),
	urutan_pasangan SMALLINT,
	is_dead BOOLEAN,
	no_akta_kematian VARCHAR(255),
	keterangan_kematian VARCHAR(255),
	id_pewaris uuid REFERENCES pewaris(id) ON DELETE CASCADE,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS ahli_waris (
	id uuid PRIMARY KEY,
	name VARCHAR(255),
	nik VARCHAR(255),
	phone VARCHAR(255),
	rt VARCHAR(255),
	rw VARCHAR(255),
	kelurahan VARCHAR(255),
	kecamatan VARCHAR(255),
	kabupaten VARCHAR(255),
	province VARCHAR(255),
	address VARCHAR(255),
	religion VARCHAR(255),
	work VARCHAR(255),
	date_of_birth VARCHAR(255),
	gender VARCHAR(255),
	keterangan_kematian VARCHAR(255),
	no_akta_kematian VARCHAR(255),
	blood_type VARCHAR(255),
	id_lampiran VARCHAR(255),
	urutan_ahli_waris SMALLINT,
	is_dead BOOLEAN,
	id_pasangan_pewaris uuid REFERENCES pasangan_pewaris(id) ON DELETE CASCADE,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS pasangan_ahli_waris (
	id uuid PRIMARY KEY,
	name VARCHAR(255),
	nik VARCHAR(255),
	phone VARCHAR(255),
	rt VARCHAR(255),
	rw VARCHAR(255),
	kelurahan VARCHAR(255),
	kecamatan VARCHAR(255),
	kabupaten VARCHAR(255),
	province VARCHAR(255),
	address VARCHAR(255),
	religion VARCHAR(255),
	work VARCHAR(255),
	date_of_birth VARCHAR(255),
	gender VARCHAR(255),
	blood_type VARCHAR(255),
	id_lampiran VARCHAR(255),
	urutan_pasangan SMALLINT,
	is_dead BOOLEAN,
	no_akta_kematian VARCHAR(255),
	keterangan_kematian VARCHAR(255),
	id_ahli_waris uuid REFERENCES ahli_waris(id) ON DELETE CASCADE,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS cucu (
	id uuid PRIMARY KEY,
	name VARCHAR(255),
	nik VARCHAR(255),
	phone VARCHAR(255),
	rt VARCHAR(255),
	rw VARCHAR(255),
	kelurahan VARCHAR(255),
	kecamatan VARCHAR(255),
	kabupaten VARCHAR(255),
	province VARCHAR(255),
	address VARCHAR(255),
	religion VARCHAR(255),
	work VARCHAR(255),
	date_of_birth VARCHAR(255),
	gender VARCHAR(255),
	blood_type VARCHAR(255),
	id_lampiran VARCHAR(255),
	urutan_pasangan SMALLINT,
	is_dead BOOLEAN,
	no_akta_kematian VARCHAR(255),
	keterangan_kematian VARCHAR(255),
	id_pasangan_ahli_waris uuid REFERENCES pasangan_ahli_waris(id) ON DELETE CASCADE,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS register_pernyataan (
	id uuid PRIMARY KEY,
	kode_registrasi VARCHAR(255),
	status VARCHAR(255),
	id_user uuid REFERENCES users(id) ON DELETE SET NULL,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS dokumen (
	id uuid PRIMARY KEY,
	name VARCHAR(255),
	file_path VARCHAR(255),
    status VARCHAR(255),
	id_register_pernyataan uuid REFERENCES register_pernyataan(id) ON DELETE CASCADE,
    approved_at TIMESTAMPTZ,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS file_uploads (
	id VARCHAR(255) PRIMARY KEY,
	file_name VARCHAR(255),
	file_path VARCHAR(255),
	status VARCHAR(255),
	approved_at TIMESTAMPTZ,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS file_uploads;
DROP TABLE IF EXISTS dokumen;
DROP TABLE IF EXISTS register_pernyataan;
DROP TABLE IF EXISTS cucu;
DROP TABLE IF EXISTS pasangan_ahli_waris;
DROP TABLE IF EXISTS ahli_waris;
DROP TABLE IF EXISTS pasangan_pewaris;
DROP TABLE IF EXISTS pewaris;
-- +goose StatementEnd
