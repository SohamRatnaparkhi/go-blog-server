-- +goose Up

CREATE TABLE
    users (
        id int NOT NULL,
        first_name VARCHAR(255) NOT NULL,
        last_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        bio text,
        isAdmin BOOLEAN NOT NULL DEFAULT FALSE,
        PRIMARY KEY(id),
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- +goose Down

DROP TABLE users;