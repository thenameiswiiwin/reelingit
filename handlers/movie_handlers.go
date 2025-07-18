package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thenameiswiiwin/reelingit/models"
)

type MovieHandler struct {
	// TODO: Add database connection or service layer if needed
}

func (m *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     550,
			Title:       "Fight Club",
			ReleaseYear: 1999,
			Genres:      []models.Genre{{ID: 18, Name: "Drama"}, {ID: 53, Name: "Thriller"}},
			Keywords:    []string{"support group", "dual identity", "anti hero", "psychological thriller"},
			Casting: []models.Actor{
				{
					ID:        1,
					FirstName: "Edward",
					LastName:  "Norton",
				},
			},
		},
		{
			ID:          2,
			TMDB_ID:     680,
			Title:       "Pulp Fiction",
			ReleaseYear: 1994,
			Genres:      []models.Genre{{ID: 80, Name: "Crime"}, {ID: 53, Name: "Thriller"}},
			Keywords:    []string{"nonlinear timeline", "hitman", "crime boss", "cult film"},
			Casting: []models.Actor{
				{
					ID:        2,
					FirstName: "John",
					LastName:  "Travolta",
				},
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode reponse", http.StatusInternalServerError)
	}
}
