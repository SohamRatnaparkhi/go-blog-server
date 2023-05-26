package routes

import (
	"github.com/SohamRatnaparkhi/go-blog-server/controllers/server"
	"github.com/SohamRatnaparkhi/go-blog-server/controllers/users"
	"github.com/go-chi/chi"
)

func SetUserRouters() chi.Router {
	var userRouter = chi.NewRouter()
	userRouter.Get("/", server.HealthCheck)
	userRouter.Post("/register", users.HandleRegisterUser)
	userRouter.Post("/login", users.HandleLoginUser)
	return userRouter
}
