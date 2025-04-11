package main

import (
	"back-end/internal/config"
	"back-end/internal/model"
	"back-end/internal/router"
	"fmt"
	"strconv"

)

func main() {

	err := config.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo de configuração")
		panic(err.Error())
	}

	model.DatabaseConnect()
	port := strconv.Itoa(config.Env.Service.Port)

	r := router.SetupRouter()
	r.Run(":" + port)
}
