package dbinterface

import (
	"database/sql"
	"os"
)

const movietablename = "movie"

func InsertIntoDb(db *sql.DB) (string, error) {
	insert, err := os.ReadFile("../../data.sql")
	if err != nil {
		return "", err
	}

	_, err = db.Query(string(insert))
	if err != nil {
		return "", err
	}

	return "table # move # updated!", nil
}

func SetupDatabase(db *sql.DB) (string, error) {
	data := `
	CREATE SCHEMA IF NOT EXISTS tmdb;
	SET SCHEMA 'tmdb';

	CREATE TABLE IF NOT EXISTS ` + movietablename + `(
		id SERIAL, 
		name VARCHAR(255)
	);
	`
	_, err := db.Exec(data)
	if err != nil {
		return "", err
	}

	return "Table and schema created", nil
}

func BreakDB(db *sql.DB) string {
	data := `
	DROP table IF EXISTS ` + movietablename + `;
	DROP SCHEMA IF EXISTS tmdb;
	`
	_, err := db.Exec(data)
	if err != nil {
		panic(err)
	}

	return "Table and schemas dropped"

}
