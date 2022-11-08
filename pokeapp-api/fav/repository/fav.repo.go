package repository

import (
	"github.com/alexparco/pokeapp-api/database"
	"github.com/alexparco/pokeapp-api/model"
)

type FavRepo interface {
	Create(fav *model.Fav) (*model.Fav, error)
	Delete(fav *model.Fav) error
	FindAll(userId uint) ([]*model.Fav, error)
}

type favRepo struct {
	*database.SqlClient
}

func NewFavRepo(db *database.SqlClient) FavRepo {
	return &favRepo{db}
}

func (f *favRepo) Create(fav *model.Fav) (*model.Fav, error) {
	stmt, err := f.Prepare("INSERT INTO pokefavs (user_id, pokemon_id) VALUES ($1, $2) RETURNING *")
	if err != nil {
		return nil, err
	}

	var pfav model.Fav
	row := stmt.QueryRow(fav.UserId, fav.PokemonId)
	err = row.Scan(&pfav.PokefavId, &pfav.UserId, &pfav.PokemonId)
	if err != nil {
		return nil, err
	}

	return &pfav, nil
}

func (f *favRepo) Delete(fav *model.Fav) error {
	stmt, err := f.Prepare("DELETE FROM pokefavs WHERE pokemon_id=$1 and user_id=$2")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(fav.PokemonId, fav.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (f *favRepo) FindAll(userId uint) ([]*model.Fav, error) {
	rows, err := f.Query("SELECT * FROM pokefavs WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}

	var favs []*model.Fav
	for rows.Next() {
		var fav model.Fav
		err = rows.Scan(&fav.PokefavId, &fav.UserId, &fav.PokemonId)
		if err != nil {
			return nil, err
		}
		favs = append(favs, &fav)
	}
	return favs, nil
}
