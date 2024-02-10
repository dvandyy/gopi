package handlers

import (
	"net/http"

	"github.com/dvandyy/gopi/api/v1/models"
	"github.com/dvandyy/gopi/config"
	"github.com/dvandyy/gopi/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Get user with UUID
// @Description  Return user with unique id
// @Tags         Users
// @Security 	 JWT_TOKEN
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

// @Summary      Register new user
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
		return c.JSON(models.Error{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(registerInput.Password)
	if err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Error while creating user",
		})
	}

	// Generate unique id with prefix
	uuid := utils.GenerateUniqueID("u-")

	// Store user data in the database
	db_err := models.CreateNewUser(uuid, registerInput.Email, hashedPassword)
	if db_err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Error while creating user",
		})
	}

	// Return a success message
	return c.JSON(models.RegisterResponse{
		Status:  fiber.StatusAccepted,
		Message: "User successfully created",
	})
}

// @Summary      Login user
// @Description  Authenticate user with jwt token
// @Tags         Users
// @Accept		 json
// @Produce      json
// @Param		 LoginRequest body models.LoginRequest true "Login user with Email and Password."
// @Success      200  {object} models.LoginResponse
// @Router       /users/login [post]
func HandleLoginUser(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)

	// Check if input is valid
	if err := c.BodyParser(loginRequest); err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Get user info from db
	userCredentials, err := models.GetUserCredentials(loginRequest.Email)
	if err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusUnauthorized,
			Message: "Invalid credentials",
		})
	}

	// Check if password is correct
	if err := utils.ComparePasswords(loginRequest.Password, userCredentials.Password); err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusUnauthorized,
			Message: "Invalid credentials",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(userCredentials.Id, loginRequest.Email, config.Get().JWT_Secret)
	if err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	// Return a success message
	return c.JSON(models.LoginResponse{
		Message: "User successfully logged in",
		Token:   token,
	})
}
