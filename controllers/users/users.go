package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/SohamRatnaparkhi/go-blog-server/db"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/SohamRatnaparkhi/go-blog-server/utils"
)

func HandleRegisterUser(res http.ResponseWriter, req *http.Request) {
	type reqBody struct {
		FirstName string         `json:"first_name"`
		LastName  string         `json:"last_name"`
		Email     string         `json:"email"`
		Password  string         `json:"password"`
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

	saltValueString := os.Getenv("BCRYPT_SALT_VALUE")

	saltValue, err := strconv.Atoi(saltValueString)

	if err != nil {
		utils.ResponseJson(res, 400, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(bodyDecoded.Password), saltValue)
	if err2 != nil {
		hashedPassword = []byte(bodyDecoded.Password)
	}

	user, err := apiConfig.CreateUser(
		req.Context(),
		database.CreateUserParams{
			ID:        uuid.New(),
			FirstName: bodyDecoded.FirstName,
			LastName:  bodyDecoded.LastName,
			Email:     bodyDecoded.Email,
			Password:  string(hashedPassword),
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
		utils.ResponseJson(res, 200, utils.MapUser(user))
	}

}
