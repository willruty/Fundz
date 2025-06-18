package router

import (
	user_controller "fundz/internal/controller/user"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(rg *gin.RouterGroup)  {

	user := rg.Group("/user")
	{
		// === User CRUD ===
		user.POST("/register", user_controller.CreateUser)
		user.GET("/login", user_controller.GetUserById)
		user.PUT("/", user_controller.UpdateUserById)
		user.DELETE("/:id", user_controller.DeleteUserById)
	}
	
}