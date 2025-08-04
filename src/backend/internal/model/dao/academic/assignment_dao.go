package dao

import (
	database "fundz/internal/database"
	entity "fundz/internal/model/entity/academic"
)

// -------
// Create
// -------
func CreateAssignment(assignment entity.Assignment) error {
	if err := database.DB.Create(&assignment).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllAssignments() ([]entity.Assignment, int64, error) {

	var assignments []entity.Assignment
	var count int64

	result := database.DB.Model(&entity.Assignment{}).Count(&count).Find(&assignments)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return assignments, result.RowsAffected, nil
}

// -------
// Read
// -------
func FindAssignmentById(id string) (entity.Assignment, error) {
	var assignment entity.Assignment
	if err := database.DB.Where("assignment_id = ?", id).First(&assignment).Error; err != nil {
		return assignment, err
	}

	return assignment, nil
}

// -------
// Update
// -------
func UpdateAssignmentById(input entity.Assignment, id string) error {

	var assignment entity.Assignment
	if err := database.DB.Model(&assignment).Where("assignment_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete
// -------
func DeleteAssignmentById(id string) error {

	var assignment entity.Assignment
	if err := database.DB.Model(&assignment).Where("assignment_id = ?", id).Delete(assignment).Error; err != nil {
		return err
	}

	return nil
}
