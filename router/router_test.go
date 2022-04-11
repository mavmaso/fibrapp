package router

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	app := fiber.New()
	SetupRoutes(app)

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t, resp.StatusCode, fiber.StatusOK, "status 200")
	assert.Equalf(t, string(body), "Hello Mundo", "same content")
}
