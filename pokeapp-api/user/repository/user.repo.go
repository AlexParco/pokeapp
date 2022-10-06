package repository

import (
	"github.com/alexparco/pokeapp-api/database"
	"github.com/alexparco/pokeapp-api/model"
	"github.com/google/uuid"
)

type UserRepo interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(userId uuid.UUID) error
	GetUserById(userId uuid.UUID) (*model.User, error)
	GetUsers() ([]*model.User, error)
	FindByEmail(user *model.User) (*model.User, error)
}

type userRepo struct {
	db *database.SqlClient
}

func NewUserRepo(db *database.SqlClient) UserRepo {
	return &userRepo{db}
}

func (u *userRepo) Create(user *model.User) (*model.User, error) {
	stmt, err := u.db.Prepare("INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, 'user') RETURNING * ")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(user.Username, user.Email, user.Password)
	uUser := model.User{}
	err = row.Scan(&uUser.UserId, &uUser.Username, &uUser.Email, &uUser.Password, &uUser.Role, &uUser.CreatedAt, &uUser.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &uUser, nil
}

func (u *userRepo) Update(user *model.User) (*model.User, error) {
	stmt, err := u.db.Prepare("UPDATE users SET username=$1, updated_at = CURRENT_TIMESTAMP WHERE user_id=$2 RETURNING *")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(user.Username, user.UserId)

	var uUser model.User
	err = row.Scan(&uUser.UserId, &uUser.Username, &uUser.Email, &uUser.Password, &uUser.Role, &uUser.CreatedAt, &uUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &uUser, nil
}

func (u *userRepo) Delete(userId uuid.UUID) error {
	_, err := u.db.Exec("DELETE FROM users WHERE user_id=$1", userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) GetUserById(userId uuid.UUID) (*model.User, error) {
	var user model.User
	row := u.db.QueryRow("SELECT * FROM users WHERE user_id=$1", userId)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
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
		if err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *userRepo) FindByEmail(user *model.User) (*model.User, error) {
	row := u.db.QueryRow("SELECT * FROM users WHERE email=$1", user.Email)
	foundUser := model.User{}
	err := row.Scan(&foundUser.UserId, &foundUser.Username, &foundUser.Email, &foundUser.Password, &foundUser.Role, &foundUser.CreatedAt, &foundUser.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &foundUser, nil
}
