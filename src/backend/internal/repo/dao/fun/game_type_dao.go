package dao

import (
	"fmt"
	"fundz/internal/database"
	"fundz/internal/repo/entity/fun"
)

// -------
// Create
// -------
func CreateGameType(game_type fun.GameType) error {
	if err := database.DB.Create(&game_type).Error; err != nil {
		return err
	}
	return nil
}

func GetAllGameType() ([]fun.GameType, int64, error) {

	var game_type []fun.GameType
	var count int64

	result := database.DB.Model(&fun.GameType{}).Count(&count).Find(&game_type)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return game_type, result.RowsAffected, nil
}

func GetGameTypeById(pk string) (fun.GameType, error) {
	var game_type fun.GameType

	if err := database.DB.Where("game_type_id = ?", pk).First(&game_type).Error; err != nil {
		return game_type, err
	}

	return game_type, nil
}

// -------
// Update
// -------
func UpdateGameTypeById(funUpdated fun.GameType, id string) error {

	query := database.DB.Model(&fun.GameType{}).Where("game_type_id = ?", id).Updates(funUpdated)
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
func DeleteGameTypeById(id string) error {
	var game_type fun.GameType

	query := database.DB.Where("game_type_id = ?", id).Delete(game_type)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
