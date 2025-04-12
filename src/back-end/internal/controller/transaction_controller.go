package controller

import (
	"back-end/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {

	var transacao model.Transaction

	if err := c.ShouldBindJSON(&transacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := model.DB.Create(&transacao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "impossivel inserir transacao > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "transacao inserida com sucesso"})
}

func GetUserTransactionsByID(c *gin.Context) {

	var results []model.Transaction

	resultado := model.DB.Where("user_id = ?", c.Param("id")).Find(&results)
	if resultado.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": resultado.Error})
		return
	}

	if len(results) > 0 {
		c.IndentedJSON(http.StatusOK,
			gin.H{
				"RowsFound": resultado.RowsAffected,
				"notas":     results,
			})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "informações de extrato não encontradas"})
		return
	}

}

func GetTransactionById(c *gin.Context) {

	var transacao model.Transaction

	resultado := model.DB.Where("transaction_id = ?", c.Param("id")).Find(&transacao)
	if resultado.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": resultado.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transacao})

}

func UpdateTransactionById(c *gin.Context) {

	var transacao model.Transaction

	if err := model.DB.Where("transaction_id = ?", c.Param("id")).First(&transacao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "transacao nao encontrada"})
		return
	}

	// Validate input
	var input model.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.DB.Model(&transacao).Where("transaction_id = ?", c.Param("id")).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "impossivel atualizar transacao > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated": transacao})
}

func DeleteTransactionById(c *gin.Context) {

	var transacao model.Transaction

	if err := model.DB.Where("transaction_id = ?", c.Param("id")).First(&transacao).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := model.DB.Model(&transacao).Where("transaction_id = ?", c.Param("id")).Delete(&transacao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "impossivel deletar transacao > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "transacao deletada com sucesso"})
}
