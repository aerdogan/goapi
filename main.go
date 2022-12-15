package main

import (
	"goapi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())     // cors
	app.Use(csrf.New())     // csrf
	app.Use(limiter.New())  // limiter
	app.Use(logger.New())   // logger
	app.Use(cache.New())    // cache
	app.Use(compress.New()) // compression

	routes.UserRoute(app)
	app.Listen(":3000")
}
