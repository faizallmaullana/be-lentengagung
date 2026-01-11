-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chat_main (
	id uuid PRIMARY KEY,
	id_register_pernyataan uuid REFERENCES register_pernyataan(id) ON DELETE CASCADE,
	chat VARCHAR(2048),
	id_user uuid REFERENCES users(id) ON DELETE SET NULL,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_document (
	id uuid PRIMARY KEY,
	id_document uuid REFERENCES dokumen(id) ON DELETE CASCADE,
	chat VARCHAR(2048),
	id_user uuid REFERENCES users(id) ON DELETE SET NULL,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_pewaris (
	id uuid PRIMARY KEY,
	id_pewaris uuid REFERENCES pewaris(id) ON DELETE CASCADE,
	chat VARCHAR(2048),
	id_user uuid REFERENCES users(id) ON DELETE SET NULL,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_pasangan_pewaris (
	id uuid PRIMARY KEY,
	id_pasangan_pewaris uuid REFERENCES pasangan_pewaris(id) ON DELETE CASCADE,
	chat VARCHAR(2048),
	id_user uuid REFERENCES users(id) ON DELETE SET NULL,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_ahli_waris (
	id uuid PRIMARY KEY,
	id_ahli_waris uuid REFERENCES ahli_waris(id) ON DELETE CASCADE,
	chat VARCHAR(2048),
	id_user uuid REFERENCES users(id) ON DELETE SET NULL,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_pasangan_ahli_waris (
	id uuid PRIMARY KEY,
	id_pasangan_ahli_waris uuid REFERENCES pasangan_ahli_waris(id) ON DELETE CASCADE,
	chat VARCHAR(2048),
	id_user uuid REFERENCES users(id) ON DELETE SET NULL,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_cucu (
	id uuid PRIMARY KEY,
	id_cucu uuid REFERENCES cucu(id) ON DELETE CASCADE,
	chat VARCHAR(2048),
	id_user uuid REFERENCES users(id) ON DELETE SET NULL,
	timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chat_cucu;
DROP TABLE IF EXISTS chat_pasangan_ahli_waris;
DROP TABLE IF EXISTS chat_ahli_waris;
DROP TABLE IF EXISTS chat_pasangan_pewaris;
DROP TABLE IF EXISTS chat_pewaris;
DROP TABLE IF EXISTS chat_document;
DROP TABLE IF EXISTS chat_main;
-- +goose StatementEnd
