package handler

import (
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
)

func (d *Env) GetJobDetails(w http.ResponseWriter, r *http.Request) {
	//! Get Claim Values and job_id
	claims := r.Context().Value("claims").(jwt.MapClaims)
	isAdmin := claims["is_admin"].(bool)
	if !isAdmin {
		SendResponse(w, models.Response{
			Message: "Only Admins can get details of this job", Status: "Error", Data: ""})
		return
	}
	jobID := chi.URLParam(r, "job_id")
	var response models.Response
	var jobExists bool
	err := d.Driver.QueryRow("SELECT EXISTS(SELECT 1 FROM job WHERE JobID = ?)", jobID).Scan(&jobExists)
	//! Check if job exists
	if err != nil {
		SendResponse(w, models.Response{Message: "Error checking job existence: " + err.Error(), Status: "Error"})
		return
	}
	if !jobExists {
		SendResponse(w, models.Response{Message: "Job does not exist", Status: "Error"})
		return
	}
	var jobDetails struct {
		JobSpecs   models.Job `json:"job_specs"`
		Applicants []struct {
			Uuid            string         `json:"uuid"`
			Name            string         `json:"name"`
			Email           string         `json:"email"`
			ProfileHeadline string         `json:"profile_headline"`
			Address         string         `json:"address"`
			ResumeDetails   models.Profile `json:"profile"`
		} `json:"applicants"`
	}
	//! Response
	response = models.Response{
		Message: "Job ID: " + jobID,
		Status:  "Success",
		Data:    jobDetails,
	}
	SendResponse(w, response)
}
