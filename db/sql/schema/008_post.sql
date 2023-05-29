-- +goose Up

ALTER Table post drop COLUMN likes;

ALTER Table post add COLUMN likes int not null DEFAULT 0;

-- +goose Down

ALTER Table post drop COLUMN likes;