package api

import (
	"net/http"

	"github.com/Satr10/1cak-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	app := fiber.New()
	// middleware untuk log
	app.Use(logger.New())

	// agar bisa diskses dimana saja
	app.Use(cors.New())

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
	return adaptor.FiberApp(app)
}
