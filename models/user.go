package models

type User struct {
	Uuid            string `json:"uuid"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	PasswordHash    string `json:"password_hash"`
	IsAdmin         bool   `json:"is_admin"`
	ProfileHeadline string `json:"profile_headline"`
	Address         string `json:"address"`
	Jwt             string `json:"jwt"`
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
