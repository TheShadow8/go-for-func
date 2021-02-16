package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/TheShadow8/go-test-fiber/models"
	"gitlab.com/TheShadow8/go-test-fiber/services"
	"gitlab.com/TheShadow8/go-test-fiber/util"
)

type AuthController interface {
	SignUp(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
}

type authController struct {
	authServices services.AuthServices
}

func NewAuthController(authServices services.AuthServices) AuthController {
	return &authController{authServices}
}

func (c *authController) SignUp(ctx *fiber.Ctx) error {
	var newUser models.User
	err := ctx.BodyParser(&newUser)

	if err != nil {
		return util.NewAppError(err, http.StatusUnprocessableEntity)
	}

	user, err := c.authServices.SignUp(&newUser)

	if err != nil {
		return util.NewAppError(err, http.StatusUnprocessableEntity)
	}

	return ctx.
		Status(http.StatusCreated).
		JSON(user)

}

func (c *authController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := c.authServices.GetUser(id)

	if err != nil {
		return util.NewAppError(err, http.StatusNotFound)
	}

	return ctx.
		Status(http.StatusOK).
		JSON(user)
}
