package main

import (
	"log"

	"github.com/KawannSouza/my-bday-invite-api/internal/config"
	"github.com/KawannSouza/my-bday-invite-api/internal/db"
	"github.com/KawannSouza/my-bday-invite-api/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main()  {
	config.LoadEnv()
	db.Connect()
	db.Migrate()

	port := config.GetEnv("PORT", "8080")
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "API is running 🎉")
	})

	e.POST("/invite/register", handlers.Register)
	e.POST("/invite/login", handlers.Login)

	log.Printf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}