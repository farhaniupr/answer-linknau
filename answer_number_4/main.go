package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	engine := echo.New()

	engine.POST("/login", HandlerLogin)
	engine.GET("/beranda", Beranda, HandlerJWT())

	err := engine.Start(":8080")
	if err != nil {
		log.Println(err.Error())
	}
}
