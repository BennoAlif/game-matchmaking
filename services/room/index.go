package room

import "gorm.io/gorm"

type RoomService struct {
	DB *gorm.DB
}
