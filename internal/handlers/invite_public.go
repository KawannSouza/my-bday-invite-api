package handlers

import (
	"net/http"

	"github.com/KawannSouza/my-bday-invite-api/internal/db"
	"github.com/KawannSouza/my-bday-invite-api/internal/model"
	"github.com/labstack/echo/v4"
)

func GetInviteByCode(c echo.Context) error {
	inviteCode := c.Param("code")

	var invite model.Invite
	if err := db.DB.First(&invite, "invite_code = ?", inviteCode).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error" : "Invite not found"})
	}

	publicData := echo.Map{
		"id": invite.ID,
		"title": invite.Title,
		"description": invite.Description,
		"background": invite.Background,
		"created_at": invite.CreatedAt,
	}

	return c.JSON(http.StatusOK, publicData)
}