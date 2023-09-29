package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"

	"net/http"

	"io/ioutil"
	"log"
)

type sentence struct {
	Worlds	    string  `json:worlds`
	Frecuency   float64 `json:frecuency`
	Limit       float32 `json:limit`
	Ideal_limit float32 `json:ideal`
}

type allsentences []sentence

var sentences = allsentences{
	{
		Worlds:      "Hello",
		Frecuency:   0.5,
		Limit:       0.5,
		Ideal_limit: 0.5,
	},
}

func getSentences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sentences)
}

func createSentence(w http.ResponseWriter, r *http.Request) {
	var newSentence sentence
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid sentence")
	}
	json.Unmarshal(reqBody, &newSentence)
	sentences = append(sentences, newSentence)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSentence)
}
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API sssss ")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/sentences", getSentences).Methods("GET")
	router.HandleFunc("/sentences", createSentence).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))

}
