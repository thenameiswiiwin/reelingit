package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thenameiswiiwin/reelingit/handlers"
	"github.com/thenameiswiiwin/reelingit/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movies-output.log")
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {
	logInstance := initializeLogger()
	movieHandler := handlers.MovieHandler{}

	// Handlers for API Endpoints
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	// Handler for Static Files (Frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Server is running on http://localhost:8080")

	const addr = ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
		logInstance.Error("failed to start server", err)
	}
}
