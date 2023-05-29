package posts

import (
	"net/http"
	"strconv"

	"github.com/SohamRatnaparkhi/go-blog-server/db"
	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/SohamRatnaparkhi/go-blog-server/utils"
)

func GetPostsByAuthor(w http.ResponseWriter, req *http.Request, user database.GetUserByEmailRow) {
	apiConfig, dbErr := db.DbInstance()
	if dbErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr)
		return
	}
	offset_string := req.URL.Query().Get("page_no")
	offset, typeCastError := strconv.Atoi(offset_string)
	if typeCastError != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, typeCastError)
		return
	}
	LIMIT := 5
	authorPosts, dbErr2 := apiConfig.ViewPostsByAuthor(req.Context(), database.ViewPostsByAuthorParams{
		AuthorID: user.ID,
		Limit:    int32(LIMIT),
		Offset:   int32(offset-1) * int32(LIMIT),
	})
	if dbErr2 != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, dbErr2)
		return
	}
	utils.ResponseJson(w, http.StatusOK, utils.MapPosts(authorPosts))
}
