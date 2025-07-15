package router

import (
	user_controller "fundz/internal/controller/user"
	"fundz/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(rg *gin.RouterGroup) {

	user := rg.Group("/user")
	{
		// === User CRUD ===
		user.POST("/auth/register", user_controller.Register)
		user.POST("/auth/login", user_controller.Login)
		// user.PUT("/auth/forgot-password", user_controller.ForgetPassword)  esqueci senha = manda email com token jwt
		// user.PUT("/auth/reset-password", user_controller.ResetPassword)  envia nova senha com token para verificacao e muda a senha hasheada no banco

		validatedUser := user.Group("/")
		validatedUser.Use(middleware.AuthMiddleware)
		{
			validatedUser.GET("auth/validate", user_controller.ValidateToken)
			// validatedUser.PUT("/auth/change-password", user_controller.ChangePassword)  só se já estiver logado
			validatedUser.PUT("/", user_controller.UpdateUserById)
			validatedUser.DELETE("/:id", user_controller.DeleteUserById)
		}
	}
}
