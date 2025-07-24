package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thenameiswiiwin/reelingit/handlers"
	"github.com/thenameiswiiwin/reelingit/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie-output.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	return logInstance
}

func main() {
	logInstance := initializeLogger()
	defer logInstance.Close()

	movieHandler := handlers.MovieHandler{}

	// Handle API routes
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	// Handle static files
	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	fmt.Printf("Starting server on %s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logInstance.Error("Failed to start server", err)
		log.Fatalf("Failed to start server on %s: %v", addr, err)
	}
}
