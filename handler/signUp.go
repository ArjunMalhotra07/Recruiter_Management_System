package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
)

type Env struct {
	Driver *sql.DB
}

func (d *Env) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	_, err = d.Driver.Exec(`INSERT INTO 
	user(Name,Email,Address,UserType,PasswordHash,ProfileHeadline) 
	VALUES (?,?,?,?,?,?)`,
		user.Name,
		user.Email,
		user.Address,
		user.Type,
		user.PasswordHash,
		user.ProfileHeadline)

	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	response := models.Response{Message: "Created new user", Status: "Success"}
	SendResponse(w, response)
}
