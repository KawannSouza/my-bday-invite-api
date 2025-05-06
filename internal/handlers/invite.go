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
	Title 		string    `json:"title"`
	Description string 	  `json:"description"`
	EventDate 	string    `json:"event_date"`
	Background 	string 	  `json:"background"`
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

	eventDate, err := time.Parse("2006-01-02 15:04:05", input.EventDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error" : "Invalid event date format"})
	}

	invite := model.Invite{
		UserID: 	 userID,
		Title: 		 input.Title,
		Description: input.Description,
		EventDate:   eventDate,
		Background:  input.Background,
	}

	if err := db.DB.Create(&invite).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Internal server error"})
	}

	return c.JSON(http.StatusCreated, invite)
}

func ListUserInvites(c echo.Context) error {
	userIDStr := c.Get("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Internal server error"})
	}

	var invites []model.Invite
	if err := db.DB.Where("user_id = ?", userID).Find(&invites).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Internal server error"})
	}

	return c.JSON(http.StatusOK, invites)
}

func UpdateInvite(c echo.Context) error {
	inviteID := c.Param("id")
	userIDStr := c.Get("user_id").(string)

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Internal server error"})
	}

	var invite model.Invite
	if err := db.DB.Where("id = ? AND user_id = ?", inviteID, userID).First(&invite).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error" : "Invite not found"})
	}

	var input struct {
		Title 		string 	  `json:"title"`
		Description string 	  `json:"description"`
		EventDate 	time.Time `json:"event_date"`
		Background 	string 	  `json:"background"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error" : "Invalid input"})
	}

	invite.Title = input.Title
	invite.Description = input.Description
	invite.EventDate = input.EventDate
	invite.Background = input.Background

	if err := db.DB.Save(&invite).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Internal server error"})
	}

	return c.JSON(http.StatusOK, invite)
}