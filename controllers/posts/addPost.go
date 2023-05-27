package posts

import (
	"net/http"

	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/SohamRatnaparkhi/go-blog-server/utils"
)

func CreatePostHandler(w http.ResponseWriter, _ *http.Request, user database.GetUserByEmailRow) {
	utils.ResponseJson(w, http.StatusOK, utils.MapLoginUser(user))
}
