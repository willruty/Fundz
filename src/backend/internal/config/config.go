package config

import (
	"fmt"
	"os"
	"path/filepath"

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
	SSlMode      string
}

type Jwt struct {
	JWTSECRET string
}

type ConfigEnv struct {
	Service  service
	Database database
	Jwt      Jwt
}

var Env ConfigEnv

func Load() error {

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(exe)

	fileConf := filepath.Join(exPath, "fundz.conf")
	cfg, err := ini.Load(fileConf)

	if err != nil {
		fmt.Println("Não foi possível encontrar ou abrir o arquivo: " + fileConf)
		return err
	}

	env := ConfigEnv{}

	// Port
	env.Service.Port = cfg.Section("SERVICE").Key("PORT").MustInt(0000)

	// Database
	env.Database.Host = cfg.Section("DATABASE").Key("HOST").String()
	env.Database.Port = cfg.Section("DATABASE").Key("PORT").MustInt(0000)
	env.Database.User = cfg.Section("DATABASE").Key("USER").String()
	env.Database.Password = cfg.Section("DATABASE").Key("PASSWORD").String()
	env.Database.DatabaseName = cfg.Section("DATABASE").Key("DBNAME").String()
	env.Database.DatabaseName = cfg.Section("DATABASE").Key("SSLMODE").String()
	
	env.Jwt.JWTSECRET = cfg.Section("JWT").Key("JWTSECRET").String()

	Env = env

	return nil

}
