package models

type User struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	PasswordHash    string   `json:"password_hash"`
	Type            string `json:"type"`
	ProfileHeadline string   `json:"profile_headline"`
	Address         string   `json:"address"`
}

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