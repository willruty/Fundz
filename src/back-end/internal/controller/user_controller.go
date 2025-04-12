package controller

// import "github.com/gin-gonic/gin"

import (
	"back-end/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var usuario model.User

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := model.DB.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao criar conta > " + err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": "Conta criada com sucesso!"})

}

func LoginUserAccount(c *gin.Context) {}

func GetDataByJWT(c *gin.Context) {}
