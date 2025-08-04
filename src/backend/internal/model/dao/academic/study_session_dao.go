package dao

import (
	database "fundz/internal/database"
	"fundz/internal/model/entity/academic"
)

// -------
// Create
// -------
func CreateStudy_session(study_session academic.Study_session) error {
	if err := database.DB.Create(&study_session).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllStudy_sessions() ([]academic.Study_session, int64, error) {

	var study_sessions []academic.Study_session
	var count int64

	result := database.DB.Model(&academic.Study_session{}).Count(&count).Find(&study_sessions)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return study_sessions, result.RowsAffected, nil
}

// -------
// Read
// -------
func FindStudy_sessionById(id string) (academic.Study_session, error) {
	var study_session academic.Study_session
	if err := database.DB.Where("study_session_id = ?", id).First(&study_session).Error; err != nil {
		return study_session, err
	}

	return study_session, nil
}

// -------
// Update
// -------
func UpdateStudy_sessionById(input academic.Study_session, id string) error {
	var study_session academic.Study_session
	if err := database.DB.Model(&study_session).Where("study_session_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete
// -------
func DeleteStudy_sessionById(id string) error {
	var study_session academic.Study_session
	if err := database.DB.Model(&study_session).Where("study_session_id = ?", id).Delete(study_session).Error; err != nil {
		return err
	}

	return nil
}
