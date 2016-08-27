package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Note struct {
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Info      string  `json:"info"`
	Audio     string  `json:"audio"`
}
type Notes []Note

// Global variables
var notes Notes

func main() {
	DEBUG := true

	if DEBUG {
		notes = Notes{Note{Title: "Paris", Latitude: 5.498, Longitude: 6.689, Info: "Gretest city", Audio: "01.mp3"}}
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/notes", notesIndex)

	http.ListenAndServe(":9000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func notesIndex(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		panic(err)
	}
}
