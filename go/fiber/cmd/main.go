package main

import (
	"os"

	"github.com/ygunayer/restbench/internal/cassandra"
	"github.com/ygunayer/restbench/internal/config"
	"github.com/ygunayer/restbench/internal/database"
	"github.com/ygunayer/restbench/internal/logger"
	"github.com/ygunayer/restbench/internal/server"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
		os.Exit(-1)
		return
	}

	if err := database.Open(&cfg.Database); err != nil {
		logger.Fatalf("Failed to connect to the database: %v", err)
		os.Exit(-1)
		return
	}

	if err := cassandra.Open(&cfg.Cassandra); err != nil {
		logger.Fatalf("Failed to connect to Cassandra: %v", err)
		os.Exit(-1)
		return
	}

	logger.Trace("Successfully connected to the database")

	srv := server.New(cfg)

	srv.OnShutdown(func() error {
		cassandra.Get().Close()
		logger.Trace("Successfully closed Cassandra connection")
		return nil
	})

	srv.OnShutdown(func() error {
		database.Get().Close()
		logger.Trace("Successfully closed PostgreSQL connection")
		return nil
	})

	if err := srv.Run(); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
		os.Exit(-1)
		return
	}

	logger.Tracef("Server is now listening at %v", cfg.Http.GetListenAddr())
}
