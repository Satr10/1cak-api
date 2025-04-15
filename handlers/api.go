package handlers

import (
	"github.com/Satr10/1cak-api/scraper"
	"github.com/gofiber/fiber/v2"
)

func ApiIndex(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Halo",
		"status":  "Success",
		"NOTE":    "USE POST URL AS REFERER IN HEADER WHEN REQUESTING IMAGE",
	})
}

func RandomMeme(c *fiber.Ctx) error {
	scraper := scraper.NewScraper()
	meme, err := scraper.RandomMeme()
	if err != nil {
		return c.JSON(fiber.Map{
			"error":  err,
			"status": "Failed",
		})
	}
	return c.JSON(meme)
}

func PopularMeme(c *fiber.Ctx) error {
	scraper := scraper.NewScraper()
	memes, err := scraper.PopularMemeImages()
	if err != nil {
		return c.JSON(fiber.Map{
			"error":  err,
			"status": "Failed",
		})
	}

	return c.JSON(memes)
}
