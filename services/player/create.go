package player

import "github.com/BennoAlif/game-matchmaking/model"

func (s PlayerService) CreateOne(username string) (model.Player, error) {
	player := model.Player{
		Username:    username,
		PlayerLevel: 1,
	}

	if err := s.DB.Create(&player).Error; err != nil {
		return player, err
	}
	return player, nil
}
