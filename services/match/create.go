package match

import "github.com/BennoAlif/game-matchmaking/model"

func (s MatchService) PlayerJoinRoom(username string, roomID uint) (model.Match, error) {
	match := model.Match{
		PlayerUsername: username,
		RoomID:         roomID,
	}

	if err := s.DB.Create(&match).Error; err != nil {
		return match, err
	}

	return match, nil
}
