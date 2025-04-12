package controller

import (
	"back-end/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGoal(c *gin.Context) {
	var meta model.Goal

	if err := c.ShouldBindJSON(&meta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := model.DB.Create(&meta).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "impossivel inserir meta > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "meta criada com sucesso"})
}

func GetUserGoalsById(c *gin.Context) {
	var results []model.Goal

	resultado := model.DB.Where("user_id = ?", c.Param("id")).Find(&results)
	if resultado.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": resultado.Error})
		return
	}

	if len(results) > 0 {
		c.IndentedJSON(http.StatusOK,
			gin.H{
				"RowsFound":  resultado.RowsAffected,
				"goals": results,
			})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "nenhuma meta encontrada"})
		return
	}
}

func GetGoalById(c *gin.Context) {

	var meta model.Goal

	err := model.DB.Where("goal_id = ?", c.Param("id")).First(&meta).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": meta})
}

func UpdateGoalById(c *gin.Context) {

	var meta model.Goal

	if err := model.DB.Where("category_id = ?", c.Param("id")).First(&meta).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "categoria nao encontrada"})
		return
	}

	// Validate input
	var input model.Goal
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.DB.Model(&meta).Where("goal_id = ?", c.Param("id")).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "impossivel atualizar meta > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated": meta})
}

func DeleteGoalById(c *gin.Context) {

	var meta model.Goal

	if err := model.DB.Where("goal_id = ?", c.Param("id")).First(&meta).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := model.DB.Model(&meta).Where("goal_id = ?", c.Param("id")).Delete(&meta).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "impossivel deletar meta > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "meta deletada com sucesso"})
}
