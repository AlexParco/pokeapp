package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexparco/pokeapp-api/comment/services"
	"github.com/alexparco/pokeapp-api/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController interface {
	Create() gin.HandlerFunc
	UpdateMessage() gin.HandlerFunc
	Delete() gin.HandlerFunc
	GetCommentsByPokeId() gin.HandlerFunc
	GetCommentById() gin.HandlerFunc
}

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) CommentController {
	return &commentController{service}
}

// @Sumary Create new message
// @Description register new comment and return comment
// @Router /comment/ [POST]
func (cmt *commentController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment model.Comment
		if err := c.BindJSON(&comment); err != nil {
			c.AbortWithStatus(404)
			return
		}

		id := c.GetString("user_id")

		parseId, err := uuid.Parse(id)

		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		comment.UserId = parseId

		createComment, err := cmt.service.Create(&comment)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		c.JSON(http.StatusOK, createComment)
	}
}

// @Sumary Update comment
// @Description update comment message
// @Router /comment/{id}/ [PATCH]
func (cmt *commentController) UpdateMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		var comment model.Comment
		if err := c.BindJSON(&comment); err != nil {
			c.AbortWithStatus(404)
			return
		}

		comment.CommentId = commentId
		comment.UserId = uuid.MustParse(c.GetString("user_id"))
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
// @Description delete message by id
// @Router /comment/{id}/ [DELETE]
func (cmt *commentController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment model.Comment

		commentId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		comment.CommentId = commentId
		comment.UserId = uuid.MustParse(c.GetString("user_id"))
		fmt.Println(comment)
		deleteComment, err := cmt.service.Delete(&comment)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(404)
			return
		}

		c.JSON(http.StatusOK, deleteComment)
	}
}

// @Sumary Get comment
// @Description Get comment by id
// @Router /comment/{id} [GET]
func (cmt *commentController) GetCommentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		comment, err := cmt.service.GetCommentById(commentId)
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		c.JSON(http.StatusOK, comment)
	}
}

// @Sumary Get comments
// @Description Get all comments by poke id
// @Router /comment?pokeid={id} [GET]
func (cmt *commentController) GetCommentsByPokeId() gin.HandlerFunc {
	return func(c *gin.Context) {
		pokeId, err := strconv.Atoi(c.Query("pokeid"))
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		fmt.Println(pokeId)
		comments, err := cmt.service.GetCommentsByPokeId(uint(pokeId))
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		c.JSON(http.StatusOK, comments)
	}
}
