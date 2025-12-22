package handlers

import (
	"github.com/astianmuchui/url-shortener/internal/models"
	"github.com/astianmuchui/url-shortener/internal/scripts"
	"github.com/gofiber/fiber/v2"
)

func HomeHandler(c *fiber.Ctx) error {
	code := c.Params("code")

	url := new(models.Url)
	url.ShortPath = code

	if err := url.Retreive(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Url not found",
		})
	}

	return c.Redirect(url.TargetPath)
}

func RegisterURLHandler(c *fiber.Ctx) error {

	payload := new(models.UrlCreateRequest)
	err := c.BodyParser(payload)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	url := new(models.Url)

	url.TargetPath = payload.TargetPath

	err = url.Create()
	path := scripts.GetURL(c)

	res := url.ToResponse(path)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unable to create",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
