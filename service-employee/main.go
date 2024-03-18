package main

import (
	"fmt"
	"service-employee/config"
	"service-employee/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config.NewPostgresDatabase()
}

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	port := 3002
	fmt.Printf("Service employee is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service employee: %v\n", err)
	}
}
