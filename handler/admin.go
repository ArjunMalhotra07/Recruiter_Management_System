package handler

// UserType enum
type UserType int

const (
	Applicant UserType = iota
	Admin
)

// String method to provide a string representation for the UserType enum
func (ut UserType) String() string {
	return [...]string{"Applicant", "Admin"}[ut]
}

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
