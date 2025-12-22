package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/astianmuchui/url-shortener/internal/env"
	"github.com/astianmuchui/url-shortener/internal/router"
	"github.com/astianmuchui/url-shortener/internal/scripts"
)

func init() {
	env.Load()
	scripts.Migrate()
}

func main() {

	app := fiber.New(fiber.Config{
		AppName: "URL Shortener Service",
		Prefork: true,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	router.Begin(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}


	app.Listen(fmt.Sprintf(":%s", port))
}
