package dao

import (
	"fmt"
	"fundz/internal/database"
	"fundz/internal/repo/entity/fun"
)

// -------
// Create
// -------
func CreateGameType(game_type fun.Game_type) error {
	if err := database.DB.Create(&game_type).Error; err != nil {
		return err
	}
	return nil
}

func GetAllGameType() ([]fun.Game_type, int64, error) {

	var game_type []fun.Game_type
	var count int64

	result := database.DB.Model(&fun.Game_type{}).Count(&count).Find(&game_type)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return game_type, result.RowsAffected, nil
}

func GetGameTypeById(pk string) (fun.Game_type, error) {
	var game_type fun.Game_type

	if err := database.DB.Where("game_type_id = ?", pk).First(&game_type).Error; err != nil {
		return game_type, err
	}

	return game_type, nil
}

// -------
// Update
// -------
func UpdateGameTypeById(funUpdated fun.Game_type, id string) error {

	query := database.DB.Model(&fun.Game_type{}).Where("game_type_id = ?", id).Updates(funUpdated)
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
	var game_type fun.Game_type

	query := database.DB.Where("game_type_id = ?", id).Delete(game_type)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
