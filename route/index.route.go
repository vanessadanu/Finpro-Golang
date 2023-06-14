package route

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/vanessadanu/Finpro-Golang.git/controller"
)

const jwtSecret = "secret-key"

func authMiddleware() func(c *fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
		SigningKey: []byte(jwtSecret),
	})
}

func RouteInit(app *fiber.App) {
	customer := app.Group("api/customer")
	customer.Post("/login", func(c *fiber.Ctx) {
		controller.CustomerLogin(c)
	})
	customer.Use(authMiddleware())
	customer.Get("/", func(c *fiber.Ctx) {
		controller.GetCustomerAll(c)
	})
	customer.Get("/orders", func(c *fiber.Ctx) {
		controller.GetCustomerOrders(c)
	})
	customer.Get("/bill", func(c *fiber.Ctx) {
		controller.GetCustomerBill(c)
	})
	// customer.Get("/:id", func(c *fiber.Ctx) {
	// 	controller.GetCustomerById(c)
	// })
	customer.Post("/logout", func(c *fiber.Ctx) {
		controller.Logout(c)
	})

	staff := app.Group("api/admin")
	staff.Post("/login", func(c *fiber.Ctx) {
		controller.StaffLogin(c)
	})
	staff.Use(authMiddleware())

	// Customer
	staff.Get("/customer", func(c *fiber.Ctx) {
		controller.GetCustomerAll(c)
	})
	staff.Get("/customer/:id", func(c *fiber.Ctx) {
		controller.GetCustomerById(c)
	})
	staff.Post("/customer", func(c *fiber.Ctx) {
		controller.CreateCustomer(c)
	})
	staff.Put("/customer/:id", func(c *fiber.Ctx) {
		controller.UpdateCustomer(c)
	})
	staff.Delete("/customer/:id", func(c *fiber.Ctx) {
		controller.DeleteCustomer(c)
	})

	//staff
	staff.Get("/staff", func(c *fiber.Ctx) {
		controller.GetStaffAll(c)
	})
	staff.Get("/staff/:id", func(c *fiber.Ctx) {
		controller.GetStaffById(c)
	})
	staff.Post("/staff", func(c *fiber.Ctx) {
		controller.CreateStaff(c)
	})
	staff.Put("/staff/:id", func(c *fiber.Ctx) {
		controller.UpdateStaff(c)
	})
	staff.Delete("/staff/:id", func(c *fiber.Ctx) {
		controller.DeleteStaff(c)
	})

	// Bill
	staff.Get("/bill", func(c *fiber.Ctx) {
		controller.GetBillAll(c)
	})
	staff.Get("/bill/:id", func(c *fiber.Ctx) {
		controller.GetBillById(c)
	})
	staff.Post("/bill", func(c *fiber.Ctx) {
		controller.CreateBill(c)
	})
	staff.Put("/bill/:id", func(c *fiber.Ctx) {
		controller.UpdateBill(c)
	})
	staff.Delete("/bill/:id", func(c *fiber.Ctx) {
		controller.DeleteBill(c)
	})

	// Order
	staff.Get("/order", func(c *fiber.Ctx) {
		controller.GetOrderAll(c)
	})
	staff.Get("/order/:id", func(c *fiber.Ctx) {
		controller.GetOrderById(c)
	})
	staff.Post("/order", func(c *fiber.Ctx) {
		controller.CreateOrder(c)
	})
	staff.Put("/order/:id", func(c *fiber.Ctx) {
		controller.UpdateOrder(c)
	})
	staff.Delete("/order/:id", func(c *fiber.Ctx) {
		controller.DeleteOrder(c)
	})

	//Pickup
	staff.Get("/pickup", func(c *fiber.Ctx) {
		controller.GetPickupAll(c)
	})
	staff.Get("/pickup/:id", func(c *fiber.Ctx) {
		controller.GetPickupById(c)
	})
	staff.Post("/pickup", func(c *fiber.Ctx) {
		controller.CreatePickup(c)
	})
	staff.Put("/pickup/:id", func(c *fiber.Ctx) {
		controller.UpdatePickup(c)
	})
	staff.Delete("/pickup/:id", func(c *fiber.Ctx) {
		controller.DeletePickup(c)
	})

	// Logout
	staff.Post("/logout", func(c *fiber.Ctx) {
		controller.Logout(c)
	})
}
