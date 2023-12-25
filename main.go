package main

import (
	"kuverse/config"
	"kuverse/controllers"
	"kuverse/middleware"
	"kuverse/routers"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadEnv()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))
	app.Use(middleware.ApiKeyMiddleware)
	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Request: %s %s", c.Method(), c.Path())
		return c.Next()
	})

	limiterMiddleware := limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
	})

	controllers.InitMongoDB()
	app.Use(limiterMiddleware)
	routers.SetupRoutes(app)
	app.Listen(":" + config.GetPort())
}
