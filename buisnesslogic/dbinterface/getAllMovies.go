package dbinterface

import (
	"backend/models"
	"database/sql"
)

func GetAllMovies(db *sql.DB) ([]models.Movie, error) {
	stmt, err := db.Prepare("SELECT id, name FROM movie;")
	if err != nil {
		return []models.Movie{}, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return []models.Movie{}, err
	}
	defer rows.Close()

	movies := []models.Movie{}

	for rows.Next() {
		m := models.Movie{}
		err := rows.Scan(&m.ID, &m.Name)
		if err != nil {
			return []models.Movie{}, err
		}
		movies = append(movies, m)
	}

	return movies, err

}
