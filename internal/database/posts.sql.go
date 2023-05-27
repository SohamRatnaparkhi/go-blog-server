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
VALUES ($1, $2, $3, $4)
RETURNING id, title, body, author_id, url, tags, created_at, updated_at, views
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
	)
	return i, err
}