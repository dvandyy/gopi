package handlers

import (
	"net/http"

	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/bodatomas/gopi/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Get user with UUID
// @Description  Return user with unique id
// @Tags         Users
// @Produce      json
// @Param 		 uid path string true "User unique ID"
// @Success      200  {object}  models.User
// @Router       /users/{uid} [get]
func HandleGetUserByID(c *fiber.Ctx) error {
	// Get id param from url
	id := c.Params("uid")
	// Get user from database
	data, err := models.GetUserByID(id)
	if err != nil {
		return c.JSON(fiber.Map{"status": http.StatusNotFound})
	}
	return c.JSON(data)
}

// @Summary      Create new user
// @Description  Create new user in database
// @Tags         Users
// @Accept		 json
// @Produce      json
// @Param		 RegisterRequest body models.RegisterRequest true "Create user with Email and Password."
// @Success      200  {object} models.RegisterResponse
// @Router       /users/register [post]
func HandleRegisterUser(c *fiber.Ctx) error {
	// Validate user input (email, password)
	registerInput, err := utils.CheckValidRegistrationInput(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error while creating user",
		})
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(registerInput.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error while creating user",
		})
	}

	// Generate unique id with prefix
	uuid := utils.GenerateUniqueID("u-")

	// Store user data in the database
	db_err := models.CreateNewUser(uuid, registerInput.Email, hashedPassword)
	if db_err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error while creating user",
		})
	}

	// Return a success message
	response := models.RegisterResponse{
		Status:  fiber.StatusAccepted,
		Message: "User successfully created",
	}
	return c.JSON(response)
}
