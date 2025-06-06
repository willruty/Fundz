package router

import (
	"fundz/internal/controller"
	"os"
	"strings"

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

	SetupFinanceRouter()

	v1 := route.Group("/Fundz")
	// vao ser 41 novos endpoints -> criar mais grupos para cada endpoint

	{

		user := v1.Group("/user")
		// === User ===
		user.POST("/register", controller.CreateUser)
		user.GET("/", controller.GetAllUsers)
		user.GET("/:id", controller.GetUserById)
		user.PUT("/", controller.UpdateUserById)
		user.DELETE("/:id", controller.DeleteUserById)

		// v1.POST("/login", user.LoginUserAccount)
		// v1.GET("/user/getdata", user.GetDataByJWT)
		// v1.GET("/dashboard", user.Dashboard)

	}

	return route
}
