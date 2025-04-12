package controller

import (
	"back-end/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var categoria model.Category

	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := model.DB.Create(&categoria).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "impossivel inserir categoria > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "categoria inserida com sucesso"})
}

func GetUserCategoriesByID(c *gin.Context) {

	var results []model.Category

	resultado := model.DB.Where("user_id = ?", c.Param("id")).Find(&results)
	if resultado.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": resultado.Error})
		return
	}

	if len(results) > 0 {
		c.IndentedJSON(http.StatusOK,
			gin.H{
				"RowsFound":  resultado.RowsAffected,
				"categories": results,
			})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "nenhuma categoria encontrada"})
		return
	}
}

func GetCategoryById(c *gin.Context) {

	var categoria model.Category

	err := model.DB.Where("category_id = ?", c.Param("id")).First(&categoria).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categoria})
}

func UpdateCategoryById(c *gin.Context) {

	var category model.Category

	if err := model.DB.Where("category_id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "categoria nao encontrada"})
		return
	}

	// Validate input
	var input model.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.DB.Model(&category).Where("category_id = ?", c.Param("id")).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "impossivel atualizar categoria > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated": category})
}

func DeleteCategoryById(c *gin.Context) {

	var category model.Category

	if err := model.DB.Where("category_id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := model.DB.Model(&category).Where("category_id = ?", c.Param("id")).Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "impossivel deletar categoria > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "categoria deletada com sucesso"})
}
