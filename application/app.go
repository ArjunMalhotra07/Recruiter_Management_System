package application

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	router http.Handler
	driver *sql.DB
}

func New() *App {
	return &App{router: AppRoutes(), driver: &sql.DB{}}
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

func (a *App) ConnectTODB() error {
	var err error
	a.driver, err = sql.Open("mysql", "root:Witcher_Arjun7@tcp(127.0.0.1:3306)/New_DB")
	if err != nil {
		return err
	}
	defer a.driver.Close()
	pingErr := a.driver.Ping()
	if pingErr != nil {
		return pingErr
	}
	fmt.Println("Connected to DB")
	return nil
}
