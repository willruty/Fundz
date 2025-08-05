package controller

import (
	dao "fundz/internal/model/dao/fun"
	entity "fundz/internal/model/entity/fun"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------
// Create
// -------
func CreatePlayer(c *gin.Context) {

	var player entity.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreatePlayer(player); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": player,
	})
}

// -------
// GetAll
// -------
func GetAllPlayers(c *gin.Context) {

	players, rowsAffected, err := dao.GetAllPlayers()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      players,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(players),
		})
}

// -------
// GetById
// -------
func GetPlayerById(c *gin.Context) {

	result, err := dao.GetPlayerById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdatePlayerById(c *gin.Context) {

	var input entity.Player
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdatePlayerById(input, input.PlayerId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeletePlayerById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.GetPlayerById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.DeletePlayerById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
