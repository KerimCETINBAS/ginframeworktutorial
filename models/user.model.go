package models

type UserContact struct {
	Phone string `json:"telefon"`
}

type User struct {
	UserContact
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginDto struct {
	Email    string
	Password string
}

type UserLoginResponse struct {
	Email string
}
