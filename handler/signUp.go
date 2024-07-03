package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	apigateway "github.com/ArjunMalhotra07/Recruiter_Management_System/api_gateway"
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
	newUUID, err1 := exec.Command("uuidgen").Output()
	if err1 != nil {
		response := models.Response{Message: err1.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}
	encText, err := apigateway.Encrypt(user.PasswordHash, apigateway.MySecret)
	if err != nil {
		fmt.Println(err)
	}

	_, err = d.Driver.Exec(`INSERT INTO 
	user(uuid, Name,Email,Address,UserType,PasswordHash,ProfileHeadline) 
	VALUES (?,?,?,?,?,?,?)`,
		newUUID,
		user.Name,
		user.Email,
		user.Address,
		user.IsAdmin,
		encText,
		user.ProfileHeadline)

	if err != nil {
		response := models.Response{Message: err.Error(), Status: "Error"}
		SendResponse(w, response)
		return
	}

	tokenString, tokenError := apigateway.CreateToken(string(newUUID), user.IsAdmin)
	if tokenError != nil {
		fmt.Println("error", tokenError)
	}

	fmt.Println("Token sent", tokenString)
	response := models.Response{Message: "Created new user", Status: "Success", Jwt: tokenString}
	SendResponse(w, response)
}
