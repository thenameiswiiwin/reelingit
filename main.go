package main

import (
	"log"
	"net/http"

	"github.com/thenameiswiiwin/reelingit/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {
	// Initialize the logger
	logInstance := initializeLogger()

	// Handle requests to the root path by serving static files from the "public" directory
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Start the HTTP server on port 8080
	const address = ":8080"
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Failed to start server on %s: %v", address, err)
		logInstance.Error("Failed to start server", err)
	}
}
