package main

import (
	"fmt"
	"log"
	"net/http"

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

	const addr = ":8080"
	fmt.Printf("starting server at http://localhost%s\n", addr)

	http.Handle("/", http.FileServer(http.Dir("public")))

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
		logInstance.Error("failed to start server", err)
	}
}
