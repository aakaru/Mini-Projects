package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(newMovie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[:idx+1]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			var newMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID = strconv.Itoa(rand.Intn(10000000))
			movies = append(movies, newMovie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

}

func main() {
	
	r := mux.NewRouter()
	
	movies = append(movies, Movie{ID: "1", Isbn: "111", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "222", Title: "Movie 2", Director: &Director{Firstname: "Jane", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "333", Title: "Movie 3", Director: &Director{Firstname: "Michael", Lastname: "Johnson"}})
	movies = append(movies, Movie{ID: "4", Isbn: "444", Title: "Movie 4", Director: &Director{Firstname: "Emily", Lastname: "Davis"}})
	movies = append(movies, Movie{ID: "5", Isbn: "555", Title: "Movie 5", Director: &Director{Firstname: "Robert", Lastname: "Brown"}})


	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieByID).Methods("GET")  
	r.HandleFunc("/movies/", createMovie).Methods("POST")       
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") 

	fmt.Println("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
