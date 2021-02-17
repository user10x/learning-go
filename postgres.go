package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	uuid "github.com/satori/go.uuid"
	"log"
)

func main() {
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=visitors user=nipun.chawla password=")

	if err != nil {
		log.Fatal("Unable to connect to %v", err, conn)
	}

	defer conn.Close()

	log.Println("Connected to database")

	//test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Not able to ping the database", err)
	}

	log.Println("")

	err = getAllRows(conn)

	if err != nil {
		log.Println("could not execute query")
	}

}

func getAllRows(conn *sql.DB) error {

	var id uuid.Generator
	var firstName, lastName string
	rows, err := conn.Query("select id, first_name,last_name from users;")

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	// check error one more time just to bee safe
	if err == rows.Err() || err != nil {
		log.Fatal("Error scanning rows", err)

	}
	return nil

}
