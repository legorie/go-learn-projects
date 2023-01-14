package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "log"
	// "encoding/json"
	// "math/rand"
	// "net/http"
	// "strconv"
	// "github.com/gorilla.mux"
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

func main() {
	fmt.Println("test")
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "3893", Title: "God Father", Director: &Director{Firstname: "James", Lastname: "Cam"}})
	movies = append(movies, Movie{"2", "3233", "II God Father", &Director{"James", "Cam"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Staring server at port 8000...")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getMovies print")
	w.Header().Set("Content.Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content.Type", "application/json")
	params := mux.Vars(r)
	var mov Movie
	for _, item := range movies {
		if item.ID == params["id"] {
			mov = item
			break
		}
	}
	json.NewEncoder(w).Encode(mov)

}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content.Type", "application/json")
	var mov Movie
	_ = json.NewDecoder(r.Body).Decode(&mov)
	mov.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, mov)
	json.NewEncoder(w).Encode(mov)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content.Type", "applicatin/json")
	params := mux.Vars(r)
	var mov Movie
	for index, item := range movies {
		if item.ID == params["id"] {
			fmt.Print("deleted item", item)
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	mov.ID = params["id"]
	_ = json.NewDecoder(r.Body).Decode(&mov)
	movies = append(movies, mov)
	json.NewEncoder(w).Encode(mov)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In delete function ...")
	w.Header().Set("Content.Type", "application/json")
	params := mux.Vars(r)
	fmt.Print(params)
	for index, item := range movies {
		if item.ID == params["id"] {
			fmt.Print("deleted item", item)
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
