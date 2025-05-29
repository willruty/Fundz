package dao

import (
	"nds_service/internal/database"
	"nds_service/internal/entity"
)

// -------
// Create
// -------
func CreateStudy_session(study_session entity.Study_session) error {
	if err := database.DB.Create(&study_session).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllStudy_sessions() ([]entity.Study_session, int64, error) {

	var study_sessions []entity.Study_session
	var count int64

	if err := database.DB.Model(&entity.Study_session{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("study_session_id asc").Find(&study_sessions) 
	return study_sessions, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindStudy_sessionById(id string) (entity.Study_session, error) {
	var study_session entity.Study_session

	if err := database.DB.Where("study_session_id = ?", id).First(&study_session).Error; err != nil {
		return study_session, err
	}

	return study_session, nil
}

// -------
// Update
// -------
func UpdateStudy_sessionById(input entity.Study_session, id string) error {

	var study_session entity.Study_session
	if err := database.DB.Where("study_session_id = ?", id).First(&study_session).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&study_session).Where("study_session_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeleteStudy_sessionById(id string) error {

	var study_session entity.Study_session
	if _, err := FindStudy_sessionById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&study_session).Where("study_session_id = ?", id).Delete(study_session).Error; err != nil {
		return err
	}

	return nil
}
