package services

import (
	"errors"

	"github.com/alexparco/pokeapp-api/comment/repository"
	"github.com/alexparco/pokeapp-api/model"
	"github.com/google/uuid"
)

type CommentService interface {
	Create(comment *model.Comment) (*model.Comment, error)
	UpdateMessage(comment *model.Comment) (*model.Comment, error)
	Delete(comment *model.Comment) (*model.Comment, error)
	GetCommentsByPokeId(pokeId uint) ([]*model.Comment, error)
	GetCommentById(comentId uuid.UUID) (*model.Comment, error)
}
type commentService struct {
	repo repository.CommentRepo
}

func NewCommentService(repo repository.CommentRepo) CommentService {
	return &commentService{repo}
}

func (c *commentService) Create(comment *model.Comment) (*model.Comment, error) {
	createUser, err := c.repo.Create(comment)
	if err != nil {
		return nil, err
	}
	return createUser, nil
}

func (c *commentService) UpdateMessage(comment *model.Comment) (*model.Comment, error) {
	_, err := c.repo.GetCommentById(comment.CommentId)
	if err != nil {
		return nil, errors.New("error comment does not exist")
	}

	updateComment, err := c.repo.UpdateMessage(comment)
	if err != nil {
		return nil, err
	}
	return updateComment, nil
}

func (c *commentService) Delete(comment *model.Comment) (*model.Comment, error) {
	commnet, err := c.repo.Delete(comment)
	if err != nil {
		return nil, err
	}
	return commnet, nil
}

func (c *commentService) GetCommentsByPokeId(pokeId uint) ([]*model.Comment, error) {
	comments, err := c.repo.GetComments(pokeId)
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (c *commentService) GetCommentById(commentId uuid.UUID) (*model.Comment, error) {
	comment, err := c.repo.GetCommentById(commentId)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
