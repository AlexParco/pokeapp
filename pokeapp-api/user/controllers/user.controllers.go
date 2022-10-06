package controllers

import (
	"fmt"
	"net/http"

	"github.com/alexparco/pokeapp-api/model"
	"github.com/alexparco/pokeapp-api/user/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	Profile() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{service}
}

// @Sumary Profile of user
// @Description get profile of user
// @Params id of user
// @Router /user/{id} [GET]
func (u *userController) Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(400)
			return
		}
		user, err := u.service.GetById(userId)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(404)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// @Sumary Profile of user
// @Description get profile of user
// @Router /user [GET]
func (u *userController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := u.service.GetUsers()
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// UserUpdate struct not correct implementation
// @Sumary Update user
// @Description update username existing user
// @Params id of user
// @Router /auth [PATCH]
func (u *userController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("user_id")

		userId, err := uuid.Parse(id)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		var userUpt model.UserUpdate
		if err := c.Bind(&userUpt); err != nil {
			c.AbortWithStatus(404)
			return
		}

		updateUser, err := u.service.Update(&model.User{
			UserId:   userId,
			Username: userUpt.Username,
		})

		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		c.JSON(http.StatusOK, updateUser)

	}
}

// @Sumary Delete user
// @Description delete user by id
// @Params id of user
// @Router /auth [DELETE]
func (u *userController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("user_id")

		userId, err := uuid.Parse(id)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		if err := u.service.Delete(userId); err != nil {
			c.AbortWithStatus(404)
			return
		}

		c.Status(http.StatusOK)
	}
}
