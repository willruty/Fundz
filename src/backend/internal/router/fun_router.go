package router

import (
	controller "fundz/internal/controller/fun"

	"github.com/gin-gonic/gin"
)

func SetupFunRouter(rg *gin.RouterGroup) int {

	fun := rg.Group("/fun")
	{

		// === Player ===
		player := fun.Group("/fun")

		player.GET("/", controller.GetAllPlayers)

		{
			// === Player CRUD ===
			player.GET("/:id", controller.GetPlayerById)
			player.POST("/", controller.CreatePlayer)
			player.PUT("/", controller.UpdatePlayerById)
			player.DELETE("/:id", controller.DeletePlayerById)
		}

		// === Match ===
		match := fun.Group("/match")
		{
			// === Match CRUD ===
			match.GET("/", controller.GetAllMatches)
			match.GET("/:id", controller.GetMatchById)
			match.POST("/", controller.CreateMatch)
			match.PUT("/", controller.UpdateMatchById)
			match.DELETE("/:id", controller.DeleteMatchById)
		}

		// === GameType ===
		game_type := fun.Group("/game_type")
		{
			// === GameType CRUD ===
			game_type.GET("/", controller.GetAllGameTypes)
			game_type.GET("/:id", controller.GetGameTypeById)
			game_type.POST("/", controller.CreateGameType)
			game_type.PUT("/", controller.UpdateGameTypeById)
			game_type.DELETE("/:id", controller.DeleteGameTypeById)
		}
	}

	return 15
}
