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
func CreateAssignment(c *gin.Context) {

	var assignment entity.Assignment

	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreateAssignment(assignment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": assignment,
	})
}

// -------
// GetAll
// -------
func GetAllAssignments(c *gin.Context) {

	assignments, rowsAffected, err := dao.FindAllAssignments()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      assignments,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(assignments),
		})
}

// -------
// GetById
// -------
func GetAssignmentById(c *gin.Context) {

	result, err := dao.FindAssignmentById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdateAssignmentById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindAssignmentById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var input entity.Assignment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdateAssignmentById(input, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeleteAssignmentById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindAssignmentById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.DeleteAssignmentById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
