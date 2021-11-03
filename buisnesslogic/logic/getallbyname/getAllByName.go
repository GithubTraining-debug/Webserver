package getallbyname

import (
	"database/sql"
	"backend/database"
	"backend/dbinterface"
	"backend/models"
)

func Setup(name string) ([]models.Movie, error) {
	DBCon, err := database.StartDB()
	if err != nil {
		return []models.Movie{}, err
	}
	return logic(name, DBCon)
}

func logic(name string, db *sql.DB) ([]models.Movie, error) {
	return dbinterface.GetAllByName(name, db)
}
