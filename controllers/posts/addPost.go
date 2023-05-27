package posts

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/SohamRatnaparkhi/go-blog-server/db"
	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/SohamRatnaparkhi/go-blog-server/utils"
	"github.com/google/uuid"
)

func CreatePostHandler(w http.ResponseWriter, req *http.Request, user database.GetUserByEmailRow) {
	type reqBody struct {
		Title sql.NullString
		Body  sql.NullString
		Tags  []string
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

	post, dbErr2 := apiConfig.CreatePost(req.Context(), database.CreatePostParams{
		ID:       uuid.New(),
		Title:    bodyDecoded.Title,
		Body:     bodyDecoded.Body,
		AuthorID: user.ID,
	})

	if dbErr2 != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr2)
		return
	}

	utils.ResponseJson(w, http.StatusOK, utils.MapPost(post))
}
