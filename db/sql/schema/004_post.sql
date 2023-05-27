-- +goose Up

CREATE TABLE
    post (
        id int NOT NULL,
        title text,
        body text,
        PRIMARY KEY(id),
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- +goose Down

DROP TABLE post;