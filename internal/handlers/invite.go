package handlers

import (
	"net/http"
	"time"

	"github.com/KawannSouza/my-bday-invite-api/internal/db"
	"github.com/KawannSouza/my-bday-invite-api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateInviteInput struct {
	Title string `json:"title"`
	Description string `json:"description"`
	EventDate time.Time `json:"event_date"`
	Background string `json:"background"`
}

func CreateInvite(c echo.Context) error {
	var input CreateInviteInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error" : "Invalid input"})
	}

	userIDStr := c.Get("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Internal server error"})
	}

	invite := model.Invite{
		UserID: 	 userID,
		Title: 		 input.Title,
		Description: input.Description,
		EventDate:   input.EventDate,
		Background:  input.Background,
	}

	if err := db.DB.Create(&invite).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Internal server error"})
	}

	return c.JSON(http.StatusCreated, invite)
}