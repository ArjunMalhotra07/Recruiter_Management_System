package handler

type User struct {
	Name            string
	Email           string
	Addr            string
	Type            UserType
	PasswordHash    string
	ProfileHeadline string
	Profile         *Profile
}
