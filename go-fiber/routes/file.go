package routes

import (
	"github.com/gofiber/fiber/v2"

	"gitlab.com/TheShadow8/go-test-fiber/controllers"
)

type fileRoutes struct {
	fileController controllers.FileController
}

func NewFileRoutes(fileController controllers.FileController) Routes {
	return &fileRoutes{fileController}
}

func (r *fileRoutes) Install(app *fiber.App) {
	app.Post("/uploads", r.fileController.Uploads)
	app.Get("/files/:id", AuthRequired, r.fileController.GetFile)
}
