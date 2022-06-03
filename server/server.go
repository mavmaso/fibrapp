package server

import (
	"fibrapp/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ServerConfig struct {
	DBConfig string
}

type App struct {
	DB     *gorm.DB
	Server *fiber.App
}

func NewServer(config ServerConfig) App {
	app := App{
		Server: fiber.New(),
		DB:     db.InitDB(config.DBConfig),
	}

	SetMiddlewares(app.Server, app.DB)

	return app
}

func SetMiddlewares(app *fiber.App, db *gorm.DB) {
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})
}
