package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thenameiswiiwin/reelingit/models"
)

type MovieHandler struct {
	// TODO: Add database connection or service layer if needed
}

func (m *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
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
	m.writeJSONResponse(w, movies)
}

func (m *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          3,
			TMDB_ID:     13,
			Title:       "Forrest Gump",
			ReleaseYear: 1994,
			Genres:      []models.Genre{{ID: 35, Name: "Comedy"}, {ID: 18, Name: "Drama"}},
			Keywords:    []string{"vietnam war", "ping pong", "love of life", "box of chocolates"},
			Casting: []models.Actor{
				{
					ID:        3,
					FirstName: "Tom",
					LastName:  "Hanks",
				},
			},
		},
		{
			ID:          4,
			TMDB_ID:     155,
			Title:       "The Dark Knight",
			ReleaseYear: 2008,
			Genres:      []models.Genre{{ID: 28, Name: "Action"}, {ID: 80, Name: "Crime"}, {ID: 18, Name: "Drama"}},
			Keywords:    []string{"vigilante", "joker", "gotham city", "psychopath"},
			Casting: []models.Actor{
				{
					ID:        4,
					FirstName: "Christian",
					LastName:  "Bale",
				},
			},
		},
	}
	m.writeJSONResponse(w, movies)
}
