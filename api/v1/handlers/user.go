package handlers

import (
	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/bodatomas/gopi/utils"
	"github.com/gofiber/fiber/v2"
)

// Register user
// @Summary      Create new user
// @Description  Create new user in database
// @Tags         Users
// @Accept		 json
// @Produce      json
// @Success      200  {string} 	"User succesfully created."
// @Router       /user/register [post]
func HandleRegisterUser(c *fiber.Ctx) error {
	// Validate user input (email, password)
	registerInput, err := utils.CheckValidRegistrationInput(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(registerInput.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot hash password",
		})
	}

	// Store user data in the database
	db_err := models.CreateNewUser(registerInput.Email, hashedPassword)
	if db_err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot create user",
		})
	}

	// Return a success message
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success": "User successfully created",
	})
}
