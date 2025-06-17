package dao

import (
	"fmt"
	"fundz/internal/database"
	"fundz/internal/repo/entity/fun"
)

// -------
// Create
// -------
func CreatePlayer(player fun.Player) error {
	if err := database.DB.Create(&player).Error; err != nil {
		return err
	}
	return nil
}

func GetAllPlayers() ([]fun.Player, int64, error) {

	var player []fun.Player
	var count int64

	result := database.DB.Model(&fun.Player{}).Count(&count).Find(&player)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return player, result.RowsAffected, nil
}

func GetPlayerById(pk string) (fun.Player, error) {
	var player fun.Player

	if err := database.DB.Where("player_id = ?", pk).First(&player).Error; err != nil {
		return player, err
	}

	return player, nil
}

// -------
// Update
// -------
func UpdatePlayerById(funUpdated fun.Player, id string) error {

	query := database.DB.Model(&fun.Player{}).Where("player_id = ?", id).Updates(funUpdated)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}

// -------
// Delete
// -------
func DeletePlayerById(id string) error {
	var player fun.Player

	query := database.DB.Where("player_id = ?", id).Delete(player)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
