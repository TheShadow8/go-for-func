package posts

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/TheShadow8/go-test-fiber/routes"
)

type postRoutes struct {
	postController PostController
}

func NewPostRoutes(postController PostController) routes.Routes {
	return &postRoutes{postController}
}

func (r *postRoutes) Install(app *fiber.App) {
	app.Post("/post", r.postController.Save)
	app.Get("/posts", r.postController.GetAll)
}
