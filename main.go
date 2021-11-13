package main

import (
	"github.com/BennoAlif/game-matchmaking/services/db"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func main() {
	DBConfig := db.DBConfig{
		DBHost: "localhost",
		DBUser: "root",
		DBName: "game_matchmaking",
		DBPort: "3306",
		DbPass: "",
	}

	app := Server{db: db.LoadDB(DBConfig)}
	db.DBMigrate(app.db)
}
