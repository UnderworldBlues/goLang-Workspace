package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type movie struct {
	movieID       string `json:"id"`
	movieTitle    string `json:"title"`
	isbn          string `json:"isbn"`
	movieDirector string `json:"director"`
}

var moviesSlice []movie // movie slice

func processSampleData(movieName string) {
	// opens file
	file, err := os.Open(movieName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// reads the file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// splits the line
		line := strings.Split(scanner.Text(), ":")
		// creates a movie
		movie := movie{line[0], line[1], line[2], line[3]}
		// appends it to the slice
		moviesSlice = append(moviesSlice, movie)
	}
}
