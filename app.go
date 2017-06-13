package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// App test
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize test
func (a *App) Initialize(user, password, dbname string) {}

// Run test
func (a *App) Run(addr string) {}
