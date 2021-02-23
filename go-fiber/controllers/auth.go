package controllers

import (
	"fmt"
	"log"
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

	user, err := c.authServices.SignUp(&newUser)

	if err != nil {
		return util.NewAppError(err, http.StatusUnprocessableEntity)
	}

	return ctx.
		Status(http.StatusCreated).
		JSON(util.NewJResponse(nil, user.SanitizeUser()))

}

func (c *authController) SignIn(ctx *fiber.Ctx) error {
	var input models.User

	err := ctx.BodyParser(&input)

	if err != nil {
 	return util.NewAppError( err ,http.StatusUnprocessableEntity)
	}

	input.Email = util.NormalizeEmail(input.Email)
	user, err := c.authServices.GetByEmail(input.Email)

	if err != nil {
		log.Printf("%s signin get email failed: %v\n", input.Email, err.Error())
		return util.NewAppError(util.ErrInvalidCredentials, http.StatusUnauthorized)
	}
	err = util.VerifyPassword(user.Password, input.Password)
	if err != nil {
		log.Printf("%s signin verify password failed: %v\n", input.Email, err.Error())
		return util.NewAppError(util.ErrInvalidCredentials, http.StatusUnauthorized)
	}

	token, err := util.NewToken(user.ID.String())

	if err != nil {
		log.Printf("%s signin create token failed: %v\n", input.Email, err.Error())
		return util.NewAppError(err, http.StatusUnauthorized)
	}
	return ctx.
		Status(http.StatusOK).
		JSON(fiber.Map{
			"user":  user.SanitizeUser(),
			"token": fmt.Sprintf("Bearer %s", token),
		})

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
