package handler

import (
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
)

func (env *Env) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "All Jobs data here", Status: "Success"}
	SendResponse(w, response)
}

func (env *Env) ApplyToJob(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "Applied to a job", Status: "Success"}
	SendResponse(w, response)
}
