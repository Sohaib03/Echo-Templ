package main

import (
	"fmt"

	"github.com/Sohaib03/GoTemplate/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Started")

	app := echo.New()

	userHandler := handler.UserHandler{}
	app.Use(withUser)
	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("user", "Sohaib")
		return next(c)
	}
}
