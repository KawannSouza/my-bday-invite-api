package main

import (
	"log"

	"github.com/KawannSouza/my-bday-invite-api/internal/config"
	"github.com/KawannSouza/my-bday-invite-api/internal/db"
	"github.com/KawannSouza/my-bday-invite-api/internal/handlers"
	"github.com/KawannSouza/my-bday-invite-api/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main()  {
	config.LoadEnv()
	db.Connect()
	db.Migrate()

	port := config.GetEnv("PORT", "8080")
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	authGroup := e.Group("/auth")
	authGroup.Use(utils.AuthMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "API is running 🎉")
	})

	e.POST("/invite/register", handlers.Register)
	e.POST("/invite/login", handlers.Login)

	e.GET("/invites/:id", handlers.GetInviteByCode)
	e.POST("/invites", handlers.ConfirmPresence)

	authGroup.GET("/invites", handlers.ListUserInvites)
	authGroup.GET("/invites/:id/confirmations", handlers.GetConfirmations)
	authGroup.POST("/invites", handlers.CreateInvite)
	authGroup.PUT("/invites/:id", handlers.UpdateInvite)
	authGroup.DELETE("/invites/:id", handlers.DeleteInvite)

	log.Printf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}