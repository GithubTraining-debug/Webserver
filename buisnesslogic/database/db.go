package database

import (
	"backend/dbinterface"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func StartDB() (*sql.DB, error) {
	return connectToDatabase()
}

func connectionString() string {
	const (
		port     = 8080
		user     = "admin"
		password = "abc123"
		host     = "172.18.0.3"
		dbname   = "postgres"
	)

	return fmt.Sprintf(
		"port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		port,
		host,
		user,
		password,
		dbname,
	)
}

func connectToDatabase() (*sql.DB, error) {
	fmt.Println("getting connection")
	db, err := sql.Open("postgres", connectionString())
	if err != nil {
		return nil, err
	}

	setSchema := "SET SCHEMA 'tmdb';"
	_, err = db.Exec(setSchema)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func Initialize() (string, error) {

	connection, err := StartDB()
	if err != nil {
		return "", err
	}
	ret, err := dbinterface.SetupDatabase(connection)
	if err != nil {
		return "", err
	}
	fmt.Printf("ret: %v\n", ret)
	ret, err = dbinterface.InsertIntoDb(connection)
	if err != nil {
		return "", err
	}
	fmt.Printf("ret: %v\n", ret)

	return "created", nil
}

func CreateTable(CreatingTable string, db *sql.DB) error {
	_, err := db.Exec(CreatingTable)
	if err != nil {
		return err
	}
	return nil
}
