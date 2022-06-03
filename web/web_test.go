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
)

func TestSetupRoutes(t *testing.T) {
	app := setupTest()

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Server.Test(req, -1)
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t, resp.StatusCode, fiber.StatusOK, "status 200")
	assert.Equalf(t, "Hello Mundo", string(body), "same content")
}

func TestShowArticle(t *testing.T) {
	app := setupTest()

	article := models.Article{Title: "um", Content: "algo"}
	app.DB.Create(&article)
	uri := fmt.Sprintf("/api/articles/%v", article.ID)

	req := httptest.NewRequest("GET", uri, nil)
	resp, _ := app.Server.Test(req, -1)
	body, _ := ioutil.ReadAll(resp.Body)
	expected, _ := json.Marshal(&article)

	assert.Equalf(t, resp.StatusCode, fiber.StatusOK, "status 200")
	assert.Equalf(t, string(expected), string(body), "same content")
}

func TestListArticles(t *testing.T) {
	app := setupTest()

	req := httptest.NewRequest("GET", "/api/articles", nil)
	resp, _ := app.Server.Test(req, -1)
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t, resp.StatusCode, fiber.StatusOK, "status 200")
	assert.Equalf(t, "[]", string(body), "same content")
}

func setupTest() server.App {
	config := server.ServerConfig{
		DBConfig: "host=localhost user=postgres password=postgres dbname=fibrapp_dev port=5432 sslmode=disable",
	}

	app := server.NewServer(config)
	SetupRoutes(app.Server)
	app.DB.Exec("DROP SCHEMA public CASCADE;")
	app.DB.Exec("CREATE SCHEMA public;")
	DB.Migrate(app.DB)

	return app
}
