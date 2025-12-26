-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id uuid PRIMARY KEY,
	email TEXT NOT NULL UNIQUE,
	password_hash TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	approved_at TIMESTAMPTZ,
	is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS profiles (
	id uuid PRIMARY KEY,
	user_id uuid REFERENCES users(id) ON DELETE CASCADE,
	nik TEXT NOT NULL UNIQUE,
	phone TEXT,
	religion TEXT,
	address TEXT,
	work TEXT,
	name TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profiles;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
