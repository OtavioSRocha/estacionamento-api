package main

import (
	"encoding/json"
	"log"
	"net/http"

	"estacionamento-api/cars"
	"estacionamento-api/parking"
	"estacionamento-api/spots"

	"github.com/gorilla/mux"
)

func main() {   
    router := mux.NewRouter()
    
    router.HandleFunc("/spots", spots.GetSpots).Methods("GET")
    router.HandleFunc("/spots/{id}", spots.GetSpot).Methods("GET")
    router.HandleFunc("/spots", spots.SetSpot).Methods("POST")

    router.HandleFunc("/cars", cars.GetCars).Methods("GET")
    router.HandleFunc("/cars/{id}", cars.GetCar).Methods("GET")
    router.HandleFunc("/cars", cars.SetCar).Methods("POST")

    router.HandleFunc("/occupySpot", parking.OccupySpot).Methods("POST")
    
    router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    json.NewEncoder(w).Encode(people)
    }
}

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}

type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person