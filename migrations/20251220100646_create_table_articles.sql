-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS articles (
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

-- Индексы для оптимизации поиска статей
CREATE INDEX IF NOT EXISTS idx_articles_author_ru ON articles(author_ru);
CREATE INDEX IF NOT EXISTS idx_articles_regions ON articles(regions);
CREATE INDEX IF NOT EXISTS idx_articles_time_periods ON articles(time_periods);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_articles_time_periods;
DROP INDEX IF EXISTS idx_articles_regions;
DROP INDEX IF EXISTS idx_articles_author_ru;
DROP TABLE IF EXISTS articles;
-- +goose StatementEnd
