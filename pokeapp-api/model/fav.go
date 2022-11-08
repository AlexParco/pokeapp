package model

type Fav struct {
	PokefavId uint `json:"pokefav_id"`
	UserId    uint `json:"user_id,omitempty"`
	PokemonId uint `json:"pokemon_id,omitempty"`
}
