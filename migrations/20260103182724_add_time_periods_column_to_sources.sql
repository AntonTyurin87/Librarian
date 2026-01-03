-- +goose Up
-- +goose StatementBegin
ALTER TABLE sources ADD COLUMN time_periods TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sources DROP COLUMN time_periods;
-- +goose StatementEnd
