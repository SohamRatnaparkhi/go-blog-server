-- name: CreatePost :one

INSERT INTO
    post (id, title, body, author_id)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: ViewPostsByAuthor :many

SELECT *
FROM post
WHERE author_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: ViewPostByID :one

SELECT * FROM post WHERE id = $1;

-- name: ViewAllPostsByPage :many

SELECT * FROM post ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: ViewAllPosts :many

SELECT * FROM post ORDER BY created_at DESC;

-- name: UpdateViewCount :one

UPDATE post
SET view_count = view_count + 1
WHERE id = $1 RETURNING *;

-- name: IncreaseLikes :one

UPDATE post SET likes = likes + 1 WHERE id = $1 RETURNING *;

-- name: DecreaseLikes :one

UPDATE post SET likes = likes - 1 WHERE id = $1 RETURNING *;