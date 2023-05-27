-- +goose Up

ALTER Table post ADD COLUMN views INT NOT NULL DEFAULT 0;

-- +goose Down

ALTER Table post DROP COLUMN views;