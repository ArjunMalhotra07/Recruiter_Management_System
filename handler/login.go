package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	apigateway "github.com/ArjunMalhotra07/Recruiter_Management_System/api_gateway"
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
	encText, _ := apigateway.Encrypt(user.PasswordHash, apigateway.Secret)
	fmt.Println(encText)
	query := `SELECT * FROM user WHERE Email=? AND PasswordHash=?`

	row := d.Driver.QueryRow(query, user.Email, encText)
	var currentUser models.User
	err = row.Scan(&currentUser.Uuid, &currentUser.Name, &currentUser.Email, &currentUser.PasswordHash, &currentUser.IsAdmin, &currentUser.ProfileHeadline, &currentUser.Address)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}

	response := models.Response{Message: "User Exists", Status: "Success"}
	SendResponse(w, response)
}
