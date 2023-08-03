package models

type User struct {
	ID       uint
	Phone    string
	Email    string
	Password string
}

type UserLoginDto struct {
	Email    string
	Password string
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func (u *UserResponse) FromUser(data User) {
	u.ID = data.ID
	u.Email = data.Email

}
