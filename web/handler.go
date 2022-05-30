package web

import (
	"fibrapp/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ListArticle(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	articles := new([]models.Article)
	db.Find(&articles)

	return c.JSON(articles)
}

func ShowArticle(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	article := new(models.Article)
	db.First(&article, id)

	return c.JSON(article)
}
