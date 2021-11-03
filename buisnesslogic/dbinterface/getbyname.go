package dbinterface

import (
	"backend/models"
	"database/sql"
)

func GetAllByName(name string, db *sql.DB) ([]models.Movie, error) {

	stmt, err := db.Prepare("SELECT * FROM movie WHERE name = $1")
	if err != nil {
		return []models.Movie{}, err
	}
	rows, err := stmt.Query(name)
	if err != nil {
		return []models.Movie{}, err
	}
	movies := []models.Movie{}

	for rows.Next() {
		m := models.Movie{}
		err := rows.Scan(&m.ID, &m.Name)
		if err != nil {
			return []models.Movie{}, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}
