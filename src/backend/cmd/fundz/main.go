package main

import (
	"fmt"
	"fundz/internal/config"
	"fundz/internal/database"
	"fundz/internal/router"
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

	r := router.SetupMainRouter()
	r.Run(":" + port)

}
