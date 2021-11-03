package dbinterface

import (
	"backend/models"
	"database/sql"
)

func InsertMovie(movie models.Movie, db *sql.DB) error {

	stmt, err := db.Prepare("INSERT INTO movie (name) VALUES ($1)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(movie.Name)
	if err != nil {
		return err
	}

	return nil
}
