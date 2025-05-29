package dao

import (
	"nds_service/internal/database"
	"nds_service/internal/entity"
)

// -------
// Create
// -------
func CreateGame(game entity.Game) error {
	if err := database.DB.Create(&game).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllGames() ([]entity.Game, int64, error) {

	var games []entity.Game
	var count int64

	if err := database.DB.Model(&entity.Game{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("game_id asc").Find(&games) 
	return games, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindGameById(id string) (entity.Game, error) {
	var game entity.Game

	if err := database.DB.Where("game_id = ?", id).First(&game).Error; err != nil {
		return game, err
	}

	return game, nil
}

// -------
// Update
// -------
func UpdateGameById(input entity.Game, id string) error {

	var game entity.Game
	if err := database.DB.Where("game_id = ?", id).First(&game).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&game).Where("game_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeleteGameById(id string) error {

	var game entity.Game
	if _, err := FindGameById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&game).Where("game_id = ?", id).Delete(game).Error; err != nil {
		return err
	}

	return nil
}
