package spots

import (
	"context"
	"database/sql"
	"encoding/json"
	"estacionamento-api/database"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	id      		string
	vehicle   		string
	isempty   		bool
	car 			Car
)

func GetSpots(w http.ResponseWriter, r *http.Request) {
	var result []Spot
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

func GetSpot(w http.ResponseWriter, r *http.Request) {
	var result []Spot
	params := mux.Vars(r)
	db := database.ConectDB()
	rows, err := db.Query("SELECT s.id, s.vehicle , s.isempty, c.model, c.licenseplate FROM spots s LEFT OUTER JOIN cars c ON c.id  = s.car WHERE s.id = " + params["id"])
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

func SetSpot(w http.ResponseWriter, r *http.Request) {
    var spot Spot
    _ = json.NewDecoder(r.Body).Decode(&spot)

	db := database.ConectDB()
	query := "INSERT INTO spots(vehicle, isempty) VALUES (?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
        log.Printf("Error %s when preparing SQL statement", err)
    }
    defer stmt.Close()
    res, err := stmt.ExecContext(ctx, spot.Vehicle, spot.IsEmpty)
    if err != nil {
        log.Printf("Error %s when inserting row into spot table", err)
    }
    rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when finding rows affected", err)
    }
    log.Printf("%d spots created ", rows)
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