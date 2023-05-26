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

func HandleRegisterUser(w http.ResponseWriter, req *http.Request) {
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
		utils.ResponseJson(w, 400, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	apiConfig, dbErr := db.DbInstance()
	if dbErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr)
		return
	}

	saltValueString := os.Getenv("BCRYPT_SALT_VALUE")

	saltValue, bcryptErr := strconv.Atoi(saltValueString)

	if bcryptErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, bcryptErr)
		return
	}
	hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(bodyDecoded.Password), saltValue)
	if err2 != nil {
		hashedPassword = []byte(bodyDecoded.Password)
	}

	user, failedToAddToDb := apiConfig.CreateUser(
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

	if failedToAddToDb != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, failedToAddToDb)
		return
	}

	// create jwt token
	token, expiryTime, jwtTokenError := utils.GetJwt(utils.Credentials{
		Email:    bodyDecoded.Email,
		Password: string(hashedPassword),
	})

	if jwtTokenError != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, jwtTokenError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "auth-token",
		Value:   token,
		Expires: expiryTime,
	})

	utils.ResponseJson(w, http.StatusCreated, utils.MapRegisterUser(user))
}
