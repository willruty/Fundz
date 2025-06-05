package dao

import (
	database "fundz/internal/database"
	entity "fundz/internal/model/entity/fun"
)

// -------
// Create
// -------
func CreateScore(score entity.Score) error {
	if err := database.DB.Create(&score).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllScores() ([]entity.Score, int64, error) {

	var scores []entity.Score
	var count int64

	if err := database.DB.Model(&entity.Score{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("score_id asc").Find(&scores) 
	return scores, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindScoreById(id string) (entity.Score, error) {
	var score entity.Score

	if err := database.DB.Where("score_id = ?", id).First(&score).Error; err != nil {
		return score, err
	}

	return score, nil
}

// -------
// Update
// -------
func UpdateScoreById(input entity.Score, id string) error {

	var score entity.Score
	if err := database.DB.Where("score_id = ?", id).First(&score).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&score).Where("score_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeleteScoreById(id string) error {

	var score entity.Score
	if _, err := FindScoreById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&score).Where("score_id = ?", id).Delete(score).Error; err != nil {
		return err
	}

	return nil
}
