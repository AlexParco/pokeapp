package model

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentId uuid.UUID `json:"comment_id"`
	UserId    uuid.UUID `json:"user_id"`
	PokemonId uint      `json:"pokemon_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
