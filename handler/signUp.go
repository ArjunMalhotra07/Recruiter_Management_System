package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	fmt.Println("Creating new user")
	response := models.Response{Message: "Created new user"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
