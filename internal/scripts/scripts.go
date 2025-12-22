package scripts

import (
	"github.com/astianmuchui/url-shortener/internal/db"
	"github.com/astianmuchui/url-shortener/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetURL(c *fiber.Ctx) string {

	protocol := c.Protocol()
	host := c.Hostname()

	baseURL := protocol + "://" + host + "/"

	return baseURL
}

func Migrate() {
	log.Info("Running Migrations ......")

	db.DB.AutoMigrate(&models.Url{})
}
