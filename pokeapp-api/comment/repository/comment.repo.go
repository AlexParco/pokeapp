package repository

import (
	"github.com/alexparco/pokeapp-api/database"
	"github.com/alexparco/pokeapp-api/model"
	"github.com/google/uuid"
)

type CommentRepo interface {
	Create(comment *model.Comment) (*model.Comment, error)
	UpdateMessage(comment *model.Comment) (*model.Comment, error)
	Delete(comment *model.Comment) (*model.Comment, error)
	GetCommentById(commentId uuid.UUID) (*model.Comment, error)
	GetComments(pokeId uint) ([]*model.Comment, error)
}

type commentRepo struct {
	*database.SqlClient
}

func NewCommentRepo(db *database.SqlClient) CommentRepo {
	return &commentRepo{db}
}

func (c *commentRepo) Create(comment *model.Comment) (*model.Comment, error) {
	stmt, err := c.Prepare("INSERT INTO comments (user_id, pokemon_id, message) VALUES ($1, $2, $3) RETURNING *")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(comment.UserId, comment.PokemonId, comment.Message)
	var cmt model.Comment
	err = row.Scan(&cmt.CommentId, &cmt.UserId, &cmt.PokemonId, &cmt.Message, &cmt.CreatedAt, &cmt.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &cmt, nil

}

func (c *commentRepo) UpdateMessage(comment *model.Comment) (*model.Comment, error) {
	stmt, err := c.Prepare("UPDATE comments SET message=$1 WHERE comment_id=$2 and user_id=$3 RETURNING *")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(comment.Message, comment.CommentId, comment.UserId)
	var cmt model.Comment
	err = row.Scan(&cmt.CommentId, &cmt.UserId, &cmt.PokemonId, &cmt.Message, &cmt.CreatedAt, &cmt.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &cmt, nil
}

func (c *commentRepo) Delete(comment *model.Comment) (*model.Comment, error) {
	stmt, err := c.Prepare("DELETE FROM comments WHERE comment_id=$1 and user_id=$2 RETURNING *")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(comment.CommentId, comment.UserId)
	var cmt model.Comment
	err = row.Scan(&cmt.CommentId, &cmt.UserId, &cmt.PokemonId, &cmt.Message, &cmt.CreatedAt, &cmt.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &cmt, nil
}

func (c *commentRepo) GetComments(pokeId uint) ([]*model.Comment, error) {
	rows, err := c.Query("SELECT * FROM comments WHERE pokemon_id=$1 ORDER BY created_at DESC;", pokeId)
	if err != nil {
		return nil, err
	}
	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.CommentId, &comment.UserId, &comment.PokemonId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (c *commentRepo) GetCommentById(commentId uuid.UUID) (*model.Comment, error) {
	row := c.QueryRow("SELECT * FROM comments WHERE comment_id=$1", commentId.String())

	var comment model.Comment
	err := row.Scan(&comment.CommentId, &comment.UserId, &comment.PokemonId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
