package routes

import (
	"blog/src/controllers"
	"github.com/gofiber/fiber/v2"
)

//SetRoute config
func SetUpRoutes(app *fiber.App) {
	//Index page route
	app.Get("/", controllers.Index)

	//Register page route
	app.Get("/register", controllers.RegisterPage)
	//Register request route
	app.Post("/register", controllers.Register)

	//Login page route
	app.Get("/login", controllers.LoginPage)
}
