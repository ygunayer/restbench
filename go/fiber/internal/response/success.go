package response

import "github.com/gofiber/fiber/v2"

const EMPTY = ""

type SuccessResponse struct {
	Data       interface{} `json:"data"`
	statusCode int
}

func respond(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(200).JSON(SuccessResponse{Data: data, statusCode: 200})
}

func respondEmpty(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).SendString(EMPTY)
}

func Ok(c *fiber.Ctx, data interface{}) error {
	return respond(c, 200, data)
}

func Created(c *fiber.Ctx, data interface{}) error {
	return respond(c, 201, data)
}

func Accepted(c *fiber.Ctx, data interface{}) error {
	return respond(c, 202, data)
}

func JustOk(c *fiber.Ctx) error {
	return respondEmpty(c, 200)
}

func JustCreated(c *fiber.Ctx) error {
	return respondEmpty(c, 201)
}

func JustAccepted(c *fiber.Ctx) error {
	return respondEmpty(c, 202)
}
