package router

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	user_controller "fundz/internal/controller/user"
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
	var modeDebug bool
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

	// SetupAcademicRouter()
	SetupFinanceRouter()
	// SetupFunRouter()

	v1 := route.Group("/fundz")

	{
		user := v1.Group("/user")
		{
			// === User CRUD ===
			user.POST("/register", user_controller.CreateUser)
			user.GET("/", user_controller.GetAllUsers)
			user.GET("/:id", user_controller.GetUserById)
			user.PUT("/", user_controller.UpdateUserById)
			user.DELETE("/:id", user_controller.DeleteUserById)
		}

		// v1.POST("/login", user.LoginUserAccount)
		// v1.POST("/register", user.RegisterUserAccount)

		// v1.GET("/user/getdata", user.GetDataByJWT)
	}

	return route
}
