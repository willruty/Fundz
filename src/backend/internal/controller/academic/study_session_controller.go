package controller

import (
	dao "fundz/internal/repo/dao/academic"
	entity "fundz/internal/repo/entity/academic"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------
// Create
// -------
func CreateStudy_session(c *gin.Context) {

	var study_session entity.Study_session

	if err := c.ShouldBindJSON(&study_session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreateStudy_session(study_session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": study_session,
	})
}

// -------
// GetAll
// -------
func GetAllStudy_sessions(c *gin.Context) {

	study_sessions, rowsAffected, err := dao.FindAllStudy_sessions()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      study_sessions,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(study_sessions),
		})
}

// -------
// GetById
// -------
func GetStudy_sessionById(c *gin.Context) {

	result, err := dao.FindStudy_sessionById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdateStudy_sessionById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindStudy_sessionById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var input entity.Study_session
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdateStudy_sessionById(input, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeleteStudy_sessionById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindStudy_sessionById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.DeleteStudy_sessionById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
