package web

import (
	"encoding/json"
	DB "fibrapp/db"
	"fibrapp/models"
	"fibrapp/server"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSetupRoutes(t *testing.T) {
	app, _ := setupTest()

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t, resp.StatusCode, fiber.StatusOK, "status 200")
	assert.Equalf(t, "Hello Mundo", string(body), "same content")
}

func TestShowArticle(t *testing.T) {
	app, db := setupTest()

	article := models.Article{Title: "um", Content: "algo"}
	db.Create(&article)
	uri := fmt.Sprintf("/api/articles/%v", article.ID)

	req := httptest.NewRequest("GET", uri, nil)
	resp, _ := app.Test(req, -1)
	body, _ := ioutil.ReadAll(resp.Body)
	expected, _ := json.Marshal(&article)

	assert.Equalf(t, resp.StatusCode, fiber.StatusOK, "status 200")
	assert.Equalf(t, string(expected), string(body), "same content")
}

func TestListArticles(t *testing.T) {
	app, _ := setupTest()

	req := httptest.NewRequest("GET", "/api/articles", nil)
	resp, _ := app.Test(req, -1)
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t, resp.StatusCode, fiber.StatusOK, "status 200")
	assert.Equalf(t, "[]", string(body), "same content")
}

func setupTest() (*fiber.App, *gorm.DB) {
	config := server.ServerConfig{
		DBConfig: "host=localhost user=postgres password=postgres dbname=fibrapp_dev port=5432 sslmode=disable",
	}

	app := server.NewServer(config)
	SetupRoutes(app)
	db := DB.InitDB("host=localhost user=postgres password=postgres dbname=fibrapp_dev port=5432 sslmode=disable")
	db.Exec("DROP SCHEMA public CASCADE;")
	db.Exec("CREATE SCHEMA public;")
	DB.Migrate(db)

	return app, db
}
