package server

import (
	"fibrapp/db"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	DBConfig string
}

func NewServer(config ServerConfig) *fiber.App {
	app := fiber.New()

	SetMiddlewares(app, config.DBConfig)

	return app
}

func SetMiddlewares(app *fiber.App, config string) {
	db := db.InitDB(config)
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})
}
