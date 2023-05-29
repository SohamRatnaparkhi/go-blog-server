package routes

import (
	"github.com/SohamRatnaparkhi/go-blog-server/handlers/server"
	"github.com/SohamRatnaparkhi/go-blog-server/handlers/users"
	"github.com/go-chi/chi"
)

func SetUserRouters() chi.Router {
	var userRouter = chi.NewRouter()
	userRouter.Get("/", server.HealthCheck)
	userRouter.Post("/register", users.HandleRegisterUser)
	userRouter.Post("/login", users.HandleLoginUser)
	return userRouter
}
