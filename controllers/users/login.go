package users

import (
	"encoding/json"
	"net/http"

	"github.com/SohamRatnaparkhi/go-blog-server/db"
	"golang.org/x/crypto/bcrypt"

	"github.com/SohamRatnaparkhi/go-blog-server/utils"
)

func HandleLoginUser(w http.ResponseWriter, req *http.Request) {
	type ReqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(req.Body)
	body := ReqBody{}

	err := decoder.Decode(&body)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
	}

	apiConfig, dbErr := db.DbInstance()
	if dbErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr)
		return
	}

	user, err := apiConfig.GetUserByEmail(req.Context(), body.Email)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	authCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if authCheck != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, authCheck)
		return
	}

	jwtToken, expiryTime, tokenErr := utils.GetJwt(utils.Credentials{
		Email:    user.Email,
		Password: user.Password,
	})

	if tokenErr != nil {
		utils.ErrorResponse(w, http.StatusForbidden, tokenErr)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "auth-token",
		Value:   jwtToken,
		Expires: expiryTime,
	})

	utils.ResponseJson(w, http.StatusOK, utils.MapLoginUser(user))
}
