-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS for_download_sources (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL UNIQUE,
    type TEXT NOT NULL,
    name_ru TEXT NOT NULL,
    name_eng TEXT,
    author_ru TEXT NOT NULL,
    year INTEGER NOT NULL,
    description TEXT,
    download_url TEXT,
    created_at TEXT,
    is_file_download BOOL NOT NULL,
    is_download BOOL NOT NULL
);

CREATE TABLE IF NOT EXISTS files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    text_source_id INTEGER,
    file_data BLOB NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS for_download_sources;
-- +goose StatementEnd
