package controller

import (
	"github.com/gofiber/fiber"
	"github.com/vanessadanu/Finpro-Golang.git/database"
	models "github.com/vanessadanu/Finpro-Golang.git/model"
)

func GetOrderAll(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("Customer").Preload("Order.Staff").Find(&orders)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get all orders",
		"data":    orders,
	})
}

func GetOrderById(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order

	database.DB.Preload("Customer").Preload("Order.Staff").Find(&order, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get order by id",
		"data":    order,
	})
}

func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)

	orderInput := new(models.OrderInput)

	if err := c.BodyParser(order); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Create(&order)

	orderInput.CustomerID = order.CustomerID
	orderInput.StaffID = order.StaffID
	orderInput.Item = order.Item
	orderInput.Quantity = order.Quantity
	orderInput.TotalPrice = order.TotalPrice

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create order",
		"data":    orderInput,
	})
}

func UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	order := new(models.Order)

	orderInput := new(models.OrderInput)

	if err := c.BodyParser(orderInput); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Model(&order).Where("id = ?", id).Updates(orderInput)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Update order",
		"data":    orderInput,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order

	database.DB.First(&order, id)
	if order.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Order not found",
		})
	}

	database.DB.Delete(&order)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Delete order",
	})
}
