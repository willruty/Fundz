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
func CreateGame_type(c *gin.Context) {

	var game_type entity.Game_type

	if err := c.ShouldBindJSON(&game_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreateGame_type(game_type); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": game_type,
	})
}

// -------
// GetAll
// -------
func GetAllGame_types(c *gin.Context) {

	game_types, rowsAffected, err := dao.FindAllGame_types()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      game_types,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(game_types),
		})
}

// -------
// GetById
// -------
func GetGame_typeById(c *gin.Context) {

	result, err := dao.FindGame_typeById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdateGame_typeById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindGame_typeById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var input entity.Game_type
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdateGame_typeById(input, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeleteGame_typeById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindGame_typeById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.DeleteGame_typeById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
