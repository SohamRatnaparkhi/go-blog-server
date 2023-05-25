package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "github.com/SohamRatnaparkhi/go-blog-server/controllers/server"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		panic("No port found")
	}
	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Mount("/v1", v1Router)

	v1Router.Get("/healthz", controllers.HealthCheck)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("\nServer starting at http://localhost:%v", port)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server started at port %v", port)
}
