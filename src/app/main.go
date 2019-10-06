package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func routes() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/score/", getScores).Methods("GET")
	router.HandleFunc("/score/{user}", getScoresForUser).Methods("GET")
	router.HandleFunc("/score/", postScore).Methods("POST")

	http.ListenAndServe(":"+PORT, router)
	/*
	  router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	  router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	*/
}

func main() {
	scores = []Score{
		Score{GameScore: 10, User: "Bob", Date: time.Now(), Level: "fort", Version: "1.2.3"},
		Score{GameScore: 32, User: "Matt", Date: time.Now(), Level: "fort", Version: "1.2.3"},
	}
	routes()
}

var scores []Score

func getScores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scores)
}

func getScoresForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	user := strings.ToLower(vars["user"])
	fmt.Println(user)

	json.NewEncoder(w).Encode(findScoresForUser(user))
}

func findScoresForUser(user string) []Score {
	var response []Score

	for _, element := range scores {
		// index is the index where we are
		// element is the element from someSlice for where we are
		if strings.ToLower(element.User) == user {
			response = append(response, element)
			// can also return each form here
			// json.NewEncoder(w).Encode(element)
		}
	}

	return response
}

func postScore(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Post score received")

	var score Score
	_ = json.NewDecoder(r.Body).Decode(&score)

	AddScore(score)

	json.NewEncoder(w).Encode(score)
}

// Add a score to memory, settings it's datetime to now.
func AddScore(score Score) {
	score.Date = time.Now()
	scores = append(scores, score)
}
