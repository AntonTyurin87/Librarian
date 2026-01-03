-- +goose Up
-- +goose StatementBegin
ALTER TABLE sources ADD COLUMN availability BOOL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sources DROP COLUMN availability;
-- +goose StatementEnd
