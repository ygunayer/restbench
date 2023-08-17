package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ygunayer/restbench/internal/logger"
	"github.com/ygunayer/restbench/internal/request"
	"github.com/ygunayer/restbench/internal/response"
	"github.com/ygunayer/restbench/internal/user"
)

func handleRegister(c *fiber.Ctx) error {
	req, err := request.Parse[RegisterUserRequest](c)

	if err != nil {
		return err
	}

	user, err := user.RegisterUser(c.Context(), req.ToCommand())

	if err != nil {
		return err
	}

	logger.Tracef("User registered successfully: %v", user)

	return response.JustCreated(c)
}

func BindRoutes(p fiber.Router) {
	g := p.Group("auth")
	g.Post("/register", handleRegister)
}
