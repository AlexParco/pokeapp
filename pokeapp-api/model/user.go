package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdate struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
