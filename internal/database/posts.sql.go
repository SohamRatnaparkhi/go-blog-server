// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createPost = `-- name: CreatePost :one

INSERT INTO
    post (id, title, body, author_id)
VALUES ($1, $2, $3, $4) RETURNING id, title, body, author_id, url, tags, created_at, updated_at, views, likes
`

type CreatePostParams struct {
	ID       uuid.UUID
	Title    sql.NullString
	Body     sql.NullString
	AuthorID uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.Title,
		arg.Body,
		arg.AuthorID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.AuthorID,
		&i.Url,
		pq.Array(&i.Tags),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Views,
		&i.Likes,
	)
	return i, err
}

const decreaseLikes = `-- name: DecreaseLikes :one

UPDATE post SET likes = likes - 1 WHERE id = $1 RETURNING id, title, body, author_id, url, tags, created_at, updated_at, views, likes
`

func (q *Queries) DecreaseLikes(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, decreaseLikes, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.AuthorID,
		&i.Url,
		pq.Array(&i.Tags),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Views,
		&i.Likes,
	)
	return i, err
}

const increaseLikes = `-- name: IncreaseLikes :one

UPDATE post SET likes = likes + 1 WHERE id = $1 RETURNING id, title, body, author_id, url, tags, created_at, updated_at, views, likes
`

func (q *Queries) IncreaseLikes(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, increaseLikes, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.AuthorID,
		&i.Url,
		pq.Array(&i.Tags),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Views,
		&i.Likes,
	)
	return i, err
}

const updateViewCount = `-- name: UpdateViewCount :one

UPDATE post
SET view_count = view_count + 1
WHERE id = $1 RETURNING id, title, body, author_id, url, tags, created_at, updated_at, views, likes
`

func (q *Queries) UpdateViewCount(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, updateViewCount, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.AuthorID,
		&i.Url,
		pq.Array(&i.Tags),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Views,
		&i.Likes,
	)
	return i, err
}

const viewAllPosts = `-- name: ViewAllPosts :many

SELECT id, title, body, author_id, url, tags, created_at, updated_at, views, likes FROM post ORDER BY created_at DESC
`

func (q *Queries) ViewAllPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, viewAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.AuthorID,
			&i.Url,
			pq.Array(&i.Tags),
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Views,
			&i.Likes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const viewAllPostsByPage = `-- name: ViewAllPostsByPage :many

SELECT id, title, body, author_id, url, tags, created_at, updated_at, views, likes FROM post ORDER BY created_at DESC LIMIT $1 OFFSET $2
`

type ViewAllPostsByPageParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ViewAllPostsByPage(ctx context.Context, arg ViewAllPostsByPageParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, viewAllPostsByPage, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.AuthorID,
			&i.Url,
			pq.Array(&i.Tags),
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Views,
			&i.Likes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const viewPostByID = `-- name: ViewPostByID :one

SELECT id, title, body, author_id, url, tags, created_at, updated_at, views, likes FROM post WHERE id = $1
`

func (q *Queries) ViewPostByID(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, viewPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.AuthorID,
		&i.Url,
		pq.Array(&i.Tags),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Views,
		&i.Likes,
	)
	return i, err
}

const viewPostsByAuthor = `-- name: ViewPostsByAuthor :many

SELECT id, title, body, author_id, url, tags, created_at, updated_at, views, likes
FROM post
WHERE author_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type ViewPostsByAuthorParams struct {
	AuthorID uuid.UUID
	Limit    int32
	Offset   int32
}

func (q *Queries) ViewPostsByAuthor(ctx context.Context, arg ViewPostsByAuthorParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, viewPostsByAuthor, arg.AuthorID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.AuthorID,
			&i.Url,
			pq.Array(&i.Tags),
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Views,
			&i.Likes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
