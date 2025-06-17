package dao

import (
	"fmt"
	database "fundz/internal/database"
	"fundz/internal/repo/entity/academic"
)

// -------
// Create
// -------
func CreateSemester(semester academic.Semester) error {
	if err := database.DB.Create(&semester).Error; err != nil {
		return err
	}
	return nil
}

func GetAllSemester() ([]academic.Semester, int64, error) {

	var semester []academic.Semester
	var count int64

	result := database.DB.Model(&academic.Semester{}).Count(&count).Find(&semester)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return semester, result.RowsAffected, nil
}

func GetSemesterById(pk string) (academic.Semester, error) {
	var semester academic.Semester

	if err := database.DB.Where("semester_id = ?", pk).First(&semester).Error; err != nil {
		return semester, err
	}

	return semester, nil
}

// -------
// Update
// -------
func UpdateSemesterById(academicUpdated academic.Semester, id string) error {

	query := database.DB.Model(&academic.Semester{}).Where("semester_id = ?", id).Updates(academicUpdated)
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
func DeleteSemesterById(id string) error {
	var semester academic.Semester

	query := database.DB.Where("semester_id = ?", id).Delete(semester)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
