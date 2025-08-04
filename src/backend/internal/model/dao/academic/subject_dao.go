package dao

import (
	"fmt"
	database "fundz/internal/database"
	"fundz/internal/model/entity/academic"
)

// -------
// Create
// -------
func CreateSubject(subject academic.Subject) error {
	if err := database.DB.Create(&subject).Error; err != nil {
		return err
	}
	return nil
}

func GetAllSubject() ([]academic.Subject, int64, error) {

	var subject []academic.Subject
	var count int64

	result := database.DB.Model(&academic.Subject{}).Count(&count).Find(&subject)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return subject, result.RowsAffected, nil
}

func GetSubjectById(pk string) (academic.Subject, error) {
	var subject academic.Subject

	if err := database.DB.Where("subject_id = ?", pk).First(&subject).Error; err != nil {
		return subject, err
	}

	return subject, nil
}

// -------
// Update
// -------
func UpdateSubjectById(academicUpdated academic.Subject, id string) error {

	query := database.DB.Model(&academic.Subject{}).Where("subject_id = ?", id).Updates(academicUpdated)
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
func DeleteSubjectById(id string) error {
	var subject academic.Subject

	query := database.DB.Where("subject_id = ?", id).Delete(subject)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
