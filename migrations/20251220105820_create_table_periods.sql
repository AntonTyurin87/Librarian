-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS periods (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    century INTEGER NOT NULL,
    era TEXT NOT NULL,
    description TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS periods;
-- +goose StatementEnd
