package pgrepo

import (
	"github.com/maximus969/users-app/internal/app/domain"
	"github.com/maximus969/users-app/internal/app/repository/models"
)

func domainToUser(user domain.User) models.User {
	return models.User{
		Id:        user.Id(),
		Firstname: user.Firstname(),
		Lastname:  user.Lastname(),
		Email:     user.Email(),
		Age:       user.Age(),
	}
}

func userToDomain(user models.User) (domain.User, error) {
	return domain.NewUser(domain.NewUserData{
		Id:        user.Id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Age:       user.Age,
	})
}
