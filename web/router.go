package web

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	router := app.Group("/")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Mundo")
	})

	api := router.Group("/api")

	api.Get("/articles", ListArticle)
	api.Get("/articles/:id", ShowArticle)
}
