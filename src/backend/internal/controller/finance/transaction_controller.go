
package controller

import (
	dao "fundz/internal/model/dao/finance"
	entity "fundz/internal/model/entity/finance"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------
// Create
// -------
func CreateTransaction(c *gin.Context) {

	var transaction entity.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreateTransaction(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

// -------
// GetAll
// -------
func GetAllTransactions(c *gin.Context) {

	transactions, rowsAffected, err := dao.FindAllTransactions()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      transactions,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(transactions),
		})
}

// -------
// GetById
// -------
func GetTransactionById(c *gin.Context) {

	result, err := dao.FindTransactionById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdateTransactionById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindTransactionById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var input entity.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdateTransactionById(input, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeleteTransactionById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindTransactionById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.DeleteTransactionById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
