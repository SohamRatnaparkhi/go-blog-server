-- +goose Up

ALTER Table post add COLUMN likes int;

-- +goose Down

ALTER Table post drop COLUMN likes;