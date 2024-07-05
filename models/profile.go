package models

type Profile struct {
	Education  []Education  `json:"education"`
	Email      string       `json:"email"`
	Experience []Experience `json:"experience"`
	Name       string       `json:"name"`
	Phone      string       `json:"phone"`
	Skills     []string     `json:"skills"`
}

type Education struct {
	Name string `json:"name"`
}

type Experience struct {
	Name  string   `json:"name"`
	Dates []string `json:"dates"`
}
