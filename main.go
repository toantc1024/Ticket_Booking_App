package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"tickets/configs"
	"tickets/database"
	"tickets/routes"
)

func init() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	database.ConnectDB(&config)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: false,
	}))

	routes.UserRoutes(micro, database.DB)
	routes.BookingRoutes(micro, database.DB)
	routes.TicketRoutes(micro, database.DB)
	routes.BookingRoutes(micro, database.DB)
	routes.PaymentRoutes(micro, database.DB)

	// Route GET healthcheck mặc định
	micro.Get("/check", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "GooPay Internship Mini-Project - ALIVE ❤️",
		})
	})

	log.Fatal(app.Listen(":8000"))

}
