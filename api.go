package main

import (
	router "github.com/SohamRatnaparkhi/go-blog-server/routes"
	"github.com/go-chi/chi"
)

func SetCompleteRouters() chi.Router {
	apiRouter := chi.NewRouter()

	userRouter := router.SetUserRouters()
	apiRouter.Mount("/users", userRouter)

	return apiRouter
}
