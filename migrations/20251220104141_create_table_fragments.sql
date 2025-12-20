-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS fragments (
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

-- Индексы для оптимизации поиска фрагментов
CREATE INDEX IF NOT EXISTS idx_fragments_regions ON fragments(regions);
CREATE INDEX IF NOT EXISTS idx_fragments_time_periods ON fragments(time_periods);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_fragments_time_periods;
DROP INDEX IF EXISTS idx_fragments_regions;
DROP TABLE IF EXISTS fragments;
-- +goose StatementEnd
