package services

import (
	"errors"

	"github.com/alexparco/pokeapp-api/model"
	"github.com/alexparco/pokeapp-api/user/repository"
)

type UserService interface {
	Register(user *model.User) (*model.User, error)
	Login(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(userId uint) error
	GetById(userId uint) (*model.User, error)
	GetUsers() ([]*model.User, error)
}

type userService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userService{repo}
}

func (u *userService) Register(user *model.User) (*model.User, error) {
	existsUser, err := u.repo.FindByUsername(user)
	if err == nil || existsUser != nil {
		return nil, errors.New("error email alredy exists")
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	createdUser, err := u.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (u *userService) Login(user *model.User) (*model.User, error) {
	foundUser, err := u.repo.FindByUsername(user)
	if err != nil {
		return nil, err
	}

	if err := foundUser.ComparePassword(user.Password); err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (u *userService) Update(user *model.User) (*model.User, error) {
	updateUser, err := u.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (u *userService) Delete(userId uint) error {
	err := u.repo.Delete(userId)
	if err != nil {
		return nil
	}

	return nil
}

func (u *userService) GetById(userId uint) (*model.User, error) {
	foundUser, err := u.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (u *userService) GetUsers() ([]*model.User, error) {
	users, err := u.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
