package db

import (
	"fmt"
	"github.com/BennoAlif/game-matchmaking/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DBConfig struct {
	DBHost string
	DBUser string
	DBName string
	DBPort string
	DBPass string
}

func LoadDB(config DBConfig) *gorm.DB {
	var err error
	conStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort,
	)

	db, err := gorm.Open(mysql.Open(conStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("!! Database Connection Loaded !!")
	return db
}

func DBMigrate(db *gorm.DB) {
	var err error
	for _, m := range model.RegisterModels() {
		err = db.Debug().AutoMigrate(m.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("!! Database Migration Succeed !!")
}