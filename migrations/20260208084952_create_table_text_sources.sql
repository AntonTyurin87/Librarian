-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS text_sources (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL UNIQUE,
    type TEXT NOT NULL,
    name_ru TEXT NOT NULL,
    name_eng TEXT,
    author_ru TEXT NOT NULL,
    year INTEGER NOT NULL,
    description TEXT,
    place_url TEXT,
    from_url TEXT,
    file_name TEXT,
    file_format TEXT,
    created_at TEXT,
    isAvailable BOOL NOT NULL
);

-- Индексы для оптимизации работы с регионами
CREATE INDEX IF NOT EXISTS idx_name_ru ON regions(name_ru);
CREATE INDEX IF NOT EXISTS idx_type ON text_sources(type);
CREATE INDEX IF NOT EXISTS idx_file_format ON text_sources(file_format);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_name_ru;
DROP INDEX IF EXISTS idx_type;
DROP INDEX IF EXISTS idx_idx_file_format;
DROP TABLE IF EXISTS text_sources;
-- +goose StatementEnd

