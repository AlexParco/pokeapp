package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexparco/pokeapp-api/comment/services"
	"github.com/alexparco/pokeapp-api/model"
	"github.com/gin-gonic/gin"
)

type CommentController interface {
	Create() gin.HandlerFunc
	UpdateMessage() gin.HandlerFunc
	Delete() gin.HandlerFunc
	GetCommentsByPokeId() gin.HandlerFunc
}

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) CommentController {
	return &commentController{service}
}

// @Sumary Create new comment
// @Description register new comment and return comment
// @Router /comment [POST]
func (cmt *commentController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment model.Comment

		if err := c.BindJSON(&comment); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(404)
			return
		}

		// user id from gin context
		id := c.GetString("user_id")
		parseId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}

		comment.UserId = uint(parseId)

		createComment, err := cmt.service.Create(&comment)
		if err != nil {
			fmt.Println(createComment)
			c.AbortWithStatus(400)
			return
		}
		fmt.Println(createComment, id)

		c.JSON(http.StatusOK, createComment)
	}
}

// @Sumary Update comment
// @Description update comment body
// @Router /comment/{id}/ [PATCH]
func (cmt *commentController) UpdateMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment model.Comment
		if err := c.BindJSON(&comment); err != nil {
			c.AbortWithStatus(404)
			return
		}

		commentId, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		comment.CommentId = uint(commentId)

		userId, err := strconv.ParseUint(c.GetString("user_id"), 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		comment.UserId = uint(userId)
		updateComment, err := cmt.service.UpdateMessage(&comment)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(404)
			return
		}

		c.JSON(http.StatusOK, updateComment)
	}
}

// @Sumary Delete comment
// @Description delete comment
// @Router /comment/{id}/ [DELETE]
func (cmt *commentController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment model.Comment

		commentId, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		comment.CommentId = uint(commentId)

		userId, err := strconv.ParseUint(c.GetString("user_id"), 10, 32)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		comment.UserId = uint(userId)

		err = cmt.service.Delete(&comment)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(404)
			return
		}

		c.Status(http.StatusOK)
	}
}

// @Sumary Get comments
// @Description Get all comments by poke id
// @Router /comment?pokeid={id} [GET]
func (cmt *commentController) GetCommentsByPokeId() gin.HandlerFunc {
	return func(c *gin.Context) {
		pokeId, err := strconv.ParseUint(c.Query("pokeid"), 10, 32)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		comments, err := cmt.service.GetCommentsByPokeId(uint(pokeId))
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		c.JSON(http.StatusOK, comments)
	}
}
