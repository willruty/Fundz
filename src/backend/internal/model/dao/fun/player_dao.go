package dao

import (
	database "fundz/internal/database"
	entity "fundz/internal/model/entity/fun"
)

// -------
// Create
// -------
func CreatePlayer(player entity.Player) error {
	if err := database.DB.Create(&player).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllPlayers() ([]entity.Player, int64, error) {

	var players []entity.Player
	var count int64

	if err := database.DB.Model(&entity.Player{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("player_id asc").Find(&players) 
	return players, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindPlayerById(id string) (entity.Player, error) {
	var player entity.Player

	if err := database.DB.Where("player_id = ?", id).First(&player).Error; err != nil {
		return player, err
	}

	return player, nil
}

// -------
// Update
// -------
func UpdatePlayerById(input entity.Player, id string) error {

	var player entity.Player
	if err := database.DB.Where("player_id = ?", id).First(&player).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&player).Where("player_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeletePlayerById(id string) error {

	var player entity.Player
	if _, err := FindPlayerById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&player).Where("player_id = ?", id).Delete(player).Error; err != nil {
		return err
	}

	return nil
}
