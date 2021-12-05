package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mux/router/entities"
	"log"
	"net/http"
)

func GetAllClubs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := SetUpConnection()

		query := `
			SELECT * FROM clubs;
		`

		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		var clubs []entities.Club

		for rows.Next() {
			var club entities.Club
			err = rows.Scan(&club.ClubName, &club.StadiumName, &club.StadiumCapacity)
			if err != nil {
				log.Fatal(err)
			}
			clubs = append(clubs, club)
		}

		rows.Close()

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(clubs)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetClub() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := SetUpConnection()

		searchedClub := mux.Vars(r)["name"]

		query := `
			SELECT * FROM clubs WHERE name = $1;
		`

		row, err := db.Query(query, searchedClub)
		if err != nil {
			log.Fatal(err)
		}

		var club entities.Club

		for row.Next() {
			err = row.Scan(&club.ClubName, &club.StadiumName, &club.StadiumCapacity)
			if err != nil {
				log.Fatal(err)
			}
		}

		row.Close()

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&club)
	}
}

func CreateNewClub() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var club entities.Club

		err := json.NewDecoder(r.Body).Decode(&club)
		if err != nil {
			log.Fatal(err)
		}

		db := SetUpConnection()

		query := `
			INSERT INTO clubs (name, stadium_name, stadium_capacity) VALUES ($1, $2, $3);
		`

		_, err = db.Exec(query, club.ClubName, club.StadiumName, club.StadiumCapacity)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(201)
		_, err = w.Write([]byte("Club has been created"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func DeleteClub() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var removedClub entities.Club

		err := json.NewDecoder(r.Body).Decode(&removedClub)
		if err != nil {
			log.Fatal(err)
		}

		db := SetUpConnection()

		query := `
			DELETE FROM clubs WHERE name = $1;
		`

		_, err = db.Exec(query, removedClub.ClubName)
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write([]byte("Club has been removed."))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func UpdateClub() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updatedClub entities.Club

		err := json.NewDecoder(r.Body).Decode(&updatedClub)
		if err != nil {
			log.Fatal(err)
		}

		clubNameMux := mux.Vars(r)["name"]

		db := SetUpConnection()

		query := `
			UPDATE clubs SET name = $1, stadium_name = $2, stadium_capacity= $3 WHERE name = $4
		`

		_, err = db.Exec(query,
			updatedClub.ClubName,
			updatedClub.StadiumName,
			updatedClub.StadiumCapacity,
			clubNameMux)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(202)
		_, err = w.Write([]byte("Club updated"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
