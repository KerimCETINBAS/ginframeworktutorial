package user

import (
	"github.com/kerimcetinbas/gogintut/models"
)

type userService struct {
	users []models.User
}

type IUserService interface {
	GetUsers() []models.User
	CreateUser(data models.User)
}

func UserService() IUserService {
	return &userService{
		users: []models.User{},
	}
}

func (s *userService) GetUsers() []models.User {
	return s.users
}

func (s *userService) CreateUser(data models.User) {
	data.Id = uint(len(s.users))

	data.Phone = "1234523425"
	s.users = append(s.users, data)
}
