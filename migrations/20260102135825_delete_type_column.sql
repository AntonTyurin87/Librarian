-- +goose Up
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_sources_type;
ALTER TABLE sources DROP COLUMN type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sources ADD COLUMN type TEXT;
CREATE INDEX IF NOT EXISTS idx_sources_type ON sources(type);
-- +goose StatementEnd
