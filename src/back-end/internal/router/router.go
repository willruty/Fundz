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
		// === Health ===
		v1.GET("/health", controller.Health)

		// === UserAccount ===
		v1.POST("/register", controller.CreateUser)
		v1.POST("/login", controller.LoginUserAccount)
		v1.GET("/user/getdata", controller.GetDataByJWT)

		// === Transaction ===
		v1.GET("/user/transactions/:id", controller.GetUserTransactionsByID)
		v1.GET("/transaction/:id", controller.GetTransactionById)

		v1.POST("/transaction", controller.CreateTransaction)
		v1.PUT("/transaction/:id", controller.UpdateTransactionById)
		v1.DELETE("/transaction/:id", controller.DeleteTransactionById)

		// === Categories ===
		v1.GET("/user/category/:id", controller.GetUserCategoriesByID)
		v1.GET("/category/:id", controller.GetCategoryById)

		v1.POST("/category", controller.CreateCategory)
		v1.PUT("/category/:id", controller.UpdateCategoryById)
		v1.DELETE("/category/:id", controller.DeleteCategoryById)

		// === Goal ===
		v1.GET("/user/goal/:id", controller.GetUserGoalsById)
		v1.GET("/goal/:id", controller.GetGoalById)

		v1.POST("/goal", controller.CreateGoal)
		v1.PUT("/goal/:id", controller.UpdateGoalById)
		v1.DELETE("/goal/:id", controller.DeleteGoalById)

	}

	return route
}
