package services

import (
	"errors"

	"github.com/alexparco/pokeapp-api/model"
	"github.com/alexparco/pokeapp-api/user/repository"
)

type AuthService interface {
	Register(user *model.User) (*model.User, error)
	Login(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(userId uint) error
}

type authService struct {
	repo repository.UserRepo
}

func NewAuthService(repo repository.UserRepo) AuthService {
	return &authService{repo}
}

func (a *authService) Register(user *model.User) (*model.User, error) {
	foundUser, err := a.repo.FindByUsername(user)
	if err == nil || foundUser != nil {
		return nil, errors.New("error username alredy exists")
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	createdUser, err := a.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (a *authService) Login(user *model.User) (*model.User, error) {
	foundUser, err := a.repo.FindByUsername(user)
	if err != nil {
		return nil, err
	}

	if err := foundUser.ComparePassword(user.Password); err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (a *authService) Update(user *model.User) (*model.User, error) {
	updateUser, err := a.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (a *authService) Delete(userId uint) error {
	err := a.repo.Delete(userId)
	if err != nil {
		return nil
	}

	return nil
}
