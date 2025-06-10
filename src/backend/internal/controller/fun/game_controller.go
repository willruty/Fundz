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
func CreateMatch(c *gin.Context) {

	var Match entity.Match

	if err := c.ShouldBindJSON(&Match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreateMatch(Match); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Match,
	})
}

// -------
// GetAll
// -------
func GetAllMatchs(c *gin.Context) {

	Matchs, rowsAffected, err := dao.FindAllMatchs()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      Matchs,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(Matchs),
		})
}

// -------
// GetById
// -------
func GetMatchById(c *gin.Context) {

	result, err := dao.FindMatchById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdateMatchById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindMatchById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var input entity.Match
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdateMatchById(input, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeleteMatchById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindMatchById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.DeleteMatchById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
