package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mux/router/entities"
	"log"
	"net/http"
)

func GetAllPlayers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbConn := SetUpConnection()

		query := `
			SELECT * FROM players;
			`

		rows, err := dbConn.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		var playersToReturn []entities.Player

		for rows.Next() {
			var player entities.Player
			err = rows.Scan(&player.Name, &player.Age, &player.JerseyNumber, &player.Club)
			if err != nil {
				log.Fatal(err)
			}
			playersToReturn = append(playersToReturn, player)
		}

		defer rows.Close()

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(playersToReturn)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetPlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := SetUpConnection()

		searchedPlayerName := mux.Vars(r)["name"]

		query := `
			SELECT * FROM players WHERE name = $1;
		`

		row, err := db.Query(query, searchedPlayerName)
		if err != nil {
			log.Fatal(err)
		}

		var player entities.Player
		for row.Next() {
			err = row.Scan(&player.Name, &player.Age, &player.JerseyNumber, &player.Club)
			if err != nil {
				log.Fatal(err)
			}
		}

		row.Close()

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(player)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CreateNewPlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newPlayer entities.Player

		err := json.NewDecoder(r.Body).Decode(&newPlayer)
		if err != nil {
			log.Fatal(err)
		}

		db := SetUpConnection()

		query := `
			INSERT INTO players (name, age, jersey_number, club) VALUES ($1, $2, $3, $4)
		`

		_, err = db.Exec(query, newPlayer.Name, newPlayer.Age, newPlayer.JerseyNumber, newPlayer.Club)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(201)
		_, err = w.Write([]byte("Player created"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func DeletePlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := SetUpConnection()

		removedPlayerName := mux.Vars(r)["name"]

		query := `
			DELETE FROM players WHERE name = $1
		`

		_, err := db.Exec(query, removedPlayerName)
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write([]byte("Player has been removed"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func UpdatePlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := SetUpConnection()

		var updatedPlayer entities.Player

		err := json.NewDecoder(r.Body).Decode(&updatedPlayer)
		if err != nil {
			log.Fatal(err)
		}

		updatedName := mux.Vars(r)["name"]

		query := `
			UPDATE players SET name = $1, age = $2, jersey_number = $3, club = $4 
					WHERE name = $5
		`

		_, err = db.Exec(query,
			updatedPlayer.Name,
			updatedPlayer.Age,
			updatedPlayer.JerseyNumber,
			updatedPlayer.Club,
			updatedName)

		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(202)
		_, err = w.Write([]byte("Player updated"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
