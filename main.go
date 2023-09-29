package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gorilla/mux"

	"net/http"

	"io/ioutil"
	"log"
)

type sentence struct {
	Worlds	    string  `json:worlds`
	Limit       float64 `json:limit`
	Ideal_limit float64 `json:ideal`
}

type allsentences []sentence

var sentences = allsentences{
	{
		Worlds:      "Hello",
		Limit:       0.5,
		Ideal_limit: 0.5,
	},
}

func getSentences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sentences)
}
func contarLetras(cadena string) map[rune]int {
	resultado := make(map[rune]int)

	// Convierte la cadena a minúsculas para que no distinga entre mayúsculas y minúsculas
	cadena = strings.ToLower(cadena)

	// Recorre la cadena y cuenta las letras
	for _, char := range cadena {
		if 'a' <= char && char <= 'z' {
			resultado[char]++
		}
	}

	return resultado
}
func dividirConteo(resultados map[rune]int, divisor float64) map[rune]float64 {
	resultadoDividido := make(map[rune]float64)

	// Itera sobre los resultados y divide el conteo por el divisor
	for letra, count := range resultados {
		resultadoDividido[letra] = float64(count) / float64(divisor)
	}

	return resultadoDividido
}
func sumarValores(resultadosDivididos map[rune]float64) float64 {
	suma := 0.0

	// Itera sobre los valores y suma
	for _, valor := range resultadosDivididos {
		suma += valor
	}

	return suma
}

func createSentence(w http.ResponseWriter, r *http.Request) {
	var newSentence sentence
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid sentence")
	}
	json.Unmarshal(reqBody, &newSentence)

	worldsWhtoutSpace := strings.ReplaceAll(newSentence.Worlds, " ", "")
	fmt.Println(worldsWhtoutSpace)

	newSentence.Limit= float64(len(worldsWhtoutSpace))
	fmt.Println(newSentence.Limit)

	resultado := contarLetras(worldsWhtoutSpace)

	for letra, count := range resultado {
		fmt.Printf("'%c': %d\n", letra, count)
	}

	resultadoDividido := dividirConteo(resultado, newSentence.Limit )
	for letra, count := range resultadoDividido {
		fmt.Printf("'%c': %.16f\n", letra, count)
	}

	suma := sumarValores(resultadoDividido)

	fmt.Printf("Suma de valores divididos: %.16f\n", suma)

	newSentence.Ideal_limit= suma

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
