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

	if err := database.DB.Model(&entity.Assignment{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("assignment_id asc").Find(&assignments)
	return assignments, result.RowsAffected, result.Error
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
	if err := database.DB.Where("assignment_id = ?", id).First(&assignment).Error; err != nil {
		return err
	}

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
	if _, err := FindAssignmentById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&assignment).Where("assignment_id = ?", id).Delete(assignment).Error; err != nil {
		return err
	}

	return nil
}
