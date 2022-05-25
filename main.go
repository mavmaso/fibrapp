package main

import (
	"fibrapp/server"
	"fibrapp/web"
)

func main() {
	config := server.ServerConfig{
		DBConfig: "host=localhost user=postgres password=postgres dbname=fibrapp_dev port=5432 sslmode=disable",
	}

	app := server.NewServer(config)
	web.SetupRoutes(app)

	app.Listen(":3000")
}
