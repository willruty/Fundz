package dao

import (
	"nds_service/internal/database"
	"nds_service/internal/entity"
)

// -------
// Create
// -------
func CreateTransaction(transaction entity.Transaction) error {
	if err := database.DB.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

// -------
// ReadAll
// -------
func FindAllTransactions() ([]entity.Transaction, int64, error) {

	var transactions []entity.Transaction
	var count int64

	if err := database.DB.Model(&entity.Transaction{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := database.DB.Order("transaction_id asc").Find(&transactions) 
	return transactions, result.RowsAffected, result.Error
}

// -------
// Read 
// -------
func FindTransactionById(id string) (entity.Transaction, error) {
	var transaction entity.Transaction

	if err := database.DB.Where("transaction_id = ?", id).First(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

// -------
// Update
// -------
func UpdateTransactionById(input entity.Transaction, id string) error {

	var transaction entity.Transaction
	if err := database.DB.Where("transaction_id = ?", id).First(&transaction).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&transaction).Where("transaction_id = ?", id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

// -------
// Delete 
// -------
func DeleteTransactionById(id string) error {

	var transaction entity.Transaction
	if _, err := FindTransactionById(id); err != nil {
		return err
	}

	if err := database.DB.Model(&transaction).Where("transaction_id = ?", id).Delete(transaction).Error; err != nil {
		return err
	}

	return nil
}
