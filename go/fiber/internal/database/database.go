package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/ygunayer/restbench/internal/config"
)

var db *bun.DB

func Open(cfg *config.DatabaseConfig) error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.Url)))
	db = bun.NewDB(sqldb, pgdialect.New())
	return db.Ping()
}

func Get() *bun.DB {
	return db
}
