package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
	app.Get("/users/:userId", controllers.Userbuy)

	app.Post("/cart/:userId", controllers.CreateCart)
	app.Get("/carts", controllers.GetAllCarts)
	app.Delete("/cart/:userId/:nameproduct", controllers.DeleteACart)
	app.Put("/cart/:userId/:nameproduct", controllers.EditACart)
	app.Get("/cart/:userId/", controllers.GetACart)

}
