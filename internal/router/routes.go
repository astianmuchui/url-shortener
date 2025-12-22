package router

import (
	"time"

	"github.com/astianmuchui/url-shortener/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Begin(app *fiber.App) {

	app.Get("/:code", handlers.HomeHandler)
	app.Route("/api/v1", func(router fiber.Router) {

		router.Use(cache.New(cache.Config{
			Expiration: 30 * time.Second,
		}))

		router.Use(limiter.New())

		router.Post("/url/register", handlers.RegisterURLHandler)

	})
}
