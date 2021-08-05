package posts

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	posts "gitlab.com/TheShadow8/go-test-fiber/modules/posts/model"
	"gitlab.com/TheShadow8/go-test-fiber/util"
)

type PostController interface {
	Save(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type postController struct {
	postService PostService
}

func NewPostController(postService PostService) PostController {
	return &postController{postService}
}

func (c *postController) Save(ctx *fiber.Ctx) error {
	var newPost posts.Post

	err := ctx.BodyParser(&newPost)

	if err != nil {
		return util.NewAppError(err, http.StatusUnprocessableEntity)
	}

	_, err = c.postService.Save(&newPost)

	if err != nil {
		return util.NewAppError(err, http.StatusUnprocessableEntity)
	}

	return ctx.
		Status(http.StatusCreated).
		JSON(util.NewJResponse(nil, newPost))
}

func (c *postController) GetAll(ctx *fiber.Ctx) error {
	posts, err := c.postService.GetAll()

	if err != nil {
		return util.NewAppError(err, http.StatusNotFound)
	}

	return ctx.Status(http.StatusOK).JSON(util.NewJResponse(nil, &posts))

}
