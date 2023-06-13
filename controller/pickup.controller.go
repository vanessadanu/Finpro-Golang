package controller

import (
	"github.com/gofiber/fiber"
	"github.com/vanessadanu/Finpro-Golang.git/database"
	models "github.com/vanessadanu/Finpro-Golang.git/model"
)

func GetPickupAll(c *fiber.Ctx) {
	var pickups []models.Pickup

	database.DB.Preload("Order").Preload("Order.Customer").Find(&pickups)

	c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get all pickups",
		"data":    pickups,
	})

}

func GetPickupById(c *fiber.Ctx) {
	id := c.Params("id")

	var pickup models.Pickup

	database.DB.Preload("Order").Preload("Order.Customer").Find(&pickup, id)

	c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get pickup by id",
		"data":    pickup,
	})
}

func CreatePickup(c *fiber.Ctx) {
	pickup := new(models.Pickup)

	pickupInput := new(models.PickupInput)

	if err := c.BodyParser(pickup); err != nil {
		c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Create(&pickup)

	pickupInput.OrderID = pickup.OrderID
	pickupInput.Date = pickup.Date

	c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create pickup",
		"data":    pickupInput,
	})
}

func UpdatePickup(c *fiber.Ctx) {
	id := c.Params("id")

	pickup := new(models.Pickup)

	pickupInput := new(models.PickupInput)

	if err := c.BodyParser(pickupInput); err != nil {
		c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Model(&pickup).Where("id = ?", id).Updates(pickupInput)

	c.JSON(fiber.Map{
		"status":  "success",
		"message": "Update pickup",
		"data":    pickupInput,
	})
}

func DeletePickup(c *fiber.Ctx) {
	id := c.Params("id")

	var pickup models.Pickup

	database.DB.First(&pickup, id)
	if pickup.ID == 0 {
		c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Pickup not found",
			"data":    nil,
		})
		return
	}

	database.DB.Delete(&pickup)

	c.JSON(fiber.Map{
		"status":  "success",
		"message": "Delete pickup",
		"data":    pickup,
	})
}
