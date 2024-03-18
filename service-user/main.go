package main

import (
	"fmt"
	"service-user/routes"

	"service-user/config"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config.NewPostgresDatabase()
}

func main() {

	app := fiber.New()
	routes.SetupRoutes(app)

	port := 3001
	fmt.Printf("Service user is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service user: %v\n", err)
	}
}
