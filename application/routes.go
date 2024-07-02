package application

import (
	"encoding/json"
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/handler"
	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AppRoutes(env *handler.Env) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", DefaultRoute)
	router.Route("/admin", func(r chi.Router) {
		AdminRoutes(r, env)
	})
	router.Route("/", func(r chi.Router) {
		LoginRoutes(r, env)
	})
	router.Route("/", func(r chi.Router) {
		JobRoutes(r, env)
	})
	return router
}

func AdminRoutes(router chi.Router, env *handler.Env) {
	router.Post("/job", env.PostJob)
	router.Get("/job/{job_id}", env.GetJobDetails)
	router.Get("/job/applicants", env.GetAllApplicants)
	router.Get("/job/applicant/{applicant_id}", env.GetApplicantData)
}
func LoginRoutes(router chi.Router, env *handler.Env) {
	router.Post("/signup", env.SignUp)
	router.Post("/login", env.LogIn)
}
func JobRoutes(router chi.Router, env *handler.Env) {
	router.Get("/jobs", env.GetAllJobs)
	router.Get("/jobs/apply?job_id={job_id}", env.ApplyToJob)
}
func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "pong🤣"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
