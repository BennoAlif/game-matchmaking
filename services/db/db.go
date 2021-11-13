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
	DbPass string
}

func LoadDB(config DBConfig) *gorm.DB {
	var err error
	conStr := fmt.Sprintf(
		"%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser, config.DBName,
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
