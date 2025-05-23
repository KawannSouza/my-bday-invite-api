package handlers

import (
	"net/http"

	"github.com/KawannSouza/my-bday-invite-api/internal/db"
	"github.com/KawannSouza/my-bday-invite-api/internal/model"
	"github.com/KawannSouza/my-bday-invite-api/internal/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

func Register (c echo.Context) error {
	var input RegisterInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error" : "Invalid input"})
	}

	if input.Username == "" || input.Email == "" || input.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "All fields are required"})
	}

	var existing model.User 
	if err := db.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return c.JSON(http.StatusConflict, echo.Map{"error" : "Email already exists"})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Failed to hash password"})
	}

	user := model.User{
		Username:  	  input.Username,
		Email: 	  	  input.Email,
		PasswordHash: string(hashed),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Failed to create user"})
	}

	user.PasswordHash = ""
	return c.JSON(http.StatusCreated, user)
}

type LoginInput struct {
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var input LoginInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error" : "Invalid data"})
	}

	var user model.User
	if err := db.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error" : "Invalid Credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Invalid Credentials"})
	}

	token, err := utils.GenerateJWT(user.ID.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error" : "Error generating token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}