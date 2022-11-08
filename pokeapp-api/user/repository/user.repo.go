package repository

import (
	"github.com/alexparco/pokeapp-api/database"
	"github.com/alexparco/pokeapp-api/model"
)

type UserRepo interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(userId uint) error
	GetUserById(userId uint) (*model.User, error)
	GetUsers() ([]*model.User, error)
	FindByUsername(user *model.User) (*model.User, error)
}

type userRepo struct {
	db *database.SqlClient
}

func NewUserRepo(db *database.SqlClient) UserRepo {
	return &userRepo{db}
}

func (u *userRepo) Create(user *model.User) (*model.User, error) {
	stmt, err := u.db.Prepare("INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING * ")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(user.Username, user.Password, user.Role)
	uUser := model.User{}
	err = row.Scan(&uUser.UserId, &uUser.Username, &uUser.Password, &uUser.Role)

	if err != nil {
		return nil, err
	}

	return &uUser, nil
}

func (u *userRepo) Update(user *model.User) (*model.User, error) {
	stmt, err := u.db.Prepare("UPDATE users SET username=$1 WHERE user_id=$2 RETURNING *")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(user.Username, user.UserId)

	var uUser model.User
	err = row.Scan(&uUser.UserId, &uUser.Username, &uUser.Password, &uUser.Role)
	if err != nil {
		return nil, err
	}

	return &uUser, nil
}

func (u *userRepo) Delete(userId uint) error {
	_, err := u.db.Exec("DELETE FROM users WHERE user_id=$1", userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) GetUserById(userId uint) (*model.User, error) {
	var user model.User
	row := u.db.QueryRow("SELECT * FROM users WHERE user_id=$1", userId)
	err := row.Scan(&user.UserId, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetUsers() ([]*model.User, error) {
	var users []*model.User

	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserId, &user.Username, &user.Password, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *userRepo) FindByUsername(user *model.User) (*model.User, error) {
	row := u.db.QueryRow("SELECT * FROM users WHERE username=$1", user.Username)
	foundUser := model.User{}
	err := row.Scan(&foundUser.UserId, &foundUser.Username, &foundUser.Password, &foundUser.Role)
	if err != nil {
		return nil, err
	}
	return &foundUser, nil
}
