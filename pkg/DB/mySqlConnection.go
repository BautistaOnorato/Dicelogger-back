package db

import (
	"database/sql"
	"fmt"
	"os"
)

// connectDB connects to the database.
func ConnectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

	// Create the data source.
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open the connection.
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	// Check the connection.
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}