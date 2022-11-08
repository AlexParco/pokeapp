package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexparco/pokeapp-api/fav/services"
	"github.com/alexparco/pokeapp-api/model"
	"github.com/gin-gonic/gin"
)

type FavController interface {
	PostFav() gin.HandlerFunc
	DeleteFav() gin.HandlerFunc
	GetAllFavsByUserId() gin.HandlerFunc
}

type favController struct {
	service services.FavService
}

func NewFavController(service services.FavService) FavController {
	return &favController{service}
}

func (f *favController) PostFav() gin.HandlerFunc {
	return func(c *gin.Context) {
		var fav model.Fav
		if err := c.BindJSON(&fav); err != nil {
			c.AbortWithStatus(404)
			return
		}

		parseId, err := strconv.ParseUint(c.GetString("user_id"), 10, 32)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(500)
			return
		}
		fav.UserId = uint(parseId)

		createFav, err := f.service.SaveFav(&fav)

		if err != nil {
			c.AbortWithStatus(400)
			return
		}
		c.JSON(http.StatusOK, createFav)
	}
}

func (f *favController) DeleteFav() gin.HandlerFunc {
	return func(c *gin.Context) {
		var fav model.Fav
		pokemonId, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
		}
		fav.PokemonId = uint(pokemonId)

		userId, err := strconv.ParseUint(c.GetString("user_id"), 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		fav.UserId = uint(userId)

		fmt.Println(fav)

		err = f.service.DeleteFav(&fav)
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		c.Status(http.StatusOK)
	}
}

// @Router /fav [GET]
func (f *favController) GetAllFavsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.ParseUint(c.GetString("user_id"), 10, 32)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(400)
			return
		}

		favs, err := f.service.FindAllFavs(uint(userId))
		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		c.JSON(http.StatusOK, favs)
	}
}
