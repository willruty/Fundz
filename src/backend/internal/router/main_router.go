package router

import (
	"os"
	"strings"

	"fundz/internal/controller"
	"fundz/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func configRouter() cors.Config {
	config := cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PUT"}
	config.ExposeHeaders = []string{"Origin", "Content-Type"}
	config.AllowCredentials = true
	return config
}

func SetupMainRouter() *gin.Engine {

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

	main := route.Group("/fundz")
	main.GET("/heath", controller.GetHealth)

	SetupUserRouter(main)
	SetupAcademicRouter(main)
	SetupFunRouter(main)
	SetupFinanceRouter(main)

	service.PrintBanner()

	return route
}
