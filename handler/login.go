package handler

import (
	"encoding/json"
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
	encText, _ := apigateway.Encrypt(user.PasswordHash, apigateway.MySecret)
	query := `SELECT Uuid FROM user WHERE Email=? AND PasswordHash=?`

	row, err := d.Driver.Query(query, user.Email, encText)
	if err != nil {
		SendResponse(w, models.Response{Message: err.Error(), Status: "Error!", Data: ""})
		return
	}
	if row.Next() {
		var currentUser models.User
		if err = row.Scan(&currentUser.Uuid); err != nil {
			response := models.Response{Message: err.Error(), Status: "Error"}
			SendResponse(w, response)
			return
		}
		tokenString, tokenError := apigateway.CreateToken(string(currentUser.Uuid), user.IsAdmin)
		if tokenError != nil {
			SendResponse(w, models.Response{Message: "Error generating JWT", Status: "Error"})
			return
		}
		response := models.Response{Message: "User exists", Status: "Success", Jwt: tokenString}
		SendResponse(w, response)
		return
	} else {
		response := models.Response{Message: "Email ID or Password doesn't match", Status: "Error"}
		SendResponse(w, response)
		return
	}
}
