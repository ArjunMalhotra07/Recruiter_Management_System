package models

type Profile struct {
	Uuid              string `json:"uuid"`
	Skills            string `json:"skills"`
	ResumeFileAddress string `json:"resume_file_address"`
	Education         string `json:"education"`
	Experience        string `json:"experience"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
}

type Education struct {
	Name string `json:"name"`
}

type Experience struct {
	Title        string `json:"title"`
	Organization string `json:"organization"`
}
