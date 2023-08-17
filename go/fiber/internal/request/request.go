package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ygunayer/restbench/internal/response"
)

var validate *validator.Validate = validator.New()

func Parse[T interface{}](c *fiber.Ctx) (*T, error) {
	req := new(T)

	if err := c.BodyParser(req); err != nil {
		return nil, response.BadRequest("common.request.parseFailed", "Invalid request").WithCause(err)
	}

	if err := validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return nil, response.BadRequest("common.request.validationFailed", fmt.Sprintf("Request validation failed: %v", &validationErrors)).WithCause(err)
	}

	return req, nil
}
