package controller

import (
	dao "fundz/internal/repo/dao/fun"
	entity "fundz/internal/repo/entity/fun"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------
// Create
// -------
func CreateGameType(c *gin.Context) {

	var gameType entity.GameType

	if err := c.ShouldBindJSON(&gameType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreateGameType(gameType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gameType,
	})
}

// -------
// GetAll
// -------
func GetAllGameTypes(c *gin.Context) {

	gameTypes, rowsAffected, err := dao.GetAllGameType()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      gameTypes,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(gameTypes),
		})
}

// -------
// GetById
// -------
func GetGameTypeById(c *gin.Context) {

	result, err := dao.GetGameTypeById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdateGameTypeById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.GetGameTypeById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var input entity.GameType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdateGameTypeById(input, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeleteGameTypeById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.GetGameTypeById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.DeleteGameTypeById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
