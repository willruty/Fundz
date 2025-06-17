package dao

import (
	"fmt"
	"fundz/internal/database"
	usuario "fundz/internal/repo/entity/user"

)

// -------
// Create
// -------
func CreateUser(user usuario.UserAccount) error {
	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUser() ([]usuario.UserAccount, int64, error) {

	var user []usuario.UserAccount

	result := database.DB.Model(&usuario.UserAccount{}).Find(&user)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return user, result.RowsAffected, nil
}

func GetUserById(pk string) (usuario.UserAccount, error) {
	var user usuario.UserAccount

	if err := database.DB.Where("user_id = ?", pk).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// -------
// Update
// -------
func UpdateUserById(userUpdated usuario.UserAccount, id string) error {

	query := database.DB.Model(&usuario.UserAccount{}).Where("user_id = ?", id).Updates(userUpdated)
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
func DeleteUserById(id string) error {
	var user usuario.UserAccount

	query := database.DB.Where("user_id = ?", id).Delete(user)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
