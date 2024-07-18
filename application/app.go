package application

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/handler"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	router http.Handler
	driver *sql.DB
}

func New(driver *sql.DB) *App {
	var env *handler.Env
	var h = handler.Env{Driver: driver}
	env = &h
	//! Or -> env := &handler.Env{Driver: driver}
	return &App{router: AppRoutes(env), driver: driver}
}
func (a *App) StartServer() error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}
	fmt.Println("Starting server")
	//! Method 2 to Start Server
	err := http.ListenAndServe(":8080", server.Handler)
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	return nil
}
