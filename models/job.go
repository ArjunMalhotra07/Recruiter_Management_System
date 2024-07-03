package models

import (
	"time"
)

type Job struct {
	Uuid              string    `json:"uuid"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	PostedOn          time.Time `json:"posted_on"`
	TotalApplications int       `json:"total_applications"`
	CompanyName       string    `json:"company_name"`
	PostedBy          User      `json:"posted_by"`
	Status            bool      `json:"active_status"`
}
