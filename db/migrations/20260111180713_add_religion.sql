-- +goose Up
-- +goose StatementBegin
ALTER TABLE profiles ADD COLUMN IF NOT EXISTS religion VARCHAR(255);
ALTER TABLE profiles ADD COLUMN IF NOT EXISTS work VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE profiles DROP COLUMN IF EXISTS religion;
ALTER TABLE profiles DROP COLUMN IF EXISTS work;
-- +goose StatementEnd
