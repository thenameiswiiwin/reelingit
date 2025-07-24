package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thenameiswiiwin/reelingit/data"
	"github.com/thenameiswiiwin/reelingit/logger"
)

type MovieHandler struct {
	Storage data.MovieStorage
	Logger  *logger.Logger
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.Logger.Error("Failed to encode JSON response", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetTopMovies()
	if err != nil {
		h.Logger.Error("Failed to get top movies", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetRandomMovies()
	if err != nil {
		h.Logger.Error("Failed to get random movies", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	h.writeJSONResponse(w, movies)
}
