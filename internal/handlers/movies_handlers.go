package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thenameiswiiwin/reelingit/internal/data"
	"github.com/thenameiswiiwin/reelingit/internal/logger"
	"github.com/thenameiswiiwin/reelingit/internal/models"
)

type MovieHandler struct {
	storage data.MovieStorage
	logger  *logger.Logger
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("Failed to encode JSON response", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (h *MovieHandler) handleStorageError(w http.ResponseWriter, err error, context string) bool {
	if err != nil {
		if err == data.ErrMovieNotFound {
			http.Error(w, context, http.StatusNotFound)
			return true
		}
		h.logger.Error(context, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return true
	}
	return false
}

func (h *MovieHandler) parseID(w http.ResponseWriter, idStr string) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid ID format", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.storage.GetTopMovies()

	if h.handleStorageError(w, err, "Failed to get top movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.logger.Info("Successfully served top movies")
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.storage.GetRandomMovies()
	if h.handleStorageError(w, err, "Failed to get random movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.logger.Info("Successfully served random movies")
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *MovieHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	order := r.URL.Query().Get("order")
	genreStr := r.URL.Query().Get("genre")

	var genre *int
	if genreStr != "" {
		genreInt, ok := h.parseID(w, genreStr)
		if !ok {
			return
		}
		genre = &genreInt
	}

	var movies []models.Movie
	var err error
	if query != "" {
		movies, err = h.storage.SearchMoviesByName(query, order, genre)
	}
	if h.handleStorageError(w, err, "Failed to search movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.logger.Info("Successfully served search results for query: " + query)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/movies/"):]
	id, ok := h.parseID(w, idStr)
	if !ok {
		return
	}

	movie, err := h.storage.GetMovieByID(id)
	if h.handleStorageError(w, err, "Failed to get movie") {
		return
	}
	if h.writeJSONResponse(w, movie) == nil {
		h.logger.Info("Successfully served movie with ID: " + strconv.Itoa(id))
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *MovieHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.storage.GetAllGenres()
	if h.handleStorageError(w, err, "Failed to get genres") {
		return
	}
	if h.writeJSONResponse(w, genres) == nil {
		h.logger.Info("Successfully served all genres")
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *AccountHandler) SaveToCollection(w http.ResponseWriter, r *http.Request) {
	type CollectionRequest struct {
		MovieID    int    `json:"movie_id"`
		Collection string `json:"collection"`
	}

	var req CollectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode collection request", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	email, ok := r.Context().Value("email").(string)
	if !ok {
		http.Error(w, "Unable to retrieve email", http.StatusInternalServerError)
		return
	}

	success, err := h.storage.SaveCollection(models.User{Email: email},
		req.MovieID, req.Collection)
	if h.handleStorageError(w, err, "Failed to save movie to collection") {
		return
	}

	response := AuthResponse{
		Success: success,
		Message: "Movie added to " + req.Collection + " successfully",
	}

	if err := h.writeJSONResponse(w, response); err == nil {
		h.logger.Info("Successfully added movie to collection: " + req.Collection)
	}
}

func (h *AccountHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value("email").(string)
	if !ok {
		http.Error(w, "Unable to retrieve email", http.StatusInternalServerError)
		return
	}
	details, err := h.storage.GetAccountDetails(email)
	if err != nil {
		http.Error(w, "Unable to retrieve collections", http.StatusInternalServerError)
		return
	}
	if err := h.writeJSONResponse(w, details.Favorites); err == nil {
		h.logger.Info("Successfully sent favorites")
	}
}

func (h *AccountHandler) GetWatchlist(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value("email").(string)
	if !ok {
		http.Error(w, "Unable to retrieve email", http.StatusInternalServerError)
		return
	}
	details, err := h.storage.GetAccountDetails(email)
	if err != nil {
		http.Error(w, "Unable to retrieve watchlist", http.StatusInternalServerError)
		return
	}
	if err := h.writeJSONResponse(w, details.Watchlist); err == nil {
		h.logger.Info("Successfully sent watchlist")
	}
}

func NewMovieHandler(storage data.MovieStorage, log *logger.Logger) *MovieHandler {
	return &MovieHandler{
		storage: storage,
		logger:  log,
	}
}
