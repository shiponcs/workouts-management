-- +goose Up
-- +goose StatementBegin
ALTER TABLE workouts
ADD COLUMN version BIGINT NOT NULL DEFAULT 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE workouts DROP COLUMN  version;
-- +goose StatementEnd