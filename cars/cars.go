package cars

import (
	"context"
	"encoding/json"
	"estacionamento-api/database"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	id 				string
	model 			string
	licenseplate 	string
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	var result []Cars
	db := database.ConectDB()

	rows, err := db.Query("SELECT * FROM cars")
	if err != nil {
		log.Fatal(err)
	}
	
	for rows.Next() {
		err := rows.Scan(&id, &model, &licenseplate)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Cars{ID: id, Model: model, LicensePlate: licenseplate})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
    json.NewEncoder(w).Encode(result)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	var result []Cars
	params := mux.Vars(r)
	db := database.ConectDB()
	rows, err := db.Query("SELECT * FROM cars WHERE cars.id = " + params["id"])
	if err != nil {
		log.Fatal(err)
	}
	
	for rows.Next() {
		err := rows.Scan(&id, &model, &licenseplate)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Cars{ID: id, Model: model, LicensePlate: licenseplate})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
    json.NewEncoder(w).Encode(result)
}

func SetCar(w http.ResponseWriter, r *http.Request) {
	var car Cars
    _ = json.NewDecoder(r.Body).Decode(&car)

	db := database.ConectDB()
	query := "INSERT INTO cars(licenseplate, model) VALUES (?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
        log.Printf("Error %s when preparing SQL statement", err)
    }
    defer stmt.Close()
    res, err := stmt.ExecContext(ctx, car.LicensePlate, car.Model)
    if err != nil {
        log.Printf("Error %s when inserting row into spot table", err)
    }
    rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when finding rows affected", err)
    }
    log.Printf("%d spots created ", rows)
}

func CarIsRegistred(licenseplate string) bool{
	var result []Cars

	db := database.ConectDB()
	rows, err := db.Query("SELECT * FROM cars WHERE cars.licenseplate = '" + licenseplate + "'")
	if err != nil {
		log.Fatal(err)
	}
	
	for rows.Next() {
		err := rows.Scan(&id, &model, &licenseplate)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Cars{ID: id, Model: model, LicensePlate: licenseplate})
	}
	fmt.Println(result)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println(len(result))

	return len(result) > 0
}


type Cars struct {
	ID      		string		`json:"id,omitempty"`
	Model 			string		`json:"model,omitempty"`
	LicensePlate 	string		`json:"licenseplate,omitempty"`
}
