package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/TheShadow8/go-test-fiber/controllers"
)

type authRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) Routes {
	return &authRoutes{authController}
}

func (r *authRoutes) Install(app *fiber.App) {
	app.Post("/signup", r.authController.SignUp)
	app.Post("/signin", r.authController.SignIn)
	app.Get("/users", r.authController.GetUsers)
	app.Get("/users/:id", AuthRequired, r.authController.GetUser)
	// app.Put("/users/:id", AuthRequired, r.authController.PutUser)
	// app.Delete("/users/:id", AuthRequired, r.authController.DeleteUser)
}
