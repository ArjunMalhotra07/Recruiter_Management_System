package handler

import (
	"encoding/json"
	"fmt"
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
	if !claims["is_admin"].(bool) {
		SendResponse(w, models.Response{Message: "Only Admins can post jobs", Status: "Error"})
	} else {
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
	//! Getting claims through JWT Token
	claims := r.Context().Value("claims").(jwt.MapClaims)
	isAdmin := claims["is_admin"].(bool)
	if !isAdmin {
		SendResponse(w, models.Response{
			Message: "Only Admins can get list of all applicants", Status: "Error", Data: ""})
		return
	} else {
		rows, err := d.Driver.Query("SELECT Uuid, Name, Email, Address, IsAdmin, ProfileHeadline FROM user")
		if err != nil {
			SendResponse(w, models.Response{Message: err.Error(), Status: "Error"})
			return
		}
		defer rows.Close()
		var applicants []models.User
		for rows.Next() {
			var applicant models.User
			if err := rows.Scan(&applicant.Uuid, &applicant.Name, &applicant.Email, &applicant.Address, &applicant.IsAdmin, &applicant.ProfileHeadline); err != nil {
				SendResponse(w, models.Response{Message: err.Error(), Status: "Error"})
				return
			}
			applicants = append(applicants, applicant)
		}
		response := models.Response{Message: "All users data", Status: "Success", Data: applicants}
		SendResponse(w, response)
	}
}

func (d *Env) GetApplicantData(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	if !claims["is_admin"].(bool) {
		SendResponse(w, models.Response{Message: "Admin Users are allowed past here", Status: "Error!", Data: ""})
		return
	}
	applicantID := chi.URLParam(r, "applicant_id")
	fmt.Println(applicantID)
	rows, err := d.Driver.Query("SELECT Uuid, Name, Email, Address, IsAdmin, ProfileHeadline FROM user WHERE Uuid = ?", applicantID)
	if err != nil {
		SendResponse(w, models.Response{Message: err.Error(), Status: "Error!", Data: ""})
		return
	}
	defer rows.Close()
	var user models.User
	if rows.Next() {
		if err := rows.Scan(&user.Uuid, &user.Name, &user.Email, &user.Address, &user.IsAdmin, &user.ProfileHeadline); err != nil {
			SendResponse(w, models.Response{Message: err.Error(), Status: "Error"})
			return
		}
	} else {
		SendResponse(w, models.Response{Message: "No user found with the given ID", Status: "Error", Data: ""})
		return
	}
	response := models.Response{Message: "Data of a single user", Status: "Success", Data: user}
	SendResponse(w, response)
}
