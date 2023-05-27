package main

import (
	controllers "github.com/SohamRatnaparkhi/go-blog-server/controllers/server"
	router "github.com/SohamRatnaparkhi/go-blog-server/routes"

	"github.com/go-chi/chi"
)

func SetCompleteRouters() chi.Router {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/", controllers.HealthCheck)

	userRouter := router.SetUserRouters()
	apiRouter.Mount("/users", userRouter)

	postsRouter := router.SetPostsRouters()
	apiRouter.Mount("/posts", postsRouter)

	return apiRouter
}
