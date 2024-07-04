package handler

import (
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
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

func (env *Env) ApplyToJob(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "Applied to a job", Status: "Success"}
	SendResponse(w, response)
}
