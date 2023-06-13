package controller

import (
	"github.com/gofiber/fiber"
	"github.com/vanessadanu/Finpro-Golang.git/database"
	models "github.com/vanessadanu/Finpro-Golang.git/model"
)

func GetStaffAll(c *fiber.Ctx) error {
	var staffs []models.Staff

	database.DB.Find(&staffs)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get all staffs",
		"data":    staffs,
	})
}

func GetStaffById(c *fiber.Ctx) error {
	id := c.Params("id")

	var staff models.Staff

	database.DB.Find(&staff, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get staff by id",
		"data":    staff,
	})
}

func CreateStaff(c *fiber.Ctx) error {
	staff := new(models.Staff)

	staffInput := new(models.StaffInput)

	if err := c.BodyParser(staff); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Create(&staff)

	staffInput.Name = staff.Name
	staffInput.Email = staff.Email

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create staff",
		"data":    staffInput,
	})
}

func UpdateStaff(c *fiber.Ctx) error {
	id := c.Params("id")

	staff := new(models.Staff)

	staffInput := new(models.StaffInput)

	if err := c.BodyParser(staff); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Model(&staff).Where("id = ?", id).Updates(staff)

	staffInput.Name = staff.Name
	staffInput.Email = staff.Email

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Update staff",
		"data":    staffInput,
	})
}

func DeleteStaff(c *fiber.Ctx) error {
	id := c.Params("id")

	var staff models.Staff

	database.DB.First(&staff, id)
	if staff.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Staff not found",
		})
	}

	database.DB.Delete(&staff)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Delete staff",
	})
}
