package router

import (
	"os"
	"strings"

	"back-end/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func configRouter() cors.Config {
	config := cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PATCH", "PUT"}
	config.ExposeHeaders = []string{"file_name"}
	return config
}

func SetupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	var modeDebug bool = true
	if len(os.Args) > 1 {
		if strings.ToLower(os.Args[1]) == "-debug" {
			modeDebug = true
			gin.SetMode(gin.DebugMode)
		}
	}

	route := gin.New()
	route.Use(gin.Recovery())

	if modeDebug {
		route.Use(gin.Logger())
	}

	route.Use(cors.New(configRouter()))

	v1 := route.Group("/cashly")

	{
		// === UserAccount ===
		// v1.POST("/register", controller.authController)
		// v1.POST("/login", controller.authController)
		// v1.GET("/user/me", controller.userController)

		// v1.GET("/transaction", controller.GetTransactions)
		// v1.POST("/transaction", controller.PostTransaction)
		// v1.PUT("/transaction/:id", controller.UpdateTransactionById)
		// v1.DELETE("/transaction/:id", controller.DeleteTransactionById)

		// v1.GET("/categories", controller.GetCategories)
		// v1.POST("/categories", controller.PostCategorie)
		// v1.PUT("/categories/:id", controller.UpdateCategorieById)
		// v1.DELETE("/categories/:id", controller.DeleteCategorieById)

		// v1.GET("/goal", controller.GetGoals)
		// v1.POST("/goal", controller.PostGoal)
		// v1.PUT("/goal/:id", controller.UpdateGoalById)
		// v1.DELETE("/goal/:id", controller.DeleteGoalById)

		v1.GET("/user", controller.GetUsers)
		v1.GET("/user", controller.GetUserById)
		v1.POST("/register", controller.CreateUser)
		v1.PUT("/user/:id", controller.UpdateUserById)
		v1.PATCH("/user/status/:id", controller.UpdatePasswordById)
		v1.DELETE("/user/:id", controller.DeleteUserById)

	}

	return route
}
