package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"dataaccess.com/constant"
)

var db *sql.DB

func main() {
	//Capture connection properties
	conStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", constant.USER, constant.PWD, constant.SERVER, constant.PORT, constant.DATABASE)
	conDb, conErr := sql.Open("postgres", conStr)

	if conErr != nil {
		log.Fatal("Error connecting to the database", conErr)
	}

	defer conDb.Close()

	pingErr := conDb.Ping()

	if pingErr != nil {
		log.Fatal("Error ")
	}

	fmt.Println("Connected in database!")
}
