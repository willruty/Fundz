package dao

import (
	database "fundz/internal/database"
	entity "fundz/internal/repo/entity/fun"
)

// -------
// Create
// -------
func CreateMatch(Match entity.Match) error {
	if err := database.DB.Create(&Match).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllMatchs() ([]entity.Match, int64, error) {

	var Matchs []entity.Match
	var count int64

	if err := database.DB.Model(&entity.Match{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("Match_id asc").Find(&Matchs) 
	return Matchs, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindMatchById(id string) (entity.Match, error) {
	var Match entity.Match

	if err := database.DB.Where("Match_id = ?", id).First(&Match).Error; err != nil {
		return Match, err
	}

	return Match, nil
}

// -------
// Update
// -------
func UpdateMatchById(input entity.Match, id string) error {

	var Match entity.Match
	if err := database.DB.Where("Match_id = ?", id).First(&Match).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&Match).Where("Match_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeleteMatchById(id string) error {

	var Match entity.Match
	if _, err := FindMatchById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&Match).Where("Match_id = ?", id).Delete(&Match).Error; err != nil {
		return err
	}

	return nil
}
