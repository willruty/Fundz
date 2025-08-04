package dao

import (
	"fmt"
	database "fundz/internal/database"
	"fundz/internal/model/entity/academic"
)

// -------
// Create
// -------
func CreateProfessor(professor academic.Professor) error {
	if err := database.DB.Create(&professor).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProfessor() ([]academic.Professor, int64, error) {

	var professor []academic.Professor
	var count int64

	result := database.DB.Model(&academic.Professor{}).Count(&count).Find(&professor)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return professor, result.RowsAffected, nil
}

func GetProfessorById(pk string) (academic.Professor, error) {
	var professor academic.Professor

	if err := database.DB.Where("professor_id = ?", pk).First(&professor).Error; err != nil {
		return professor, err
	}

	return professor, nil
}

// -------
// Update
// -------
func UpdateProfessorById(academicUpdated academic.Professor, id string) error {

	query := database.DB.Model(&academic.Professor{}).Where("professor_id = ?", id).Updates(academicUpdated)
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
func DeleteProfessorById(id string) error {
	var professor academic.Professor

	query := database.DB.Where("professor_id = ?", id).Delete(professor)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
