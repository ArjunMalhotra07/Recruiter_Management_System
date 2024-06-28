package models

type User struct {
	Name            string
	Email           string
	Addr            string
	Type            UserType
	PasswordHash    string
	ProfileHeadline string
	Profile         *Profile
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
