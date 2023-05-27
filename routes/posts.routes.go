package routes

import (
	"github.com/SohamRatnaparkhi/go-blog-server/controllers/posts"
	"github.com/SohamRatnaparkhi/go-blog-server/controllers/server"
	"github.com/SohamRatnaparkhi/go-blog-server/middleware"
	"github.com/go-chi/chi"
)

func SetPostsRouters() chi.Router {
	var userRouter = chi.NewRouter()
	userRouter.Get("/", server.HealthCheck)
	userRouter.Get("/create", middleware.Auth(middleware.AuthHandler(posts.CreatePostHandler)))
	return userRouter
}
