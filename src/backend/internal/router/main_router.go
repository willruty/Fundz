package router

import (
	"fmt"
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
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PATCH", "PUT"}
	config.ExposeHeaders = []string{"file_name"}
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

	academicEndpointsCount := SetupAcademicRouter(main)
	funEndpointsCount := SetupFunRouter(main)
	financeEndpointsCount := SetupFinanceRouter(main)
	userEndpointsCount := SetupUserRouter(main)

	endpoints := fmt.Sprintf("Academic -> %d | ", academicEndpointsCount) +
		fmt.Sprintf("Fun -> %d | ", funEndpointsCount) +
		fmt.Sprintf("Finance -> %d | ", financeEndpointsCount) +
		fmt.Sprintf("User -> %d", userEndpointsCount)

	service.PrintBanner(endpoints)

	main.GET("/heath", controller.GetHealth)

	return route
}
