package main

import (
	handlers "github.com/SohamRatnaparkhi/go-blog-server/handlers/server"
	router "github.com/SohamRatnaparkhi/go-blog-server/routes"

	"github.com/go-chi/chi"
)

func SetCompleteRouters() chi.Router {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/", handlers.HealthCheck)

	userRouter := router.SetUserRouters()
	apiRouter.Mount("/users", userRouter)

	postsRouter := router.SetPostsRouters()
	apiRouter.Mount("/posts", postsRouter)

	return apiRouter
}
