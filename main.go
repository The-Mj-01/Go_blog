package main

import (
	"blog/src/database"
	"blog/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./src/views/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	database.ConnectDb()
	Session := database.Instance
	Session.SetSession()

	routes.SetUpRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
