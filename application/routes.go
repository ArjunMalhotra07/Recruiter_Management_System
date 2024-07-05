package application

import (
	"encoding/json"
	"net/http"

	apigateway "github.com/ArjunMalhotra07/Recruiter_Management_System/api_gateway"
	mymiddleware "github.com/ArjunMalhotra07/Recruiter_Management_System/my_middleware"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/handler"
	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AppRoutes(env *handler.Env) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", DefaultRoute)
	router.Route("/", func(r chi.Router) {
		LoginRoutes(r, env)
	})
	router.Route("/admin", func(r chi.Router) {
		r.Use(mymiddleware.JwtVerify(apigateway.Secret))
		AdminRoutes(r, env)
	})
	router.Route("/jobs", func(r chi.Router) {
		r.Use(mymiddleware.JwtVerify(apigateway.Secret))
		JobRoutes(r, env)
	})
	router.With(mymiddleware.JwtVerify(apigateway.Secret)).Post("/uploadResume", env.UploadResume)
	return router
}
func JobRoutes(router chi.Router, env *handler.Env) {
	router.Get("/", env.GetAllJobs)
	//! To be implemented
	router.Get("/apply?job_id={job_id}", env.ApplyToJob)
}
func AdminRoutes(router chi.Router, env *handler.Env) {
	router.Post("/job", env.PostJob)
	//! To Be implemented
	router.Get("/job/{job_id}", env.GetJobDetails)
	router.Get("/applicants", env.GetAllApplicants)
	router.Get("/applicant/{applicant_id}", env.GetApplicantData)
}
func LoginRoutes(router chi.Router, env *handler.Env) {
	router.Post("/signup", env.SignUp)
	router.Post("/login", env.LogIn)
}
func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "pong🤣"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
