package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
)

func (d *Env) LogIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	query := `SELECT * FROM user WHERE Email=? AND PasswordHash=?`

	rows, err1 := d.Driver.Query(query, user.Email, user.PasswordHash)
	var currentUser models.User
	if rows.Next() {
		for rows.Next() {
			if err := rows.Scan(&currentUser.Name, &currentUser.Email, &currentUser.PasswordHash, &currentUser.Type, &currentUser.ProfileHeadline, &currentUser.Address); err != nil {
				return
			}
		}
		fmt.Println(currentUser)
	} else {
		response := models.Response{Message: err1.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}

	if err1 != nil {
		response := models.Response{Message: err1.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	response := models.Response{Message: "User Exists", Status: "Success"}
	SendResponse(w, response)
}
