-- +goose Up

CREATE TABLE
    post (
        id UUID NOT NULL,
        title text,
        body text,
        author_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        url text,
        tags text [],
        PRIMARY KEY(id),
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- +goose Down

DROP TABLE post;