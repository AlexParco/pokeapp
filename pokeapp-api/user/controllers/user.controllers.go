package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexparco/pokeapp-api/model"
	"github.com/alexparco/pokeapp-api/user/services"
	"github.com/gin-gonic/gin"
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

		parseId, err := strconv.ParseUint(c.GetString("user_id"), 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		user, err := u.service.GetById(uint(parseId))
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

		userId, err := strconv.ParseUint(c.GetString("user_id"), 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		var userUpt model.UserUpdate
		if err := c.Bind(&userUpt); err != nil {
			c.AbortWithStatus(404)
			return
		}

		updateUser, err := u.service.Update(&model.User{
			UserId:   uint(userId),
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

		parseId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}

		if err := u.service.Delete(uint(parseId)); err != nil {
			c.AbortWithStatus(404)
			return
		}

		c.Status(http.StatusOK)
	}
}
