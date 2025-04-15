package main

import (
	"log"

	"github.com/Satr10/1cak-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{})

	// middleware untuk log
	app.Use(logger.New())

	// TIDAK DIGUNAKAN DULU KARENA 0 USER NGAPAIN DI LIMIT WKWKWKWK
	//middleware untuk rate limiting
	// app.Use(limiter.New(limiter.Config{
	// 	Next: func(c *fiber.Ctx) bool {
	// 		return c.IP() == "127.0.0.1"
	// 	},
	// 	Max:        20,
	// 	Expiration: 30 * time.Second,
	// 	KeyGenerator: func(c *fiber.Ctx) string {
	// 		return c.Get("x-forwarded-for")
	// 	},
	// 	LimitReached: func(c *fiber.Ctx) error {
	// 		return c.JSON(fiber.Map{
	// 			"Message": "You're Too Fast Pal",
	// 			"Status":  "Failed",
	// 		})
	// 	},
	// }))

	router.SetupRouter(app)

	log.Fatal(app.Listen(":5001"))
}
