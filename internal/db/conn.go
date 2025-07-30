package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 1234
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {

	// Define the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%d dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Verify the connection to the database
	// This will attempt to ping the database to ensure the connection is valid
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// If the connection is successful, print a message
	fmt.Println("Connected to " + dbname)

	// Return the database connection
	return db, nil
}
