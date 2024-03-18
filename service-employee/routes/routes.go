package routes

import (
	h "service-employee/handler"
	"service-employee/repo"
	service "service-employee/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App) {

	employeHandler := h.NewEmployeeControllerImpl(service.NewEmployeeServiceImpl(repo.New(&gorm.DB{})))

	app.Post("/employee", employeHandler.CreateEmployee)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from service-employee")
	})
}
