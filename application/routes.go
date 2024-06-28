package application

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/ArjunMalhotra07/Recruiter_Management_System.git/handler/"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/handler"
	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func AppRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", DefaultRoute)
	fmt.Println("Running on Port 8080!")
	router.Route("/admin", adminRoutes)
	router.Route("/", BaseRoutes)
	return router
}
func BaseRoutes(router chi.Router) {
	router.Post("/signup", handler.SignUp)
}
func adminRoutes(router chi.Router) {
	// orderHandler := &handler.Order{}
	// router.Post("/", orderHandler.Create)
	// router.Get("/", orderHandler.List)
	// router.Get("/{id}", orderHandler.GetByID)
	// router.Put("/{id}", orderHandler.UpdateByID)
	// router.Delete("/{id}", orderHandler.DeleteByID)
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "pong🤣"}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the response as JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}
