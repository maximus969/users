package httpserver

import (
	"github.com/maximus969/users-app/internal/app/domain"
)

func toResponseUser(user domain.User) UserResponse {
	return UserResponse{
		Id:        user.Id(),
		Firstname: user.Firstname(),
		Lastname:  user.Lastname(),
		Age:       user.Age(),
		Email:     user.Email(),
	}
}

func toDomainUser(userRequest UserRequest) (domain.User, error) {
	return domain.NewUser(domain.NewUserData{
		Firstname: userRequest.Firstname,
		Lastname:  userRequest.Lastname,
		Email:     userRequest.Email,
		Age:       userRequest.Age,
	})
}
