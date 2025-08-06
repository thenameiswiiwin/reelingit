package data

import (
	"database/sql"
	"errors"
	"regexp"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/thenameiswiiwin/reelingit/internal/logger"
	"github.com/thenameiswiiwin/reelingit/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AccountRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewAccountRepository(db *sql.DB, log *logger.Logger) (*AccountRepository, error) {
	return &AccountRepository{
		db:     db,
		logger: log,
	}, nil
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (r *AccountRepository) Register(name, email, password string) (bool, error) {
	if name == "" || email == "" || password == "" {
		r.logger.Error("Registration validation failed: missing fields", nil)
		return false, ErrRegistrationValidation
	}
	if len(name) < 4 {
		r.logger.Error("Registration validation failed: name too short", nil)
		return false, errors.New("name must be at least 4 characters long")
	}
	if len(password) < 6 {
		r.logger.Error("Registration validation failed: password too short", nil)
		return false, errors.New("password must be at least 6 characters long")
	}
	if !emailRegex.MatchString(email) {
		r.logger.Error("Registration validation failed: invalid email", nil)
		return false, errors.New("invalid email format")
	}
	var exists bool
	err := r.db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)
	`, email).Scan(&exists)
	if err != nil {
		r.logger.Error("Failed to check if user exists", err)
		return false, err
	}
	if exists {
		r.logger.Error("User already exists with email: "+email, nil)
		return false, ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		r.logger.Error("Failed to hash password", err)
		return false, err
	}

	query := `
		INSERT INTO users (name, email, password_hashed, time_created) VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var userID int
	err = r.db.QueryRow(
		query,
		name,
		email,
		string(hashedPassword),
		time.Now(),
	).Scan(&userID)
	if err != nil {
		r.logger.Error("Failed to insert new user", err)
		return false, err
	}

	return true, nil
}

func (r *AccountRepository) Authenticate(email string, password string) (bool, error) {
	if email == "" || password == "" {
		r.logger.Error("Authentication validation failed: missing fields", nil)
		return false, ErrAuthenticationValidation
	}

	var user models.User
	query := `
		SELECT id, name, email, password_hashed
		FROM users 
		WHERE email = $1 AND time_deleted IS NULL
	`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err == sql.ErrNoRows {
		r.logger.Error("User not found for email: "+email, nil)
		return false, ErrAuthenticationValidation
	}
	if err != nil {
		r.logger.Error("Failed to query user by email", err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		r.logger.Error("Password verification failed for user: "+email, err)
		return false, ErrAuthenticationValidation
	}

	updateQuery := `
		UPDATE users 
		SET last_login = $1
		WHERE id = $2
	`
	_, err = r.db.Exec(updateQuery, time.Now(), user.ID)
	if err != nil {
		r.logger.Error("Failed to update last login time for user: "+email, err)
	}

	return true, nil
}

func (r *AccountRepository) GetAccountDetails(email string) (models.User, error) {
	var user models.User
	query := `
		SELECT id, name, email
		FROM users 
		WHERE email = $1 AND time_deleted IS NULL
	`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)
	if err == sql.ErrNoRows {
		r.logger.Error("User not found for email: "+email, nil)
		return models.User{}, ErrUserNotFound
	}
	if err != nil {
		r.logger.Error("Failed to query user by email", err)
		return models.User{}, err
	}

	favoritesQuery := `
		SELECT m.id, m.tmdb_id, m.title, m.tagline, m.release_year, 
		       m.overview, m.score, m.popularity, m.language, 
		       m.poster_url, m.trailer_url
		FROM movies m
		JOIN user_movies um ON m.id = um.movie_id
		WHERE um.user_id = $1 AND um.relation_type = 'favorite'
	`
	favoriteRows, err := r.db.Query(favoritesQuery, user.ID)
	if err != nil {
		r.logger.Error("Failed to query user favorites", err)
		return user, err
	}
	defer favoriteRows.Close()

	for favoriteRows.Next() {
		var m models.Movie
		if err := favoriteRows.Scan(
			&m.ID, &m.TMDB_ID, &m.Title, &m.Tagline, &m.ReleaseYear,
			&m.Overview, &m.Score, &m.Popularity, &m.Language,
			&m.PosterURL, &m.TrailerURL,
		); err != nil {
			r.logger.Error("Failed to scan favorite movie row", err)
			return user, err
		}
		user.Favorites = append(user.Favorites, m)
	}

	watchlistQuery := `
		SELECT m.id, m.tmdb_id, m.title, m.tagline, m.release_year, 
		       m.overview, m.score, m.popularity, m.language, 
		       m.poster_url, m.trailer_url
		FROM movies m
		JOIN user_movies um ON m.id = um.movie_id
		WHERE um.user_id = $1 AND um.relation_type = 'watchlist'
	`
	watchlistRows, err := r.db.Query(watchlistQuery, user.ID)
	if err != nil {
		r.logger.Error("Failed to query user watchlist", err)
		return user, err
	}
	defer watchlistRows.Close()

	for watchlistRows.Next() {
		var m models.Movie
		if err := watchlistRows.Scan(
			&m.ID, &m.TMDB_ID, &m.Title, &m.Tagline, &m.ReleaseYear,
			&m.Overview, &m.Score, &m.Popularity, &m.Language,
			&m.PosterURL, &m.TrailerURL,
		); err != nil {
			r.logger.Error("Failed to scan watchlist movie row", err)
			return user, err
		}
		user.Watchlist = append(user.Watchlist, m)
	}

	return user, nil
}

func (r *AccountRepository) SaveCollection(user models.User, movieID int, collection string) (bool, error) {
	if movieID <= 0 {
		r.logger.Error("SaveCollection failed: invalid movie ID", nil)
		return false, errors.New("invalid movie ID")
	}
	if collection != "favorite" && collection != "watchlist" {
		r.logger.Error("SaveCollection failed: invalid collection type", nil)
		return false, errors.New("collection must be 'favorite' or 'watchlist'")
	}

	var userID int
	err := r.db.QueryRow(`
		SELECT id 
		FROM users 
		WHERE email = $1 AND time_deleted IS NULL
	`, user.Email).Scan(&userID)
	if err == sql.ErrNoRows {
		r.logger.Error("User not found for email: "+user.Email, nil)
		return false, ErrUserNotFound
	}
	if err != nil {
		r.logger.Error("Failed to query user by email", err)
		return false, err
	}

	var exists bool
	err = r.db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 
			FROM user_movies 
			WHERE user_id = $1 
			AND movie_id = $2 
			AND relation_type = $3
		)
	`, userID, movieID, collection).Scan(&exists)
	if err != nil {
		r.logger.Error("Failed to check if movie exists in collection", err)
		return false, err
	}
	if exists {
		r.logger.Info("Movie " + strconv.Itoa(movieID) + " already exists in " + collection + " for user")
		return true, nil
	}

	query := `
		INSERT INTO user_movies (user_id, movie_id, relation_type, time_added)
		VALUES ($1, $2, $3, $4)
	`
	_, err = r.db.Exec(query, userID, movieID, collection, time.Now())
	if err != nil {
		r.logger.Error("Failed to insert movie into user collection", err)
		return false, err
	}

	r.logger.Info("Movie " + string(movieID) + " added to " + collection + " for user: " + user.Email)
	return true, nil
}

var (
	ErrRegistrationValidation   = errors.New("registration validation failed")
	ErrAuthenticationValidation = errors.New("authentication validation failed")
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrUserNotFound             = errors.New("user not found")
)
