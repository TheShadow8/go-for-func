package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/TheShadow8/go-test-fiber/models"
	"gitlab.com/TheShadow8/go-test-fiber/services"
	"gitlab.com/TheShadow8/go-test-fiber/util"
)

var uploadPath = fmt.Sprintf("./uploads")

type FileController interface {
	Uploads(ctx *fiber.Ctx) error
	GetFile(ctx *fiber.Ctx) error
}

type fileController struct {
	fileServices services.FileServices
}

func NewFileController(fileServices services.FileServices) FileController {
	return &fileController{fileServices}
}

func (c *fileController) Uploads(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["files"]

	var f = models.File{}

	err = ctx.BodyParser(&f)

	err = os.MkdirAll(uploadPath, os.ModePerm)

	if err != nil {
		return err
	}

	var filesData []*models.File

	for _, file := range files {
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

		var fileInfo models.File

		name := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		path := fmt.Sprintf("%s/%s", uploadPath, name)

		err := ctx.SaveFile(file, path)

		if err != nil {
			return err
		}

		fileInfo.Project = f.Project
		fileInfo.FileName = name
		fileInfo.FilePath = path[1:]

		fmt.Println("fI", fileInfo)

		filesData = append(filesData, &fileInfo)

	}

	_, err = c.fileServices.Save(filesData)

	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(util.NewJResponse(nil, filesData))

}

func (c *fileController) GetFile(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	file, err := c.fileServices.GetById(id)

	if err != nil {
		return util.NewAppError(err, http.StatusNotFound)
	}

	return ctx.
		Status(http.StatusOK).
		JSON(util.NewJResponse(nil, &file))
}
