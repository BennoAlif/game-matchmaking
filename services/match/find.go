package match

import "github.com/BennoAlif/game-matchmaking/model"

func (s MatchService) FindAll() ([]model.Match, error) {
	var mathces []model.Match
	result := s.DB.Find(&mathces)
	return mathces, result.Error
}

func (s MatchService) FundUserInRoom(username string, roomID uint) (model.Match, error) {
	var match model.Match
	result := s.DB.Where(&model.Match{
		PlayerUsername: username,
		RoomID:         roomID,
	})

	return match, result.Error
}
