package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//  Structure regarding how data is present for Rolls
type Roll struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Ingredients string `json:"ingredients"`
}

// creating a dummy DB to store info
var rolls []Roll

// get list of all roles present
func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(rolls)
}

// get data about one roll by it's unique id
func getOneRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// adding a new roll's info
func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)
	json.NewEncoder(w).Encode(newRoll)
}

// updating the data regarding a roll
func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			var newRole Roll
			json.NewDecoder(r.Body).Decode(&newRole)
			newRole.ID = params["id"]
			rolls = append(rolls, newRole)
			json.NewEncoder(w).Encode(newRole)
			return
		}
	}

}

// deleting the data for a specific roll
func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			break
		}
	}
}

// calling various endpoints and storing details about 1 roll in dummy db
func main() {
	rolls = append(rolls, Roll{ID: "1", Price: 40, Name: "Egg Roll", Ingredients: "Egg, Chili sauce, spices"})
	router := mux.NewRouter()
	router.HandleFunc("/rolls", getRolls).Methods("GET")
	router.HandleFunc("/rolls/{id}", getOneRoll).Methods("GET")
	router.HandleFunc("/rolls", createRoll).Methods("POST")
	router.HandleFunc("/rolls/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/rolls/{id}", deleteRoll).Methods("DELETE")
	log.Fatalln(http.ListenAndServe(":3001", router))
}
