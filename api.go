package main

import (
	controllers "github.com/SohamRatnaparkhi/go-blog-server/controllers/server"
	router "github.com/SohamRatnaparkhi/go-blog-server/routes"

	"github.com/go-chi/chi"
)

func SetCompleteRouters() chi.Router {
	apiRouter := chi.NewRouter()

	userRouter := router.SetUserRouters()
	apiRouter.Get("/", controllers.HealthCheck)
	apiRouter.Mount("/users", userRouter)

	return apiRouter
}
