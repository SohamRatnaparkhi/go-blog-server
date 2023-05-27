-- +goose Up

ALTER TABLE users ADD COLUMN password TEXT NOT NULL DEFAULT '';

-- +goose Down

ALTER TABLE users DROP COLUMN password;