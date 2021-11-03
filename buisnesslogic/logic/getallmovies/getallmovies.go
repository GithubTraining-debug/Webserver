package getallmovies

import (
	"backend/database"
	"backend/dbinterface"
	"backend/models"
	"database/sql"
)

func Setup() ([]models.Movie, error) {
	DBCon, err := database.StartDB()
	if err != nil {
		return []models.Movie{}, err
	}
	return logic(DBCon)
}

func logic(db *sql.DB) ([]models.Movie, error) {
	return dbinterface.GetAllMovies(db)
}
