package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"movies-api/models"
	"movies-api/routes"
	"movies-api/controllers"
)

func main() {
	router := mux.NewRouter()

	// Dummy data
	modelsData := []models.Movie{
		{ID: "1", Title: "Interstellar", Genre: "Sci-Fi", Rating: 8.6, IsPrime: true},
		{ID: "2", Title: "Jawan", Genre: "Action", Rating: 7.5, IsPrime: true},
	}

	// assign to global slice
	controllers.Movies = modelsData

	routes.RegisterMovieRoutes(router)

	log.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
