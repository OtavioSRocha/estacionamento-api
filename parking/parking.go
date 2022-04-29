package parking

import (
	"context"
	"encoding/json"
	"estacionamento-api/cars"
	"estacionamento-api/database"
	"estacionamento-api/utils"
	"log"
	"net/http"
	"time"
)


func OccupySpot(w http.ResponseWriter, r *http.Request) {
	var parking Parking
    _ = json.NewDecoder(r.Body).Decode(&parking)

	if cars.CarIsRegistred(parking.LicensePlate) {

		db := database.ConectDB()
		query := "UPDATE spots SET car = ?, isempty = ? WHERE id = ?"
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			log.Printf("Error %s when preparing SQL statement", err)
		}
		defer stmt.Close()
		res, err := stmt.ExecContext(ctx, parking.LicensePlate, false, parking.SpotId)
		if err != nil {
			log.Printf("Error %s when inserting row into spot table", err)
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
		}
		log.Printf("%d spots created ", rows)

		utils.RequestResponse(w, 1, "Approved")

	} else {
		
		utils.RequestResponse(w, 2, "Erro: Car is not registred")
	}
}

func UnoccupySpot(w http.ResponseWriter, r *http.Request) {
	var parking Parking
    _ = json.NewDecoder(r.Body).Decode(&parking)

	db := database.ConectDB()
	query := "UPDATE spots SET car = ?, isempty = ? WHERE id = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, nil, true, parking.SpotId)
	if err != nil {
		log.Printf("Error %s when inserting row into spot table", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
	}
	log.Printf("%d spots created ", rows)

	utils.RequestResponse(w, 1, "Approved")

}


type Parking struct {
	SpotId			string `json:"spotid,omitempty"`
	LicensePlate	string `json:"licenseplate,omitempty"`
}