package handler

import (
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
)

func (d *Env) GetJobDetails(w http.ResponseWriter, r *http.Request) {
	// Get Claim Values and job_id
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
	// Check if job exists
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

	//! Fetch job details
	err = d.Driver.QueryRow("SELECT * FROM job WHERE JobID = ?", jobID).Scan(
		&jobDetails.JobSpecs.Uuid,
		&jobDetails.JobSpecs.Title,
		&jobDetails.JobSpecs.Description,
		&jobDetails.JobSpecs.PostedOn,
		&jobDetails.JobSpecs.TotalApplications,
		&jobDetails.JobSpecs.CompanyName,
		&jobDetails.JobSpecs.PostedBy,
		&jobDetails.JobSpecs.Status,
	)
	if err != nil {
		SendResponse(w, models.Response{Message: "Error fetching job details: " + err.Error(), Status: "Error"})
		return
	}

	//! Fetch applicants' user IDs
	rows, err := d.Driver.Query("SELECT UserID FROM job_application WHERE JobID = ?", jobID)
	if err != nil {
		SendResponse(w, models.Response{Message: "Error fetching applicants: " + err.Error(), Status: "Error"})
		return
	}
	defer rows.Close()

	var userIDs []string
	for rows.Next() {
		var userID string
		if err := rows.Scan(&userID); err != nil {
			SendResponse(w, models.Response{Message: "Error scanning user ID: " + err.Error(), Status: "Error"})
			return
		}
		userIDs = append(userIDs, userID)
	}

	//! Fetch user details and profile for each applicant
	for _, userID := range userIDs {
		var applicant struct {
			Uuid            string         `json:"uuid"`
			Name            string         `json:"name"`
			Email           string         `json:"email"`
			ProfileHeadline string         `json:"profile_headline"`
			Address         string         `json:"address"`
			ResumeDetails   models.Profile `json:"profile"`
		}

		//! Fetch user details
		err = d.Driver.QueryRow("SELECT Uuid, Name, Email, ProfileHeadline, Address FROM user WHERE Uuid = ?", userID).Scan(
			&applicant.Uuid,
			&applicant.Name,
			&applicant.Email,
			&applicant.ProfileHeadline,
			&applicant.Address,
		)
		if err != nil {
			SendResponse(w, models.Response{Message: "Error fetching user details: " + err.Error(), Status: "Error"})
			return
		}

		//! Fetch profile details
		err = d.Driver.QueryRow("SELECT * FROM profile WHERE Uuid = ?", userID).Scan(
			&applicant.ResumeDetails.Uuid,
			&applicant.ResumeDetails.ResumeFileAddress,
			&applicant.ResumeDetails.Skills,
			&applicant.ResumeDetails.Education,
			&applicant.ResumeDetails.Experience,
			&applicant.ResumeDetails.Name,
			&applicant.ResumeDetails.Email,
			&applicant.ResumeDetails.Phone,
		)
		if err != nil {
			SendResponse(w, models.Response{Message: "Error fetching profile details: " + err.Error(), Status: "Error"})
			return
		}

		jobDetails.Applicants = append(jobDetails.Applicants, applicant)
	}

	//! Response
	response = models.Response{
		Message: "Job ID: " + jobID,
		Status:  "Success",
		Data:    jobDetails,
	}
	SendResponse(w, response)
}
