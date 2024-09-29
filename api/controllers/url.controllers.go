package controllers

import (
	"errors"
	"url-shortner/api/config"
	"url-shortner/api/models"
	"url-shortner/api/types"
	"url-shortner/api/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShortenUrl(c *fiber.Ctx) error {
	// Parse the request body into the clientUrl struct
	clientUrl := types.Url{}
	err := c.BodyParser(&clientUrl)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
	}

	// Generate a unique short ID
	var new_id string
	if err := utils.IdCreator(&new_id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate short ID")
	}

	// Create a new Url record
	url := models.Url{
		LongUrl: clientUrl.LongUrl,
		ShortId: new_id,
	}

	// Insert the new URL record into the database
	result := config.DB.Create(&url)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save URL to the database")
	}

	// Return success message
	return c.SendString("URL created with short ID: " + new_id)
}

func RedirectToUrl(c *fiber.Ctx) error {

	short_id := c.Params("id")
	if len(short_id) != 10 {
		return c.Status(fiber.StatusBadRequest).SendString("invalid short url")
	}
	var url models.Url
	result := config.DB.Where("short_id = ?", short_id).First(&url)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// If there's a database error, return it
		return c.Status(fiber.StatusBadRequest).SendString("short url doesnt exist")
	}

	return c.Redirect(url.LongUrl, fiber.StatusPermanentRedirect)
}
