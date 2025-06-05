package dao

import (
	database "fundz/internal/database"
	entity "fundz/internal/model/entity/finance"
)

// -------
// Create
// -------
func CreateGoal(goal entity.Goal) error {
	if err := database.DB.Create(&goal).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllGoals() ([]entity.Goal, int64, error) {

	var goals []entity.Goal
	var count int64

	if err := database.DB.Model(&entity.Goal{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("goal_id asc").Find(&goals) 
	return goals, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindGoalById(id string) (entity.Goal, error) {
	var goal entity.Goal

	if err := database.DB.Where("goal_id = ?", id).First(&goal).Error; err != nil {
		return goal, err
	}

	return goal, nil
}

// -------
// Update
// -------
func UpdateGoalById(input entity.Goal, id string) error {

	var goal entity.Goal
	if err := database.DB.Where("goal_id = ?", id).First(&goal).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&goal).Where("goal_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeleteGoalById(id string) error {

	var goal entity.Goal
	if _, err := FindGoalById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&goal).Where("goal_id = ?", id).Delete(goal).Error; err != nil {
		return err
	}

	return nil
}
