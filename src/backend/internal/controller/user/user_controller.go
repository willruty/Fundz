package user

import (
	dao "fundz/internal/model/dao/user"
	usuario "fundz/internal/model/entity/user"
	"fundz/internal/service"
	"fundz/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// -------
// Create
// -------
func Register(c *gin.Context) {

	var user usuario.UserAccount

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao hashear a senha"})
		return
	}
	user.Password = hashedPassword
	user.User_id = uuid.NewString()

	if err := dao.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	token, err := service.GenerateJWT(user.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// -------
// Login
// -------
func Login(c *gin.Context) {

	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	user, err := dao.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Email ou senha inválidos"})
		return
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Email ou senha inválidos"})
		return
	}

	token, err := service.GenerateJWT(user.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func ValidateToken(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token válido",
		"user_id": userID,
	})
}

// -------
// UpdateById
// -------
func UpdateUserById(c *gin.Context) {

	var input usuario.UserAccount
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

	if err := dao.DeleteUserById(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete record " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "row deleted"})
}
