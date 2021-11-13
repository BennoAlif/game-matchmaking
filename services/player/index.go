package player

import "gorm.io/gorm"

type PlayerService struct {
	DB *gorm.DB
}