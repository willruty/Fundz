package dao

import (
	"fmt"
	database "fundz/internal/database"
	"fundz/internal/repo/entity/finance"
)

// -------
// Create
// -------
func CreateAccount(account finance.Account) error {
	if err := database.DB.Create(&account).Error; err != nil {
		return err
	}
	return nil
}

func GetAllAccount() ([]finance.Account, int64, error) {

	var account []finance.Account
	var count int64

	result := database.DB.Model(&finance.Account{}).Count(&count).Find(&account)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return account, result.RowsAffected, nil
}

func GetAccountById(pk string) (finance.Account, error) {
	var account finance.Account

	if err := database.DB.Where("act_id = ?", pk).First(&account).Error; err != nil {
		return account, err
	}

	return account, nil
}

// -------
// Update
// -------
func UpdateAccountById(financeUpdated finance.Account, id string) error {

	query := database.DB.Model(&finance.Account{}).Where("act_id = ?", id).Updates(financeUpdated)
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
func DeleteAccountById(id string) error {
	var account finance.Account

	query := database.DB.Where("act_id = ?", id).Delete(account)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
