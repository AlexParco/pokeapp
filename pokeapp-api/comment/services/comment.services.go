package services

import (
	"errors"

	"github.com/alexparco/pokeapp-api/comment/repository"
	"github.com/alexparco/pokeapp-api/model"
)

type CommentService interface {
	Create(comment *model.Comment) (*model.Comment, error)
	UpdateMessage(comment *model.Comment) (*model.Comment, error)
	Delete(comment *model.Comment) error
	GetCommentsByPokeId(pokeId uint) ([]*model.Comment, error)
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
	if !c.repo.ExistsByCommentId(comment.CommentId) {
		return nil, errors.New("error comment does not exist")
	}

	updateComment, err := c.repo.UpdateMessage(comment)
	if err != nil {
		return nil, err
	}
	return updateComment, nil
}

func (c *commentService) Delete(comment *model.Comment) error {
	if !c.repo.ExistsByCommentId(comment.CommentId) {
		return errors.New("error comment does not exist")
	}

	err := c.repo.Delete(comment)
	if err != nil {
		return err
	}
	return nil
}

func (c *commentService) GetCommentsByPokeId(pokeId uint) ([]*model.Comment, error) {
	comments, err := c.repo.GetCommentsById(pokeId)
	if err != nil {
		return nil, err
	}
	return comments, err
}
