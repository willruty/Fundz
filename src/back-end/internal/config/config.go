package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type service struct {
	Port int
}

type database struct {
	Port         int
	Host         string
	DatabaseName string
	User         string
	Password     string
	JwtSecret    []byte
}

type ConfigEnv struct {
	Service  service
	Database database
}

var Env ConfigEnv

func Load() error {

	// exe, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// exPath := filepath.Dir(exe)

	fileConf := "D:/Cashly-BudgetFlow/Cashly/configs/cashly.conf"
	cfg, err := ini.Load(fileConf)

	if err != nil {
		fmt.Println("Não foi possível encontrar ou abrir o arquivo: " + fileConf)
		return err
	}

	env := ConfigEnv{}

	// Port
	env.Service.Port = cfg.Section("service").Key("port").MustInt(8080)

	// Database
	env.Database.Host = cfg.Section("database").Key("host").String()
	env.Database.Port = cfg.Section("database").Key("port").MustInt(5432)
	env.Database.User = cfg.Section("database").Key("user").String()
	env.Database.Password = cfg.Section("database").Key("password").String()
	env.Database.DatabaseName = cfg.Section("database").Key("dbname").String()
	env.Database.JwtSecret = []byte(cfg.Section("database").Key("jwtSecret").String())

	Env = env

	return nil
}
