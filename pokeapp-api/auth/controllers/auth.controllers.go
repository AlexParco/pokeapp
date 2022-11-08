package controllers

import (
	"fmt"
	"net/http"

	"github.com/alexparco/pokeapp-api/auth/services"
	"github.com/alexparco/pokeapp-api/model"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type authController struct {
	authService services.AuthService
	jwtService  services.JwtService
}

func NewAuthController(authService services.AuthService, jwtService services.JwtService) AuthController {
	return &authController{authService, jwtService}
}

// @Sumary Register user
// @Description register new user and return user with token
// @Router /user/register [POST]
func (a *authController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(400)
			return
		}

		createdUser, err := a.authService.Register(&user)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(404)
			return
		}
		fmt.Println(user.Username)

		token, err := a.jwtService.GeneratedToken(createdUser)
		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":  createdUser,
			"token": token,
		})
	}
}

// UserLogin struct not correct implementation
// @Sumary Login user
// @Description login useer and return user with token
// @Router /user/login [POST]
func (a *authController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userLogin model.UserLogin

		if err := c.BindJSON(&userLogin); err != nil {
			c.AbortWithStatus(400)
			return
		}

		loginUser, err := a.authService.Login(&model.User{
			Username: userLogin.Username,
			Password: userLogin.Password,
		})
		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		token, err := a.jwtService.GeneratedToken(loginUser)
		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"user":  loginUser,
			"token": token,
		})
	}
}
