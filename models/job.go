package models

import (
	"time"
)

type Job struct {
	Title             string
	Description       string
	PostedOn          time.Time
	TotalApplications int
	CompanyName       string
	PostedBy          User
}
