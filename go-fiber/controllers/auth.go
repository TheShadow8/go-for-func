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
	SignIn(ctx *fiber.Ctx) error
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

	err = c.authServices.SignUp(&newUser)

	if err != nil {
		return util.NewAppError(err, http.StatusUnprocessableEntity)
	}

	return ctx.
		Status(http.StatusCreated).
		JSON(util.NewJResponse(nil, newUser.SanitizeUser()))

}

func (c *authController) SignIn(ctx *fiber.Ctx) error {
	var input models.User

	err := ctx.BodyParser(&input)

	if err != nil {
		return util.NewAppError(err, http.StatusUnprocessableEntity)
	}

	user, token, err := c.authServices.SignIn(&input)

	if err != nil {
		return util.NewAppError(err, http.StatusUnauthorized)
	}

	resData := map[string]interface{}{
		"user":  user,
		"token": token,
	}

	return ctx.
		Status(http.StatusOK).
		JSON(util.NewJResponse(nil, resData))

}

func (c *authController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := c.authServices.GetUser(id)

	if err != nil {
		return util.NewAppError(err, http.StatusNotFound)
	}

	return ctx.
		Status(http.StatusOK).
		JSON(util.NewJResponse(nil, &user))
}
