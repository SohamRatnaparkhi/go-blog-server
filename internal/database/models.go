// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID
	Title     sql.NullString
	Body      sql.NullString
	AuthorID  uuid.UUID
	Url       sql.NullString
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
	Views     int32
}

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Bio       sql.NullString
	Isadmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Password  string
}
