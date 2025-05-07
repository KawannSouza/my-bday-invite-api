package handlers

import (
	"net/http"

	"github.com/KawannSouza/my-bday-invite-api/internal/db"
	"github.com/KawannSouza/my-bday-invite-api/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

func ConfirmPresence(c echo.Context) error {
	inviteCode := c.Param("code")

	var invite model.Invite
	if err := db.DB.First(&invite, "invite_code = ?", inviteCode).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error" : "Invite not found"})
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := c.Bind(&input); err != nil || input.Name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error" : "Invalid input"})
	}

	confirmation := model.Confirmation{
		ID: 	  uuid.New(),
		InviteID: invite.ID,
		Name:     input.Name,
	}

	if err := db.DB.Create(&confirmation).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Failed to save confirmation"})
	}

	return c.JSON(http.StatusCreated, confirmation)
}