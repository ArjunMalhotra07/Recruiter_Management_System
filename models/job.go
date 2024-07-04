package models

type Job struct {
	Uuid              string `json:"uuid"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	PostedOn          string `json:"posted_on"`
	TotalApplications int    `json:"total_applications"`
	CompanyName       string `json:"company_name"`
	PostedBy          string `json:"posted_by"`
	Status            bool   `json:"active_status"`
}
