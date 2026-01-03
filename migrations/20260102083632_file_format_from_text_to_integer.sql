-- +goose Up
-- +goose StatementBegin
ALTER TABLE books ADD COLUMN format INTEGER;
ALTER TABLE articles ADD COLUMN format INTEGER;
ALTER TABLE fragments ADD COLUMN format INTEGER;
ALTER TABLE photos ADD COLUMN format INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books ADD COLUMN file_format TEXT;
ALTER TABLE articles ADD COLUMN file_format TEXT;
ALTER TABLE fragments ADD COLUMN file_format TEXT;
ALTER TABLE photos ADD COLUMN file_format TEXT;
-- +goose StatementEnd
