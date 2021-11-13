package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Title       string
	Description sql.NullString
	Game        string
	MaxPlayer   int
	IsActive    bool
	GameMode    string
	MinLevel    int
	Matches     []Match `gorm:"foreignKey:RoomID"`
	OwnerUsername       string `gorm:"size:191"`
}
