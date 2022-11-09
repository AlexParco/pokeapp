package model

type Comment struct {
	CommentId uint   `json:"comment_id"`
	Body      string `json:"body"`
	UserId    uint   `json:"user_id"`
	PokemonId uint   `json:"pokemon_id"`
}

type CommentPayload struct {
	CommentId uint   `json:"comment_id"`
	Body      string `json:"body"`
	Username  string `json:"username"`
	PokemonId uint   `json:"pokemon_id"`
}
