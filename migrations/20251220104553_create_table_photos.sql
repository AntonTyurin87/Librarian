-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS photos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER,
    name_ru TEXT NOT NULL,
    name_native TEXT,
    filming_location TEXT,
    regions TEXT,
    time_periods TEXT,
    description TEXT,
    file_format TEXT
);

-- Индексы для оптимизации работы с фотографиями
CREATE INDEX IF NOT EXISTS idx_photos_group_id ON photos(group_id);
CREATE INDEX IF NOT EXISTS idx_photos_filming_location ON photos(filming_location);
CREATE INDEX IF NOT EXISTS idx_photos_regions ON photos(regions);
CREATE INDEX IF NOT EXISTS idx_photos_time_periods ON photos(time_periods);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_photos_time_periods;
DROP INDEX IF EXISTS idx_photos_regions;
DROP INDEX IF EXISTS idx_photos_filming_location;
DROP INDEX IF EXISTS idx_photos_group_id;
DROP TABLE IF EXISTS photos;
-- +goose StatementEnd
