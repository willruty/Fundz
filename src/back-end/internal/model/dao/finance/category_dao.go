package dao

import (
	"nds_service/internal/database"
	"nds_service/internal/entity"
)

// -------
// Create
// -------
func CreateCategory(category entity.Category) error {
	if err := database.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllCategorys() ([]entity.Category, int64, error) {

	var categorys []entity.Category
	var count int64

	if err := database.DB.Model(&entity.Category{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("category_id asc").Find(&categorys) 
	return categorys, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindCategoryById(id string) (entity.Category, error) {
	var category entity.Category

	if err := database.DB.Where("category_id = ?", id).First(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

// -------
// Update
// -------
func UpdateCategoryById(input entity.Category, id string) error {

	var category entity.Category
	if err := database.DB.Where("category_id = ?", id).First(&category).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&category).Where("category_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeleteCategoryById(id string) error {

	var category entity.Category
	if _, err := FindCategoryById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&category).Where("category_id = ?", id).Delete(category).Error; err != nil {
		return err
	}

	return nil
}
