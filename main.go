package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"net/url"
	"os"
)

type Note struct {
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
	Address   string  `json:"address"`
	Audio     string  `json:"audio"`
	Tagger    string  `json:"tagger"`
}
type Notes []Note

// Global variables
var notes Notes
var basePath string

func main() {
	basePath = "files/"
	DEBUG := true

	if DEBUG {
		notes = Notes{
			Note{Title: "Spazio Officina", Latitude: 45.837, Longitude: 9.027, City: "Chiasso", Address: "Svizzera, Via D. Alighieri 4, 6830 Chiasso", Audio: "chiasso_01.m4a", Tagger: "Elena"},
			Note{Title: "M.A.X. Museo", Latitude: 47.8, Longitude: 10.0, City: "Chiasso", Address: "Svizzera, Via D. Alighieri 6, 6830 Chiasso", Audio: "chiasso_02.m4a", Tagger: "Andrea"}}
	}

	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/notes", notesIndex)
	r.HandleFunc("/audio/{filename:[a-z]+[_][0-9]+[.](mp4|m4a)}", getAudioFile)

	http.Handle("/", r)

	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Get the Port from the environment so we can run on Heroku
func getPort() string {
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

func getAudioFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	log.Println("Requested: ", filename)
	http.ServeFile(w, r, filename)
}
