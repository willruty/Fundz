package main

import (
	"fmt"
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
