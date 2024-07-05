package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/dgrijalva/jwt-go"
)

func (d *Env) UploadResume(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	if claims["is_admin"].(bool) {
		SendResponse(w, models.Response{Message: "Only Applicant Type Users can upload resume!", Status: "Error"})
		return
	}
	file, header, err := r.FormFile("resume")
	if err != nil {
		SendResponse(w, models.Response{Message: err.Error(), Status: "Error"})
		return
	}
	defer file.Close()
	//! Save resume to a directory
	resumeDir := "./resumes"
	if err := os.MkdirAll(resumeDir, os.ModePerm); err != nil {
		SendResponse(w, models.Response{Message: "Error creating resume directory: " + err.Error(), Status: "Error"})
		return
	}
	resumeFilePath := filepath.Join(resumeDir, header.Filename)

	// Create the file on the server
	outFile, err := os.Create(resumeFilePath)
	if err != nil {
		SendResponse(w, models.Response{Message: "Error saving file: " + err.Error(), Status: "Error"})
		return
	}
	defer outFile.Close()

	// Copy the uploaded file to the created file on the server
	if _, err := io.Copy(outFile, file); err != nil {
		SendResponse(w, models.Response{Message: "Error saving file: " + err.Error(), Status: "Error"})
		return
	}

	// Read the saved file into a byte slice
	fileContent, err := os.ReadFile(resumeFilePath)
	if err != nil {
		SendResponse(w, models.Response{Message: "Error reading file: " + err.Error(), Status: "Error"})
		return
	}

	//! API call to resume parser
	apiKey := "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1"
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", bytes.NewReader(fileContent))
	if err != nil {
		SendResponse(w, models.Response{Message: "Error creating request: " + err.Error(), Status: "Error"})
		return
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		SendResponse(w, models.Response{Message: "Error uploading file: " + err.Error(), Status: "Error"})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		SendResponse(w, models.Response{Message: "Error from resume parsing API: " + string(bodyBytes), Status: "Error"})
		return
	}

	//! Parsing Resume
	var parsedResume models.Profile
	if err := json.NewDecoder(resp.Body).Decode(&parsedResume); err != nil {
		SendResponse(w, models.Response{Message: "Error decoding response: " + err.Error(), Status: "Error"})
		return
	}

	// Extracting and joining Education and Experience names
	var educationNames []string
	for _, edu := range parsedResume.Education {
		educationNames = append(educationNames, edu.Name)
	}
	education := strings.Join(educationNames, ", ")

	var experienceNames []string
	for _, exp := range parsedResume.Experience {
		experienceNames = append(experienceNames, exp.Name)
	}
	experience := strings.Join(experienceNames, ", ")

	//! Insert into profile table
	_, err = d.Driver.Exec("INSERT INTO profile (Uuid, Name, Email, Phone, ResumeFileAddress, Skills, Education, Experience) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		claims["uuid"], parsedResume.Name, parsedResume.Email, parsedResume.Phone, resumeFilePath, strings.Join(parsedResume.Skills, ", "), education, experience)
	if err != nil {
		SendResponse(w, models.Response{Message: "Error saving profile: " + err.Error(), Status: "Error"})
		return
	}
	SendResponse(w, models.Response{Message: "Resume Uploaded Successfully!", Status: "Success"})
}
