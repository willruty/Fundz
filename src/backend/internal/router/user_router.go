package router

import (
	user_controller "fundz/internal/controller/user"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(rg *gin.RouterGroup) int {

	user := rg.Group("/user")
	{
		// === User CRUD ===
		user.POST("/auth/register", user_controller.Register)
		user.POST("/auth/login", user_controller.Login)

		// user.PUT("/auth/forgot-password", user_controller.ForgetPassword)  esqueci senha = manda email com token jwt
		// user.PUT("/auth/reset-password", user_controller.ResetPassword)  envia nova senha com token para verificacao e muda a senha hasheada no banco
		// user.PUT("/auth/change-password", user_controller.ChangePassword)  só se já estiver logado

		user.PUT("/", user_controller.UpdateUserById)
		user.DELETE("/:id", user_controller.DeleteUserById)
	}

	return 4
}
