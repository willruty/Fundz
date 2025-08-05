package controller

import (
	dao "fundz/internal/model/dao/academic"
	entity "fundz/internal/model/entity/academic"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------
// Create
// -------
func CreateSubject(c *gin.Context) {

	var subject entity.Subject

	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreateSubject(subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subject,
	})
}

// -------
// GetAll
// -------
func GetAllSubjects(c *gin.Context) {

	subjects, rowsAffected, err := dao.GetAllSubject()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      subjects,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(subjects),
		})
}

// -------
// GetById
// -------
func GetSubjectById(c *gin.Context) {

	result, err := dao.GetSubjectById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdateSubjectById(c *gin.Context) {

	var input entity.Subject
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdateSubjectById(input, input.SubjectId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeleteSubjectById(c *gin.Context) {

	if err := dao.DeleteSubjectById(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
