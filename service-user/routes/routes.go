package routes

import (
	"service-user/config"
	"service-user/handler"
	"service-user/middleware"
	repository "service-user/repo"
	service "service-user/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := config.GetPostgresDatabase()

	userRepo := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepo)
	userController := handler.NewUserControllerImpl(userService)
	middleware := middleware.NewAuthImpl(userRepo)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from service-user")
	})
	app.Post("/user/register", userController.Register)
	app.Post("/user/login", userController.Login)
	app.Get("/user/auth", middleware.Authentication, userController.Auth)
}
