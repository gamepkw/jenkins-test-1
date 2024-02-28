package main

import (
	"log"
	_handler "main/handlers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_ "time/tzdata"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	e := echo.New()
	_handler.NewHandler(e)
	log.Fatal(e.Start("0.0.0.0:8090"))
}
