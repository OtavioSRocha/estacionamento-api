package vagas

import (
	"encoding/json"
	"estacionamento-api/database"
	"fmt"
	"log"
	"net/http"
)

var (
	result 		[]Spot
	id      	string
	vehicle   	string
	isempty   	bool
	car 		*string
)


func GetSpot(w http.ResponseWriter, r *http.Request){
	db := database.ConectDB()

	rows, err := db.Query("SELECT * FROM spots")
	if err != nil {
		fmt.Println("Erro no select")
		// log.Fatal(err)
	}
	
	for rows.Next() {
		err := rows.Scan(&id, &vehicle, &isempty, &car)
		if err != nil {
			fmt.Println("Erro nas rows")
			log.Fatal(err)
		}
		result = append(result, Spot{ID: id, Vehicle: vehicle, IsEmpty: isempty})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
    json.NewEncoder(w).Encode(result)
}

type Spot struct {
	ID      	string		`json:"id,omitempty"`
	Vehicle 	string		`json:"vehicle,omitempty"`
	IsEmpty 	bool		`json:"isempty"`
	Car			*Car 		`json:"car,omitempty"`
}

type Car struct {
	LicensePlate	string		`json:"licenseplate,omitempty"`
	Model			string 		`json:"model,omitempty"`
}