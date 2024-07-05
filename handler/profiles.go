package handler

import (
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
)

func (d *Env) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	rows, err := d.Driver.Query("SELECT * FROM profile")
	if err != nil {
		SendResponse(w, models.Response{Message: err.Error(), Status: "Error"})
		return
	}
	defer rows.Close()
	var profiles []models.Profile
	for rows.Next() {
		var profile models.Profile
		if err := rows.Scan(&profile.Uuid, &profile.ResumeFileAddress, &profile.Skills, &profile.Education, &profile.Experience, &profile.Name, &profile.Email, &profile.Phone); err != nil {
			SendResponse(w, models.Response{Message: err.Error(), Status: "Error"})
			return
		}
		profiles = append(profiles, profile)
	}
	if len(profiles) != 0 {
		response := models.Response{Message: "All Profiles data", Status: "Success", Data: profiles}
		SendResponse(w, response)
	} else {
		response := models.Response{Message: "No New Profiles", Status: "Success", Data: []models.Job{}}
		SendResponse(w, response)
	}
}
