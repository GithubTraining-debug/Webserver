package database

import (
	"backend/dbinterface"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func StartDB() (*sql.DB, error) {
	return connectToDatabase()
}

func connectionString(s DBSettings) string {
	fmt.Printf(
		"port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Port,
		s.Host,
		s.User,
		s.Password,
		s.DBname,
	)
	return fmt.Sprintf(
		"port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Port,
		s.Host,
		s.User,
		s.Password,
		s.DBname,
	)
}

func connectToDatabase() (*sql.DB, error) {
	fmt.Println("getting connection to database")

	settings := DBSettings{}
	settingsBytes, err := os.ReadFile("Settings.json")
	if err != nil {
		return &sql.DB{}, fmt.Errorf("failed to read settingsfile Error: %s", err.Error())
	}
	err = json.Unmarshal(settingsBytes, &settings)
	if err != nil {
		return &sql.DB{}, fmt.Errorf("failed to parse settings Error: %s", err.Error())
	}
	db, err := sql.Open("postgres", connectionString(settings))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db Error: %s", err.Error())
	}
	setSchema := "SET SCHEMA 'tmdb';"
	_, err = db.Exec(setSchema)
	if err != nil {
		return nil, fmt.Errorf("failed to Set Schema: %s", err.Error())
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
