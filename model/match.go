package model

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	PlayerUsername string
	RoomID         string
}
