package dao

import (
	"fmt"
	database "fundz/internal/database"
	"fundz/internal/model/entity/finance"
)

// -------
// Create
// -------
func CreateTransaction(transaction finance.Transaction) error {
	if err := database.DB.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTransaction() ([]finance.Transaction, int64, error) {

	var transaction []finance.Transaction
	var count int64

	result := database.DB.Model(&finance.Transaction{}).Count(&count).Find(&transaction)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return transaction, result.RowsAffected, nil
}

func GetTransactionById(pk string) (finance.Transaction, error) {
	var transaction finance.Transaction

	if err := database.DB.Where("transaction_id = ?", pk).First(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

// -------
// Update
// -------
func UpdateTransactionById(financeUpdated finance.Transaction, id string) error {

	query := database.DB.Model(&finance.Transaction{}).Where("transaction_id = ?", id).Updates(financeUpdated)
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
func DeleteTransactionById(id string) error {
	var transaction finance.Transaction

	query := database.DB.Where("transaction_id = ?", id).Delete(transaction)
	if err := query.Error; err != nil {
		return err
	} else if query.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado")
	}

	return nil
}
