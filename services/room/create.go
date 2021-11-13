package room

import (
	"database/sql"
	"github.com/BennoAlif/game-matchmaking/model"
)

func (s RoomService) CraeteOne(ownerUsername string, title, string, description string) (model.Room, error) {
	room := model.Room{
		Title: title,
		Description: sql.NullString{
			String: description,
		},
		Game:          "Warframe",
		MaxPlayer:     4,
		IsActive:      true,
		GameMode:      "Normal",
		MinLevel:      1,
		OwnerUsername: ownerUsername,
	}

	if err := s.DB.Create(&room).Error; err != nil {
		return room, err
	}

	return room, nil
}
