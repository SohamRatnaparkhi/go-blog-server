package users

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/SohamRatnaparkhi/go-blog-server/db"
	"github.com/google/uuid"

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
			Error: err.Error(),
		})
		return
	}

	apiConfig, dbErr := db.DbInstance()
	if dbErr != nil {
		utils.ResponseJson(res, 400, struct {
			Error string `json:"error"`
		}{
			Error: dbErr.Error(),
		})
		return
	}

	user, err := apiConfig.CreateUser(
		req.Context(),
		database.CreateUserParams{
			ID:        uuid.New(),
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
			Error: err.Error(),
		})
		return
		// fmt.Println(err.Error())
	} else {
		utils.ResponseJson(res, 200, user)
	}

}
