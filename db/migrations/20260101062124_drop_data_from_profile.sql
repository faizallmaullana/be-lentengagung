-- +goose Up
-- +goose StatementBegin
ALTER TABLE profiles
	DROP COLUMN IF EXISTS religion,
	DROP COLUMN IF EXISTS work;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE profiles
	ADD COLUMN IF NOT EXISTS religion text,
	ADD COLUMN IF NOT EXISTS work text;
-- +goose StatementEnd
