package handler

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"time"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
)

func (d *Env) PostJob(w http.ResponseWriter, r *http.Request) {
	//! Getting claims through JWT Token
	claims := r.Context().Value("claims").(jwt.MapClaims)
	postedBy := claims["uuid"].(string)
	//! Parsing data
	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	//! Generating new id for job
	newUUID, err1 := exec.Command("uuidgen").Output()
	if err1 != nil {
		response := models.Response{Message: err1.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	//! Insert into Jobs Table
	_, err = d.Driver.Exec(`INSERT INTO 
	Job(JobID, Title,Description,PostedOn,TotalApplications,CompanyName,PostedBy) 
	VALUES (?,?,?,?,?,?,?)`,
		newUUID,
		job.Title,
		job.Description,
		time.Now(),
		job.TotalApplications,
		job.CompanyName,
		postedBy)

	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	//! Send Response
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
