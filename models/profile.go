package models

type Profile struct {
	Applicant         *User
	ResumeFileAddress string
	Skills            string
	Education         string
	Experience        string
	Name              string
	Email             string
	Phone             string
}
