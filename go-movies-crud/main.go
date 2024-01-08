package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")           // pull all entries
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // pull a single entry
	router.HandleFunc("/movies", createMovie).Methods("POST")        // create a new entry
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // update an entry
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // delete an entry
	fmt.Println("Server up and running\n")
	log.Fatal(http.ListenAndServe(":8000", router))
	processSampleData("movieDataSample.csv")
	fmt.Println("hi", moviesSlice[0].movieTitle)
}
