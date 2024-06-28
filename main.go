package main

import (
	"fmt"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/application"
)

func main() {
	app := application.New()
	err := app.StartServer()

	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
