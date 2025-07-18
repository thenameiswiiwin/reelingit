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

	const addr = ":8080"
	fmt.Printf("starting server at http://localhost%s\n", addr)

	// Handler for Static Files (Frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Handlers for API Endpoints
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
		logInstance.Error("failed to start server", err)
	}
}
