package controller

import (
	dao "fundz/internal/model/dao/finance"
	entity "fundz/internal/model/entity/finance"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------
// Create
// -------
func CreateCategory(c *gin.Context) {

	var category entity.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreateCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

// -------
// GetAll
// -------
func GetAllCategorys(c *gin.Context) {

	categorys, rowsAffected, err := dao.FindAllCategorys()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      categorys,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(categorys),
		})
}

// -------
// GetById
// -------
func GetCategoryById(c *gin.Context) {

	result, err := dao.FindCategoryById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdateCategoryById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindCategoryById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var input entity.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdateCategoryById(input, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeleteCategoryById(c *gin.Context) {

	id := c.Param("id")

	if _, err := dao.FindCategoryById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.DeleteCategoryById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
