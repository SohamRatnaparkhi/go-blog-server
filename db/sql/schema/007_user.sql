-- +goose Up

alter table users add column followers int;

alter table users add column following int;

-- +goose Down

alter table users drop column followers;

alter table users drop column following;