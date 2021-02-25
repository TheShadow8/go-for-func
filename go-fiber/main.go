package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"gitlab.com/TheShadow8/go-test-fiber/controllers"
	"gitlab.com/TheShadow8/go-test-fiber/db"
	"gitlab.com/TheShadow8/go-test-fiber/repository"
	"gitlab.com/TheShadow8/go-test-fiber/routes"
	"gitlab.com/TheShadow8/go-test-fiber/services"
	"gitlab.com/TheShadow8/go-test-fiber/util"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Panicln(err)
	}
}

func main() {

	conn := db.NewConnection()

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Statuscode defaults to 500
			code := fiber.StatusInternalServerError

			// Retreive the custom statuscode if it's an *util.AppError
			if e, ok := err.(*util.AppError); ok {
				code = e.GetStatus()
			}

			if err != nil {

				return ctx.Status(code).JSON(util.NewJResponse(err, nil))
			}

			// Return from handler
			return nil
		},
	})

	uploadPath := fmt.Sprintf("./uploads")

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World !!!")
	})

	app.Static("/uploads", uploadPath)

	filesRepo := repository.NewFilesRepository(conn)
	fileServices := services.NewFileService(filesRepo)
	fileController := controllers.NewFileController(fileServices)
	fileRoutes := routes.NewFileRoutes(fileController)
	fileRoutes.Install(app)

	usersRepo := repository.NewUsersRepository(conn)
	authServices := services.NewAuthServices(usersRepo)
	authController := controllers.NewAuthController(authServices)
	authRoutes := routes.NewAuthRoutes(authController)
	authRoutes.Install(app)

	log.Fatal(app.Listen(":3000"))
}
