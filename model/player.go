package model

import (
	"gorm.io/gorm"
	"time"
)

type Player struct {
	Username    string `gorm:"primaryKey"`
	PlayerLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Matches     []Match        `gorm:"foreignKey:PlayerUsername"`
	Room        Room           `gorm:"foreignKey:OwnerUsername"`
}
