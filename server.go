package main

import (
	"alexpy.com/julia/db"
	"alexpy.com/julia/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(handlers.LogRequest)

	// Routes
	e.GET("/", handlers.Home)
	db.InitDB()

	e.GET("/users", handlers.HandleGetAllUser)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/user/:id", handlers.HandleUpdateUser)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
