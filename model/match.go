package model

import (
	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	PlayerUsername	string `gorm:"size:191"`
	RoomID			uint
}