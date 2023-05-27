-- name: CreatePost :one

INSERT INTO
    post (id, title, body, author_id)
VALUES ($1, $2, $3, $4)
RETURNING *;