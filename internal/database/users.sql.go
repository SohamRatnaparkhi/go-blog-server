// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: users.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one

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
RETURNING id, first_name, last_name, email, bio, isadmin, created_at, updated_at, password, followers, following
`

type CreateUserParams struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Password  string
	Email     string
	Bio       sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Password,
		arg.Email,
		arg.Bio,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Bio,
		&i.Isadmin,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Password,
		&i.Followers,
		&i.Following,
	)
	return i, err
}

const getUser = `-- name: GetUser :one

SELECT
    id,
    first_name,
    last_name,
    email,
    bio
FROM users
WHERE email = $1 AND password = $2
`

type GetUserParams struct {
	Email    string
	Password string
}

type GetUserRow struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Bio       sql.NullString
}

func (q *Queries) GetUser(ctx context.Context, arg GetUserParams) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, arg.Email, arg.Password)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Bio,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one

SELECT
    id,
    email,
    password,
    first_name,
    bio,
    last_name
FROM users
WHERE email = $1
`

type GetUserByEmailRow struct {
	ID        uuid.UUID
	Email     string
	Password  string
	FirstName string
	Bio       sql.NullString
	LastName  string
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.Bio,
		&i.LastName,
	)
	return i, err
}
