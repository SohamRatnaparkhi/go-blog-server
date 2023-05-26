package utils

import (
	"database/sql"
	"time"

	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/google/uuid"
)

type DbUser struct {
	ID        uuid.UUID      `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Bio       sql.NullString `json:"bio"`
	Isadmin   bool           `json:"is_admin"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func MapUser(dbUser database.User) DbUser {
	return DbUser{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Bio:       dbUser.Bio,
		Isadmin:   dbUser.Isadmin,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}
