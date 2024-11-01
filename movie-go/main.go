package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	c "movie-go/constants"
	h "movie-go/handlers"
	"net/http"
)

func main() {
	h.InitMovies()
	r := mux.NewRouter()

	r.Headers("Access-Control-Allow-Origin", "*")
	r.Headers("Content-Type", "application/json")

	r.HandleFunc(c.GetMoviesEndpoint, h.GetMovies).Methods(c.GET)
	r.HandleFunc(c.GetMovieIDEndpoint, h.GetMovieById).Methods(c.GET)
	r.HandleFunc(c.CreateMovieEndpoint, h.CreateMovie).Methods(c.POST)
	r.HandleFunc(c.UpdateMovieEndpoint, h.UpdateMovieById).Methods(c.PUT)
	r.HandleFunc(c.DeleteMovieEndpoint, h.DeleteMovie).Methods(c.DELETE)

	fmt.Printf("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
