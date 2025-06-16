# reelingit

A Fullstack SPA App with Go and Vanilla JS

## Technologies

* **Backend**: Go
* **Database**: Postgres
* **Frontend**: HTML, CSS, JS
* **Communication**: JSON RESTful APIs

## Libraries

* [Air](github.com/air-verse/air) Live reload for Go apps
* [GoDotEnv](github.com/joho/godotenv) A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).
* [PG](github.com/lib/pq) Pure Go Postgres driver for database/sql

```sh
go get github.com/joho/godotenv
go get github.com/lib/pq
go install github.com/cosmtrek/air@latest
```

## Data

* The Movie Database (TMDB)
* ~5,000 subset movies with meta data
* Images come from TMDB online server
* Video trailers from YouTube

## Features available

* See top and recent movies
* Search movies
* Movie details
* User registration
* User authentication
* Favorites movies
* Watchlist

## Architecture

![Architecture](./public/images/architecture/architecture_1.png)
![Diagram](./public/images/architecture/architecture_2.png)
![GET/](./public/images/architecture/architecture_3.png)
![Static](./public/images/architecture/architecture_4.png)
![index.html](./public/images/architecture/architecture_5.png)
![/movies/id](./public/images/architecture/architecture_6.png)
![MovieDetailsPage](./public/images/architecture/architecture_7.png)
![Get /api/movies/id](./public/images/architecture/architecture_8.png)
![MovieHandler](./public/images/architecture/architecture_9.png)
![GetMovie](./public/images/architecture/architecture_10.png)
![SELECT*](./public/images/architecture/architecture_11.png)
![Movie](./public/images/architecture/architecture_12.png)
![JSON](./public/images/architecture/architecture_13.png)
![HTML](./public/images/architecture/architecture_14.png)

## Database Schema

### Models 

* Movie
* Genre
* Actor
* User

![Entity](./public/images/models/models_1.png)
![Table](./public/images/models/models_2.png)
