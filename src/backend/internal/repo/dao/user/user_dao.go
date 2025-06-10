package dao

import (
	database "fundz/internal/database"
	entity "fundz/internal/repo/entity/user"
)

// -------
// Create
// -------
func CreateUser(user entity.UserAccount) error {
	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllUsers() ([]entity.UserAccount, int64, error) {

	var users []entity.UserAccount
	var count int64

	if err := database.DB.Model(&entity.UserAccount{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("user_id asc").Find(&users)
	return users, result.RowsAffected, result.Error
}

// -------
// Read
// -------
func FindUserById(id string) (entity.UserAccount, error) {
	var user entity.UserAccount

	if err := database.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// -------
// Update
// -------
func UpdateUserById(input entity.UserAccount, id string) error {

	var user entity.UserAccount
	if err := database.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&user).Where("user_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete
// -------
func DeleteUserById(id string) error {

	var user entity.UserAccount
	if _, err := FindUserById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&user).Where("user_id = ?", id).Delete(user).Error; err != nil {
		return err
	}

	return nil
}
