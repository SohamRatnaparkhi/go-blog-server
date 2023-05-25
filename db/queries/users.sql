-- name: CreateUser :one

INSERT INTO
    users (
        first_name,
        last_name,
        email,
        bio
    )
VALUES ($1, $2, $3, $4)
RETURNING *;