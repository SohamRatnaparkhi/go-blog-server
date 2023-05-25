package users

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/SohamRatnaparkhi/go-blog-server/db"

	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/SohamRatnaparkhi/go-blog-server/utils"
)

func HandleCreateUser(res http.ResponseWriter, req *http.Request) {
	type reqBody struct {
		FirstName string         `json:"first_name"`
		LastName  string         `json:"last_name"`
		Email     string         `json:"email"`
		Bio       sql.NullString `json:"bio"`
	}
	decoder := json.NewDecoder(req.Body)

	bodyDecoded := reqBody{}

	if err := decoder.Decode(&bodyDecoded); err != nil {
		utils.ResponseJson(res, 400, struct {
			Error string `json:"error"`
		}{
			Error: "Failed to create user",
		})
		return
	}

	apiConfig := db.DbClient

	user, err := apiConfig.CreateUser(
		req.Context(),
		database.CreateUserParams{
			FirstName: bodyDecoded.FirstName,
			LastName:  bodyDecoded.LastName,
			Email:     bodyDecoded.Email,
			Bio:       bodyDecoded.Bio,
		},
	)

	if err != nil {
		utils.ResponseJson(res, 400, struct {
			Error string `json:"error"`
		}{
			Error: "Failed to create user",
		})
		return
	}

	utils.ResponseJson(res, 200, user)
}
