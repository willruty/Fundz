package dao

import (
	"nds_service/internal/database"
	"nds_service/internal/entity"
)

// -------
// Create
// -------
func CreateGame_type(game_type entity.Game_type) error {
	if err := database.DB.Create(&game_type).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllGame_types() ([]entity.Game_type, int64, error) {

	var game_types []entity.Game_type
	var count int64

	if err := database.DB.Model(&entity.Game_type{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("game_type_id asc").Find(&game_types) 
	return game_types, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindGame_typeById(id string) (entity.Game_type, error) {
	var game_type entity.Game_type

	if err := database.DB.Where("game_type_id = ?", id).First(&game_type).Error; err != nil {
		return game_type, err
	}

	return game_type, nil
}

// -------
// Update
// -------
func UpdateGame_typeById(input entity.Game_type, id string) error {

	var game_type entity.Game_type
	if err := database.DB.Where("game_type_id = ?", id).First(&game_type).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&game_type).Where("game_type_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeleteGame_typeById(id string) error {

	var game_type entity.Game_type
	if _, err := FindGame_typeById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&game_type).Where("game_type_id = ?", id).Delete(game_type).Error; err != nil {
		return err
	}

	return nil
}
