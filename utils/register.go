package utils

import (
	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/gofiber/fiber/v2"
)

func CheckValidRegistrationInput(context *fiber.Ctx) (*models.RegisterRequest, error) {

	register := new(models.RegisterRequest)

	// Check if input is valid
	if err := context.BodyParser(register); err != nil {
		return nil, err
	}

	// Check for valid email
	if err := CheckValidEmail(register.Email); err != nil {
		return nil, err
	}

	// Check for strong password
	if err := CheckStrongPassword(register.Password); err != nil {
		return nil, err
	}

	return register, nil
}
