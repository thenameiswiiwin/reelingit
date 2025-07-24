package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thenameiswiiwin/reelingit/models"
)

type MovieHandler struct {
	// TODO: logging can be added later
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     101,
			Title:       "Inception",
			ReleaseYear: 2010,
			Genres: []models.Genre{
				{
					ID:   1,
					Name: "Science Fiction",
				},
			},
			Keywords: []string{"dream", "subconscious", "heist"},
			Casting: []models.Actor{
				{
					ID:        1,
					FirstName: "Leonardo",
					LastName:  "DiCaprio",
				},
			},
		},
		{
			ID:          2,
			TMDB_ID:     102,
			Title:       "The Dark Knight",
			ReleaseYear: 2008,
			Genres: []models.Genre{
				{
					ID:   2,
					Name: "Action",
				},
			},
			Keywords: []string{"batman", "joker", "crime"},
			Casting: []models.Actor{
				{
					ID:        2,
					FirstName: "Christian",
					LastName:  "Bale",
				},
			},
		},
		{
			ID:          3,
			TMDB_ID:     103,
			Title:       "Interstellar",
			ReleaseYear: 2014,
			Genres: []models.Genre{
				{
					ID:   1,
					Name: "Science Fiction",
				},
			},
			Keywords: []string{"space", "time travel", "exploration"},
			Casting: []models.Actor{
				{
					ID:        3,
					FirstName: "Matthew",
					LastName:  "McConaughey",
				},
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		// TODO: Log the error using a logger
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
