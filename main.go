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
		log.Fatal("No port found")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//Routers

	// database connection
	// db_url := os.Getenv("DB_URL")
	// if db_url == "" {
	// 	panic("No database connection found")
	// }

	// db, dbErr := sql.Open("postgres", db_url)
	// if dbErr != nil {
	// 	panic("Failed to connect to database")
	// }

	// dbQueries := database.New(db)

	// fmt.Print(dbQueries)

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Get("/health", controllers.HealthCheck)

	apiRouter := SetCompleteRouters()
	v1Router.Mount("/api", apiRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("\nServer starting at http://localhost:%v\n", port)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server started at port %v", port)
}
