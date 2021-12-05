package api

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//should not be stored like that, I made it only because it was easier to learn new things
//assign values by your needs
const (
	host     = ""
	port     = 0
	user     = ""
	password = ""
	dbname   = ""
)

func SetUpConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
