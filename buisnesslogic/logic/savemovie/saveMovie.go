package savemovie

import (
	"backend/database"
	"backend/dbinterface"
	"backend/models"
	"database/sql"
)

func SetupSaveIt(movie models.Movie) error {
	DBCon, err := database.StartDB()
	if err != nil {
		return err
	}
	return save(movie, DBCon)
}

func save(movie models.Movie, db *sql.DB) error {
	return dbinterface.InsertMovie(movie, db)
}
