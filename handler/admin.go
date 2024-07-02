package handler

import (
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/go-chi/chi/v5"
)

func (d *Env) PostJob(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "Created a new Job opening", Status: "Success"}
	SendResponse(w, response)
}

func (d *Env) GetJobDetails(w http.ResponseWriter, r *http.Request) {
	jobID := chi.URLParam(r, "job_id")
	var response models.Response

	if jobID != "" {
		response = models.Response{
			Message: "Job ID: " + jobID,
			Status:  "Success",
		}
	}
	SendResponse(w, response)
}
func (d *Env) GetAllApplicants(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "All users data", Status: "Success"}
	SendResponse(w, response)
}

func (d *Env) GetApplicantData(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "Data of a single user", Status: "Success"}
	SendResponse(w, response)
}
