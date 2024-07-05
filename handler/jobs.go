package handler

import (
	"net/http"
	"os/exec"
	"time"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/dgrijalva/jwt-go"
)

func (d *Env) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	rows, err := d.Driver.Query("SELECT * FROM job WHERE Status = ?", false)
	if err != nil {
		SendResponse(w, models.Response{Message: err.Error(), Status: "Error"})
		return
	}
	defer rows.Close()
	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.Uuid, &job.Title, &job.Description, &job.PostedOn, &job.TotalApplications, &job.CompanyName, &job.PostedBy, &job.Status); err != nil {
			SendResponse(w, models.Response{Message: err.Error(), Status: "Error"})
			return
		}
		jobs = append(jobs, job)
	}
	if len(jobs) != 0 {
		response := models.Response{Message: "All users data", Status: "Success", Data: jobs}
		SendResponse(w, response)
	} else {
		response := models.Response{Message: "No New Jobs", Status: "Success", Data: []models.Job{}}
		SendResponse(w, response)
	}
}

func (d *Env) ApplyToJob(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	if claims["is_admin"].(bool) {
		SendResponse(w, models.Response{Message: "Only Applicant Users can apply to this job!", Status: "Error"})
		return
	}
	uuid := claims["uuid"].(string)
	jobID := r.URL.Query().Get("job_id") // Extract job_id from query parameters
	if jobID == "" {
		SendResponse(w, models.Response{Message: "job_id parameter is required", Status: "Error"})
		return
	}
	//! Check if the job with jobID exists
	var jobExists bool
	err := d.Driver.QueryRow("SELECT EXISTS(SELECT 1 FROM job WHERE JobID = ?)", jobID).Scan(&jobExists)
	if err != nil {
		SendResponse(w, models.Response{Message: "Error checking job existence: " + err.Error(), Status: "Error"})
		return
	}
	if !jobExists {
		SendResponse(w, models.Response{Message: "Job does not exist", Status: "Error"})
		return
	}
	// ! Generating new applicationID
	applicationID, err := exec.Command("uuidgen").Output()
	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	_, err = d.Driver.Exec(`INSERT INTO 
job_application(ApplicationID, JobID,UserID,AppliedOn) 
VALUES (?,?,?,?)`,
		applicationID,
		jobID,
		uuid,
		time.Now(),
	)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	response := models.Response{Message: "Successfully Applied!", Status: "Success"}
	SendResponse(w, response)

}
