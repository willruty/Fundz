package model

import (
	"back-end/internal/config"
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	HOST     = config.Env.Database.Host
	PORT     = config.Env.Database.Port
	USER     = config.Env.Database.User
	PASSWORD = config.Env.Database.Password
	DBNAME   = config.Env.Database.DatabaseName
)

var DB *gorm.DB

func DatabaseConnect() {

	gorm.DefaultTableNameHandler = func(db *gorm.DB, table string) string {
		return table
	}

	dsn := url.URL{
		User:     url.UserPassword(config.Env.Database.User, config.Env.Database.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", config.Env.Database.Host, config.Env.Database.Port),
		Path:     config.Env.Database.DatabaseName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	database, err := gorm.Open("postgres", dsn.String())

	if err != nil {
		panic(err.Error())
	}

	DB = database
	DB.SingularTable(true)

	DB.AutoMigrate(&User{})
	// DB.Debug()

}
