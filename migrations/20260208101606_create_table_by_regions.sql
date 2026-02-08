-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS by_regions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    source_id INTEGER NOT NULL,
    region_id INTEGER NOT NULL,
    type TEXT NOT NULL,
    pages TEXT NOT NULL
);

-- Индексы для оптимизации работы с регионами
CREATE INDEX IF NOT EXISTS idx_region_id ON by_regions(region_id);
CREATE INDEX IF NOT EXISTS idx_type ON by_regions(type);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_type;
DROP INDEX IF EXISTS idx_region_id;
DROP TABLE IF EXISTS by_regions;
-- +goose StatementEnd
