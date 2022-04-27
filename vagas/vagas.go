package vagas

import (
	"database/sql"
	"encoding/json"
	"estacionamento-api/database"
	"log"
	"net/http"
)

var (
	result 			[]Spot
	id      		string
	vehicle   		string
	isempty   		bool
	car 			Car
)


func GetSpot(w http.ResponseWriter, r *http.Request){
	db := database.ConectDB()

	rows, err := db.Query("SELECT s.id, s.vehicle , s.isempty, c.model, c.licenseplate FROM spots s LEFT OUTER JOIN cars c ON c.id  = s.car")
	if err != nil {
		log.Fatal(err)
	}
	
	for rows.Next() {
		err := rows.Scan(&id, &vehicle, &isempty, &car.Model, &car.LicensePlate)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Spot{ID: id, Vehicle: vehicle, IsEmpty: isempty, Car: &Car{LicensePlate:car.LicensePlate, Model: car.Model} })
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
	LicensePlate	sql.NullString		`json:"licenseplate,omitempty"`
	Model			sql.NullString 		`json:"model,omitempty"`
}