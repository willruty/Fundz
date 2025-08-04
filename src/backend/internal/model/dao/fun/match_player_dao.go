package dao

import (
	"fmt"
	"fundz/internal/database"
	"fundz/internal/model/entity/fun"
)

// -------
// Create
// -------
func CreateMatchPlayer(match_player fun.MatchPlayer) error {
	if err := database.DB.Create(&match_player).Error; err != nil {
		return err
	}
	return nil
}

func GetAllMatchPlayer() ([]fun.MatchPlayer, int64, error) {

	var match_player []fun.MatchPlayer
	var count int64

	result := database.DB.Model(&fun.Match{}).Count(&count).Find(&match_player)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return match_player, result.RowsAffected, nil
}

func GetMatchPlayerById(pk string) (fun.MatchPlayer, error) {
	var match_player fun.MatchPlayer

	if err := database.DB.Where("match_id = ?", pk).First(&match_player).Error; err != nil {
		return match_player, err
	}

	return match_player, nil
}

// -------
// Update
// -------
func UpdateMatchPlayerById(funUpdated fun.MatchPlayer, id string) error {

	query := database.DB.Model(&fun.MatchPlayer{}).Where("match_id = ?", id).Updates(funUpdated)
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
func DeleteMatchPlayerById(id string) error {
	var match_player fun.MatchPlayer

	query := database.DB.Where("match_id = ?", id).Delete(match_player)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
