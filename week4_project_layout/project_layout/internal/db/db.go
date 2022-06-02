package db

import (
	"database/sql"

	"project-layout/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewDb)

func NewDb(cfg *config.Config) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", cfg.Database.Dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
