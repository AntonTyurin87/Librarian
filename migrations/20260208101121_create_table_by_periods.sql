-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS by_periods (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    source_id INTEGER NOT NULL,
    period_id INTEGER NOT NULL,
    type TEXT NOT NULL,
    pages TEXT NOT NULL
);

-- Индексы для оптимизации работы с регионами
CREATE INDEX IF NOT EXISTS idx_period_id ON by_periods(period_id);
CREATE INDEX IF NOT EXISTS idx_type ON by_periods(type);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_type;
DROP INDEX IF EXISTS idx_period_id;
DROP TABLE IF EXISTS by_periods;
-- +goose StatementEnd