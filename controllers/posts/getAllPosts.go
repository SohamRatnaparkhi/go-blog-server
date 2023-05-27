package posts

import (
	"net/http"

	"github.com/SohamRatnaparkhi/go-blog-server/db"
	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/SohamRatnaparkhi/go-blog-server/utils"
)

func GetAllPosts(w http.ResponseWriter, req *http.Request, _ database.GetUserByEmailRow) {
	apiConfig, dbErr := db.DbInstance()
	if dbErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr)
		return
	}

	authorPosts, dbErr2 := apiConfig.ViewAllPosts(req.Context())
	if dbErr2 != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr2)
		return
	}
	utils.ResponseJson(w, http.StatusOK, utils.MapPosts(authorPosts))
}
