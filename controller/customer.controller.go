package controller

import (
	"github.com/gofiber/fiber"
	"github.com/vanessadanu/Finpro-Golang.git/database"
	models "github.com/vanessadanu/Finpro-Golang.git/model"
	"gorm.io/gorm"
)

func GetCustomerAll(c *fiber.Ctx) error {
	var customers []models.Customer

	database.DB.Find(&customers)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get all customers",
		"data":    customers,
	})
}

func GetCustomerById(c *fiber.Ctx) error {
	id := c.Params("id")

	var customer models.Customer

	database.DB.Find(&customer, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get customer by id",
		"data":    customer,
	})
}

func GetCustomerOrders(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order

	if err := database.DB.Preload("Customer").Preload("Staff").First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No orders found for the customer",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve the orders",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get order by id",
		"data":    order,
	})
}

func GetCustomerBill(c *fiber.Ctx) error {
	id := c.Params("id")

	var bill models.Bill

	err := database.DB.Preload("Order.Customer").Preload("Order.Staff").First(&bill, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Bill not found for the customer",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve the bill",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get bill by id",
		"data":    bill,
	})
}

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)

	customerInput := new(models.CustomerInput)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Create(&customer)

	customerInput.Name = customer.Name
	customerInput.Email = customer.Email
	customerInput.Address = customer.Address

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create customer",
		"data":    customerInput,
	})
}

func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	customer := new(models.Customer)

	customerInput := new(models.CustomerInput)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Model(&customer).Where("id = ?", id).Updates(customer)

	customerInput.Name = customer.Name
	customerInput.Email = customer.Email
	customerInput.Address = customer.Address

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Update customer",
		"data":    customerInput,
	})
}

func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	var customer models.Customer

	database.DB.First(&customer, id)
	if customer.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Customer not found",
		})
	}

	database.DB.Delete(&customer)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Delete customer",
	})
}
