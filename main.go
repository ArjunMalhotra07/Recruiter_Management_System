package main

import (
	"database/sql"
	"fmt"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/application"
)

func main() {
	driver, err := sql.Open("mysql", "root:Witcher_Arjun7@tcp(127.0.0.1:3306)/New_DB")
	if err != nil {
		fmt.Println("Can't connect to DB:", err)
		return
	}
	defer driver.Close()

	app := application.New(driver)

	err = app.StartServer()
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
}
