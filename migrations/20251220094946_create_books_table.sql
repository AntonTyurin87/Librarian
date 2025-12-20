-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name_ru TEXT NOT NULL,
    name_native TEXT,
    author_ru TEXT,
    year INTEGER,
    regions TEXT,
    time_periods TEXT,
    description TEXT,
    file_format TEXT
);

-- Индексы для оптимизации поиска
CREATE INDEX IF NOT EXISTS idx_books_name_ru ON books(name_ru);
CREATE INDEX IF NOT EXISTS idx_books_author_ru ON books(author_ru);
CREATE INDEX IF NOT EXISTS idx_books_regions ON books(regions);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_books_time_periods;
DROP INDEX IF EXISTS idx_books_regions;
DROP INDEX IF EXISTS idx_books_author_ru;
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
