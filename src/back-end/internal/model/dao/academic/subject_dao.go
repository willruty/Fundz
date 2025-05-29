package dao

import (
	"nds_service/internal/database"
	"nds_service/internal/entity"
)

// -------
// Create
// -------
func CreateSubject(subject entity.Subject) error {
	if err := database.DB.Create(&subject).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllSubjects() ([]entity.Subject, int64, error) {

	var subjects []entity.Subject
	var count int64

	if err := database.DB.Model(&entity.Subject{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("subject_id asc").Find(&subjects) 
	return subjects, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindSubjectById(id string) (entity.Subject, error) {
	var subject entity.Subject

	if err := database.DB.Where("subject_id = ?", id).First(&subject).Error; err != nil {
		return subject, err
	}

	return subject, nil
}

// -------
// Update
// -------
func UpdateSubjectById(input entity.Subject, id string) error {

	var subject entity.Subject
	if err := database.DB.Where("subject_id = ?", id).First(&subject).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&subject).Where("subject_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeleteSubjectById(id string) error {

	var subject entity.Subject
	if _, err := FindSubjectById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&subject).Where("subject_id = ?", id).Delete(subject).Error; err != nil {
		return err
	}

	return nil
}
