-- +goose Up
-- +goose StatementBegin
ALTER TABLE books DROP COLUMN file_format;
ALTER TABLE articles DROP COLUMN file_format;
ALTER TABLE fragments DROP COLUMN file_format;
ALTER TABLE photos DROP COLUMN file_format;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books ADD COLUMN file_format TEXT;
ALTER TABLE articles ADD COLUMN file_format TEXT;
ALTER TABLE fragments ADD COLUMN file_format TEXT;
ALTER TABLE photos ADD COLUMN file_format TEXT;
-- +goose StatementEnd
