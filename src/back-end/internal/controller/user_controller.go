package controller

// import "github.com/gin-gonic/gin"

import (
	"back-end/internal/config"
	"back-end/internal/model"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

func GenerateJWT(user_ID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user_ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.Env.Database.JwtSecret)
}

func CreateUser(c *gin.Context) {

	var usuario model.User

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	// verify if exists
	var existing model.User
	if err := model.DB.Where("user_email = ?", usuario.UserEmail).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email já está em uso"})
		return
	}

	// parse password to hash
	hashedPassword, err := HashPassword(usuario.UserPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao gerar hash > " + err.Error()})
		return
	}
	usuario.UserPassword = hashedPassword

	if err := model.DB.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao criar conta > " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Conta criada com sucesso!"})

}

func LoginUserAccount(c *gin.Context) {

	var req struct {
		Email    string `json:"user_email"`
		Password string `json:"user_password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var usuario model.User
	if err := model.DB.Where("user_email = ?", req.Email).
		First(&usuario).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Usuário ou senha inválidos"})
		return
	}

	if isOK, err := CheckPasswordHash(req.Password, usuario.UserPassword); !isOK {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "senha inválida " + err.Error()})
		return
	}

	token, err := GenerateJWT(usuario.UserID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao gerar JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetDataByJWT(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "token nao oferecido"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Formato do token inválido"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isOk := token.Method.(*jwt.SigningMethodHMAC); !isOk {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}

		return config.Env.Database.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Erro ao validar token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Token inválido"})
		return
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if !isOk || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Token mal formatado"})
		return
	}

	userID, isOK := claims["user_id"].(string)
	if !isOK {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	var usuario model.User
	if err := model.DB.First(&usuario, "user_id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":    usuario.UserID,
		"user_name":  usuario.UserName,
		"user_email": usuario.UserEmail,
	})

}
