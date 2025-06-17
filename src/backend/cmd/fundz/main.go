package main

import (
	"fmt"
	"fundz/internal/config"
	"fundz/internal/database"
	"fundz/internal/router"
	"fundz/internal/service"
	"strconv"
)

func main() {

	err := config.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo de configuração")
		panic(err.Error())
	}

	database.DatabaseConnect()
	port := strconv.Itoa(config.Env.Service.Port)

	service.PrintBanner()

	r := router.SetupRouter()
	r.Run(":" + port)

}
