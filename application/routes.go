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
	router.Get("/profiles", env.GetAllProfiles)

	return router
}
func JobRoutes(router chi.Router, env *handler.Env) {
	router.Get("/", env.GetAllJobs)
	//! To be implemented
	router.Get("/apply", env.ApplyToJob)
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

/*
! Gin Routing
func main() {
    r := gin.Default()
    r.Use(gin.Logger())

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong🤣"})
    })

    auth := r.Group("/")
    {
        LoginRoutes(auth)
    }

    admin := r.Group("/admin")
    admin.Use(JwtVerify(apigateway.Secret))
    {
        AdminRoutes(admin)
    }

    jobs := r.Group("/jobs")
    jobs.Use(JwtVerify(apigateway.Secret))
    {
        JobRoutes(jobs)
    }

    r.POST("/uploadResume", JwtVerify(apigateway.Secret), env.UploadResume)
    r.GET("/profiles", env.GetAllProfiles)

    r.Run() // Listen and serve on 0.0.0.0:8080
}

func LoginRoutes(router *gin.RouterGroup) {
    router.POST("/signup", env.SignUp)
    router.POST("/login", env.LogIn)
}

func AdminRoutes(router *gin.RouterGroup) {
    router.POST("/job", env.PostJob)
    router.GET("/job/:job_id", env.GetJobDetails)
    router.GET("/applicants", env.GetAllApplicants)
    router.GET("/applicant/:applicant_id", env.GetApplicantData)
}

func JobRoutes(router *gin.RouterGroup) {
    router.GET("/", env.GetAllJobs)
    router.GET("/apply", env.ApplyToJob)
}
! Echo Routing

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"message": "pong🤣"})
    })

    auth := e.Group("/")
    LoginRoutes(auth)

    admin := e.Group("/admin")
    admin.Use(JwtVerify(apigateway.Secret))
    AdminRoutes(admin)

    jobs := e.Group("/jobs")
    jobs.Use(JwtVerify(apigateway.Secret))
    JobRoutes(jobs)

    e.POST("/uploadResume", env.UploadResume, JwtVerify(apigateway.Secret))
    e.GET("/profiles", env.GetAllProfiles)

    e.Start(":8080") // Listen and serve on 0.0.0.0:8080
}

func LoginRoutes(router *echo.Group) {
    router.POST("/signup", env.SignUp)
    router.POST("/login", env.LogIn)
}

func AdminRoutes(router *echo.Group) {
    router.POST("/job", env.PostJob)
    router.GET("/job/:job_id", env.GetJobDetails)
    router.GET("/applicants", env.GetAllApplicants)
    router.GET("/applicant/:applicant_id", env.GetApplicantData)
}

func JobRoutes(router *echo.Group) {
    router.GET("/", env.GetAllJobs)
    router.GET("/apply", env.ApplyToJob)
}
*/