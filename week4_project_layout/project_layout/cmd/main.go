package main

import (
	"database/sql"
	"log"
)

type App struct {
	db *sql.DB
}

func NewApp(db *sql.DB) *App {
	return &App{
		db: db,
	}
}

func main() {
	app, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}

	var version string
	row := app.db.QueryRow("SELECT VERSION()")
	if err := row.Scan(&version); err != nil {
		log.Fatal(err)
	}
	log.Println(version)
}
