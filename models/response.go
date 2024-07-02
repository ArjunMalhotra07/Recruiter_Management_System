package models

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Jwt     string `json:"jwt"`
}
