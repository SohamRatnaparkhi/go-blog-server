-- name: CreateUser :one

INSERT INTO
    users (
        id,
        first_name,
        last_name,
        password,
        email,
        bio
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one

SELECT
    first_name,
    last_name,
    email,
    bio
FROM users
WHERE email = $1 AND password = $2;

-- name: GetUserByEmail :one

SELECT
    email,
    password,
    first_name,
    bio,
    last_name
FROM users
WHERE email = $1;