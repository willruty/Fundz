package dao

import (
	"fmt"
	database "fundz/internal/database"
	"fundz/internal/repo/entity/academic"
)

// -------
// Create
// -------
func CreateStudent(student academic.Student) error {
	if err := database.DB.Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func GetAllStudent() ([]academic.Student, int64, error) {

	var student []academic.Student
	var count int64

	result := database.DB.Model(&academic.Student{}).Count(&count).Find(&student)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return student, result.RowsAffected, nil
}

func GetStudentById(pk string) (academic.Student, error) {
	var student academic.Student

	if err := database.DB.Where("student_id = ?", pk).First(&student).Error; err != nil {
		return student, err
	}

	return student, nil
}

// -------
// Update
// -------
func UpdateStudentById(academicUpdated academic.Student, id string) error {

	query := database.DB.Model(&academic.Student{}).Where("student_id = ?", id).Updates(academicUpdated)
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
func DeleteStudentById(id string) error {
	var student academic.Student

	query := database.DB.Where("student_id = ?", id).Delete(student)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
