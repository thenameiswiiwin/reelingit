package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/thenameiswiiwin/reelingit/internal/data"
	"github.com/thenameiswiiwin/reelingit/internal/handlers"
	"github.com/thenameiswiiwin/reelingit/internal/logger"
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

	if err := godotenv.Load(filepath.Join(getProjectRoot(), ".env")); err != nil {
		log.Printf("No .env file found, proceeding without it: %v", err)
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to create movie repository: %v", err)
	}

	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	http.HandleFunc("/api/account/register", movieHandler.GetGenres)
	http.HandleFunc("/api/account/authenticate", movieHandler.GetGenres)

	http.Handle("/", http.FileServer(http.Dir("web")))

	const addr = ":8080"
	fmt.Printf("Starting server on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Failed to start server", err)
		log.Fatalf("Failed to start server on %s: %v", addr, err)
	}
}

func getProjectRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "../../")
}
