package routes

import (
	"github.com/SohamRatnaparkhi/go-blog-server/controllers/posts"
	"github.com/SohamRatnaparkhi/go-blog-server/controllers/server"
	"github.com/SohamRatnaparkhi/go-blog-server/middleware"
	"github.com/go-chi/chi"
)

func SetPostsRouters() chi.Router {
	var postRouter = chi.NewRouter()
	postRouter.Get("/", server.HealthCheck)
	postRouter.Post("/create", middleware.Auth(middleware.AuthHandler(posts.CreatePostHandler)))
	postRouter.Get("/getAll", middleware.Auth(middleware.AuthHandler(posts.GetAllPosts)))
	postRouter.Get("/getByAuthor", middleware.Auth(middleware.AuthHandler(posts.GetPostsByAuthor)))
	postRouter.Get("/getById", middleware.Auth(middleware.AuthHandler(posts.GetPostById)))
	return postRouter
}
