package utils

import (
	"database/sql"

	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/google/uuid"
)

type PostMap struct {
	ID       uuid.UUID      `json:"id"`
	Title    sql.NullString `json:"title"`
	Body     sql.NullString `json:"body"`
	AuthorID uuid.UUID      `json:"author_id"`
	Url      sql.NullString `json:"url"`
	Tags     []string       `json:"tags"`
	Views    int32          `json:"views"`
}

func MapPost(post database.Post) PostMap {
	return PostMap{
		ID:       post.ID,
		Title:    post.Title,
		Body:     post.Body,
		AuthorID: post.AuthorID,
		Url:      post.Url,
		Tags:     post.Tags,
		Views:    post.Views,
	}
}
