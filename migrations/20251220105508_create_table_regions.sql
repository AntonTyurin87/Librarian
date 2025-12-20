-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS regions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name_ru TEXT NOT NULL,
    description TEXT
);

-- Индексы для оптимизации работы с регионами
CREATE INDEX IF NOT EXISTS idx_regions_name_ru ON regions(name_ru);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_regions_name_ru;
DROP TABLE IF EXISTS regions;
-- +goose StatementEnd
