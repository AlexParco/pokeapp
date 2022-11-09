package repository

import (
	"github.com/alexparco/pokeapp-api/database"
	"github.com/alexparco/pokeapp-api/model"
)

type CommentRepo interface {
	Create(comment *model.Comment) (*model.Comment, error)
	UpdateMessage(comment *model.Comment) (*model.Comment, error)
	Delete(comment *model.Comment) error
	GetCommentsById(pokeId uint) ([]*model.Comment, error)
	ExistsByCommentId(commentId uint) bool
}

type commentRepo struct {
	*database.SqlClient
}

func NewCommentRepo(db *database.SqlClient) CommentRepo {
	return &commentRepo{db}
}

func (c *commentRepo) Create(comment *model.Comment) (*model.Comment, error) {
	var cmt model.Comment
	stmt, err := c.Prepare("INSERT INTO comments (user_id, pokemon_id, body) VALUES ($1, $2, $3) RETURNING comment_id, body, user_id, pokemon_id")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(comment.UserId, comment.PokemonId, comment.Body)
	err = row.Scan(&cmt.CommentId, &cmt.Body, &cmt.UserId, &cmt.PokemonId)
	if err != nil {
		return nil, err
	}

	return &cmt, nil

}

func (c *commentRepo) UpdateMessage(comment *model.Comment) (*model.Comment, error) {
	stmt, err := c.Prepare("UPDATE comments SET body=$1 WHERE comment_id=$2 and user_id=$3 RETURNING *")
	if err != nil {
		return nil, err
	}
	var cmt model.Comment
	row := stmt.QueryRow(comment.Body, comment.CommentId, comment.UserId)
	err = row.Scan(&cmt.CommentId, &cmt.Body, &cmt.UserId, &cmt.PokemonId)
	if err != nil {
		return nil, err
	}
	return &cmt, nil
}

func (c *commentRepo) Delete(comment *model.Comment) error {
	stmt, err := c.Prepare("DELETE FROM comments WHERE comment_id=$1 and user_id=$2")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(comment.CommentId, comment.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (c *commentRepo) GetCommentsById(pokeId uint) ([]*model.Comment, error) {
	rows, err := c.Query("SELECT comment_id, user_id, pokemon_id, body FROM comments WHERE pokemon_id=$1 ORDER BY created_at DESC;", pokeId)
	if err != nil {
		return nil, err
	}
	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.CommentId, &comment.UserId, &comment.PokemonId, &comment.Body)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (c *commentRepo) ExistsByCommentId(commentId uint) bool {
	row := c.QueryRow("SELECT * FROM comments WHERE comment_id=$1", commentId)

	var comment model.Comment
	err := row.Scan(&comment.CommentId, &comment.UserId, &comment.PokemonId, &comment.Body)

	return err != nil
}
