package dao

import (
	"fmt"
	"fundz/internal/database"
	"fundz/internal/repo/entity/fun"

)

// -------
// Create
// -------
func CreateMatch(match fun.Match) error {
	if err := database.DB.Create(&match).Error; err != nil {
		return err
	}
	return nil
}

func GetAllMatch() ([]fun.Match, int64, error) {

	var match []fun.Match
	var count int64

	result := database.DB.Model(&fun.Match{}).Count(&count).Find(&match)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return match, result.RowsAffected, nil
}

func GetMatchById(pk string) (fun.Match, error) {
	var match fun.Match

	if err := database.DB.Where("match_id = ?", pk).First(&match).Error; err != nil {
		return match, err
	}

	return match, nil
}

// -------
// Update
// -------
func UpdateMatchById(funUpdated fun.Match, id string) error {

	query := database.DB.Model(&fun.Match{}).Where("match_id = ?", id).Updates(funUpdated)
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
func DeleteMatchById(id string) error {
	var match fun.Match

	query := database.DB.Where("match_id = ?", id).Delete(match)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
