package controller

import (
	"github.com/gofiber/fiber"
	"github.com/vanessadanu/Finpro-Golang.git/database"
	models "github.com/vanessadanu/Finpro-Golang.git/model"
)

func GetBillAll(c *fiber.Ctx) error {
	var bills []models.Bill

	database.DB.Preload("Order").Preload("Order.Customer").Preload("Order.Staff").Find(&bills)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get all bills",
		"data":    bills,
	})
}

func GetBillById(c *fiber.Ctx) error {
	id := c.Params("id")

	var bill models.Bill

	database.DB.Preload("Order").Preload("Order.Customer").Preload("Order.Staff").Find(&bill, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get bill by id",
		"data":    bill,
	})
}

func CreateBill(c *fiber.Ctx) error {
	bill := new(models.Bill)

	billInput := new(models.BillInput)

	if err := c.BodyParser(bill); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Create(&bill)

	billInput.Order_id = bill.Order_id
	billInput.Amount = bill.Amount
	billInput.Paid = bill.Paid

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create bill",
		"data":    billInput,
	})
}

func UpdateBill(c *fiber.Ctx) error {
	id := c.Params("id")

	bill := new(models.Bill)

	billInput := new(models.BillInput)

	if err := c.BodyParser(billInput); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Model(&bill).Where("id = ?", id).Updates(billInput)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Update bill",
		"data":    billInput,
	})

}

func DeleteBill(c *fiber.Ctx) error {
	id := c.Params("id")

	var bill models.Bill

	database.DB.First(&bill, id)
	if bill.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Bill not found",
		})
	}

	database.DB.Delete(&bill)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Delete bill",
	})
}
