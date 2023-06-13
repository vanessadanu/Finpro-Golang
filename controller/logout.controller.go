package controller

import (
	"github.com/gofiber/fiber"
	auth "github.com/vanessadanu/Finpro-Golang.git/auth"
)

func Logout(c *fiber.Ctx) error {
	// Get the token from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing authorization header",
		})
	}

	// Extract the token from the Authorization header
	tokenString := auth.ExtractToken(authHeader)

	// Invalidate the token by adding it to the blacklist
	err := auth.AddToBlacklist(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to logout",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
