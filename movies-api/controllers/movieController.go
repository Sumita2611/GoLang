package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"movies-api/models"
)

var Movies []models.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, movie := range Movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode("Movie not found")
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, movie := range Movies {
		if movie.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)

			var updatedMovie models.Movie
			json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = params["id"]

			Movies = append(Movies, updatedMovie)
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
	json.NewEncoder(w).Encode("Movie not found")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, movie := range Movies {
		if movie.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			json.NewEncoder(w).Encode("Movie deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Movie not found")
}
