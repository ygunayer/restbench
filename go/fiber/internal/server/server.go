package server

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/ygunayer/restbench/internal/auth"
	"github.com/ygunayer/restbench/internal/config"
	"github.com/ygunayer/restbench/internal/logger"
	"github.com/ygunayer/restbench/internal/response"
	"github.com/ygunayer/restbench/internal/transcript"
)

type Server struct {
	cfg           *config.Config
	app           *fiber.App
	shutdownHooks []ShutdownHandler
	hookMutex     sync.Mutex
}

type ShutdownHandler func() error

func New(cfg *config.Config) Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: response.SendError,
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("{\"status\": \"healthy\"}")
	})

	v1Group := app.Group("/api/v1")
	auth.BindRoutes(v1Group)
	transcript.BindV1Routes(v1Group)

	v2Group := app.Group("/api/v2")
	transcript.BindV2Routes(v2Group)

	return Server{
		cfg:           cfg,
		app:           app,
		shutdownHooks: make([]ShutdownHandler, 0, 10),
		hookMutex:     sync.Mutex{},
	}
}

func (srv *Server) Run() error {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)

	go func() {
		sig := <-c

		logger.Tracef("Received signal %v, shutting down server gracefully", sig)

		srv.hookMutex.Lock()
		defer srv.hookMutex.Unlock()

		for _, handler := range srv.shutdownHooks {
			if err := handler(); err != nil {
				logger.Tracef("Failed to run a shutdown hook: %v", err)
			}
		}

		if err := srv.Shutdown(); err != nil {
			logger.Fatalf("Failed to shut down server gracefully: %v", err)
		}
	}()

	return srv.app.Listen(srv.cfg.Http.GetListenAddr())
}

func (s *Server) Shutdown() error {
	if s.app == nil {
		return nil
	}

	return s.app.Shutdown()
}

func (s *Server) OnShutdown(handler ShutdownHandler) {
	s.hookMutex.Lock()
	defer s.hookMutex.Unlock()
	s.shutdownHooks = append(s.shutdownHooks, handler)
}
