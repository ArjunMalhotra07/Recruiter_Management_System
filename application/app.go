package application

import (
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{router: AppRoutes()}
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
