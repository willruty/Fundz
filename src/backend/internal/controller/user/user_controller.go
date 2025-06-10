package user

import (
	dao "fundz/internal/repo/dao/user"
	entity "fundz/internal/repo/entity/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------
// Create
// -------
func CreateUser(c *gin.Context) {

	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// -------
// GetAll
// -------
func GetAllUsers(c *gin.Context) {

	users, rowsAffected, err := dao.FindAllUsers()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Nenhum registro encontrado: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"results":      users,
			"RowsAffected": rowsAffected,
			"RecordCount":  len(users),
		})
}

// -------
// GetById
// -------
func GetUserById(c *gin.Context) {

	result, err := dao.FindUserById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// -------
// UpdateById
// -------
func UpdateUserById(c *gin.Context) {

	var input entity.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.UpdateUserById(input, input.User_id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to update record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// -------
// DeleteById
// -------
func DeleteUserById(c *gin.Context) {

	if _, err := dao.FindUserById(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := dao.DeleteUserById(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
