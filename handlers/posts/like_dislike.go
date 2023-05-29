package posts

import (
	"net/http"

	"github.com/SohamRatnaparkhi/go-blog-server/db"
	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/SohamRatnaparkhi/go-blog-server/utils"
	"github.com/google/uuid"
)

func LikePost(w http.ResponseWriter, req *http.Request, _ database.GetUserByEmailRow) {
	apiConfig, dbErr := db.DbInstance()
	if dbErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr)
		return
	}
	post_id := req.URL.Query().Get("post_id")
	post_uuid, parseErr := uuid.Parse(post_id)
	if parseErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, parseErr)
		return
	}
	post, err := apiConfig.IncreaseLikes(req.Context(), post_uuid)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.ResponseJson(w, http.StatusOK, utils.MapPost(post))
}

func DislikePost(w http.ResponseWriter, req *http.Request, _ database.GetUserByEmailRow) {
	apiConfig, dbErr := db.DbInstance()
	if dbErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr)
		return
	}
	post_id := req.URL.Query().Get("post_id")
	post_uuid, parseErr := uuid.Parse(post_id)
	if parseErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, parseErr)
		return
	}
	post, err := apiConfig.DecreaseLikes(req.Context(), post_uuid)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.ResponseJson(w, http.StatusOK, utils.MapPost(post))
}
