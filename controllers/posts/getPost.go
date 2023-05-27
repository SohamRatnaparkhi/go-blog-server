package posts

import (
	"net/http"

	"github.com/SohamRatnaparkhi/go-blog-server/db"
	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/SohamRatnaparkhi/go-blog-server/utils"
	"github.com/google/uuid"
)

func GetPostById(w http.ResponseWriter, req *http.Request, _ database.GetUserByEmailRow) {
	apiConfig, dbErr := db.DbInstance()
	if dbErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr)
		return
	}
	post_id := req.URL.Query().Get("post_id")
	post_uuid, parseErr := uuid.Parse(post_id)
	if parseErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, parseErr)
	}
	post, dbErr2 := apiConfig.ViewPostByID(req.Context(), post_uuid)
	if dbErr2 != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr2)
		return
	}
	utils.ResponseJson(w, http.StatusOK, utils.MapPost(post))
}
