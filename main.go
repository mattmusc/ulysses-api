package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
		notes = Notes{
			Note{Title: "london", Latitude: 5.498, Longitude: 6.689, Info: "gretest city", Audio: "01.mp3"},
			Note{Title: "paris", Latitude: 48.8567, Longitude: 2.3508, Info: "often called the City of Light.", Audio: "01.mp3"},
			Note{Title: "rome", Latitude: 41.9, Longitude: 12.5, Info: "has a whole country inside it.", Audio: "01.mp3"}}
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/notes", notesIndex)

	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "9000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	} else {
		fmt.Println("INFO: Listening on " + port)
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func notesIndex(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		panic(err)
	}
}
