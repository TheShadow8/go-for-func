package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/TheShadow8/go-test-fiber/util"
	"net/http"
	jwtware "github.com/gofiber/jwt/v2"
)

type Routes interface {
	Install(app *fiber.App)
}

func AuthRequired(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:    util.JwtSecretKey,
		SigningMethod: util.JwtSigningMethod,
		TokenLookup:   "header:Authorization",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.
				Status(http.StatusUnauthorized).
				JSON(util.NewJResponse(err,nil))
		},
	})(ctx)
}
