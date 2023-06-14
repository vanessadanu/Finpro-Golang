package controller

import (
	"github.com/gofiber/fiber"
)

// func Logout(c *fiber.Ctx) error {
// 	// Extract the token from the Authorization header
// 	authHeader := c.Get("Authorization")
// 	if authHeader == "" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Missing authorization header",
// 		})
// 	}
// 	tokenString := auth.ExtractToken(authHeader)

// 	// Generate a new token with a short expiration time
// 	expiredToken, err := auth.GenerateExpiredToken(tokenString)
// 	if err != nil {
// 		// Log the error for debugging
// 		fmt.Println("Failed to generate expired token:", err)

// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to logout",
// 		})
// 	}

// 	// Set the new token in the response header to replace the existing token
// 	c.Set("Authorization", expiredToken)

// 	return c.JSON(fiber.Map{
// 		"message": "Logout successful",
// 	})
// }

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("jwt") // Clear the JWT cookie
	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
