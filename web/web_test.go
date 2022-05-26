package web

import (
	"fibrapp/server"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	app := setupTest()

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t, resp.StatusCode, fiber.StatusOK, "status 200")
	assert.Equalf(t, "Hello Mundo", string(body), "same content")
}

func TestArticles(t *testing.T) {
	app := setupTest()

	req := httptest.NewRequest("GET", "/api/articles", nil)
	resp, _ := app.Test(req, -1)
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t, resp.StatusCode, fiber.StatusOK, "status 200")
	assert.Equalf(t, "[]", string(body), "same content")
}

func setupTest() *fiber.App {
	config := server.ServerConfig{
		DBConfig: "host=localhost user=postgres password=postgres dbname=fibrapp_dev port=5432 sslmode=disable",
	}

	app := server.NewServer(config)
	SetupRoutes(app)

	return app
}
