package response

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ygunayer/restbench/internal/logger"
)

type ApiError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	statusCode int
	cause      error
}

type ErrorResponse struct {
	Error *ApiError `json:"error"`
}

func (e *ApiError) WithCause(c error) *ApiError {
	e.cause = c
	return e
}

func (e *ApiError) Respond(c *fiber.Ctx) error {
	return c.Status(e.statusCode).JSON(ErrorResponse{Error: e})
}

func InferError(e error) *ApiError {
	var apiError *ApiError

	if errors.As(e, &apiError) {
		return apiError
	}

	var fiberError *fiber.Error
	if errors.As(e, &fiberError) {
		return &ApiError{Code: "generic.unknown", Message: fiberError.Message, statusCode: fiberError.Code}
	}

	return &ApiError{Code: "generic.unknown", Message: "An unknown error has occurred", statusCode: 500}
}

func SendError(c *fiber.Ctx, e error) error {
	logger.Error(e)
	return InferError(e).Respond(c)
}

func (e *ApiError) Error() string {
	if e.cause == nil {
		return e.Message
	}

	return fmt.Sprintf("%s (Caused by: %v)", e.Message, e.cause)
}

func Unknown() *ApiError {
	return &ApiError{Code: "generic.unknown", Message: "An unknown error has occurred", statusCode: 500}
}

func NotFound(code string, message string) *ApiError {
	return &ApiError{Code: code, Message: message, statusCode: 404}
}

func BadRequest(code string, message string) *ApiError {
	return &ApiError{Code: code, Message: message, statusCode: 400}
}

func Forbidden(code string, message string) *ApiError {
	return &ApiError{Code: code, Message: message, statusCode: 403}
}
