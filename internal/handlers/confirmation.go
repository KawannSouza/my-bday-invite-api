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

func GetConfirmations(c echo.Context) error {
	userIDStr := c.Get("user_id").(string)
	inviteID := c.Param("id")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Invalid user ID"})
	}

	var invite model.Invite
	if err := db.DB.First(&invite, "id = ?", inviteID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error" : "invite not found"})
	}

	if invite.UserID != userID {
		return c.JSON(http.StatusForbidden, echo.Map{"error" : "You do not have permission to access this invite"})
	}

	var confirmations []model.Confirmation
	if err := db.DB.Where("invite_id = ?", invite.ID).Find(&confirmations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Failed to retrieve confirmations"})
	}

	return c.JSON(http.StatusOK, confirmations)
}