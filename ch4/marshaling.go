package main

import (
	"encoding/json"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"` // This is called a field tag.
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve Mcqueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)

	if err != nil {
		log.Fatalf("JSON Marshaling failed! %s", err)
	}

	log.Printf("%s", data)
}
