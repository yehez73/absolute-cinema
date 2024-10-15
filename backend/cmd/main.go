package main

import (
	database "backend/databases"
	"backend/routes"
	"backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Main setup
	e := routes.Route()
	database.Connect()
	database.InitCollection()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())
	customValidator := &utils.CustomValidator{Validator: validator.New()}
	e.Validator = customValidator
	e.Logger.Fatal(e.Start(":8080"))
}