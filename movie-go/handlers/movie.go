package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand/v2"
	m "movie-go/models"
	"net/http"
	"strconv"
)

var movies []m.Movie

func InitMovies() {
	movies = append(movies,
		m.Movie{
			ID:       "1",
			Isbn:     "111",
			Title:    "Movie 1",
			Director: &m.Director{FirstName: "John", LastName: "Doe"}})

	movies = append(movies,
		m.Movie{
			ID:       "2",
			Isbn:     "222",
			Title:    "Movie 2",
			Director: &m.Director{FirstName: "Jane", LastName: "Dave"}})
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID != params["id"] {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie m.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	movie.ID = strconv.Itoa(rand.IntN(1_000_000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var movie m.Movie
	for idx, _movie := range movies {
		if _movie.ID != params["id"] {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if _movie.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for idx, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}
