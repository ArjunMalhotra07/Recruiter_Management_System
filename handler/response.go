package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
)

func SendResponse(w http.ResponseWriter, response models.Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
