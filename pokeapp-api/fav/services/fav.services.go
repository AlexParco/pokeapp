package services

import (
	"github.com/alexparco/pokeapp-api/fav/repository"
	"github.com/alexparco/pokeapp-api/model"
)

type FavService interface {
	SaveFav(fav *model.Fav) (*model.Fav, error)
	DeleteFav(fav *model.Fav) error
	FindAllFavs(userId uint) ([]*model.Fav, error)
}

type favService struct {
	repo repository.FavRepo
}

func NewFavService(repo repository.FavRepo) FavService {
	return &favService{repo}
}

func (f *favService) SaveFav(fav *model.Fav) (*model.Fav, error) {
	fav, err := f.repo.Create(fav)
	if err != nil {
		return nil, err
	}
	return fav, nil
}

func (f *favService) DeleteFav(fav *model.Fav) error {
	err := f.repo.Delete(fav)
	if err != nil {
		return err
	}
	return nil
}

func (f *favService) FindAllFavs(userId uint) ([]*model.Fav, error) {
	favs, err := f.repo.FindAll(userId)
	if err != nil {
		return nil, err
	}
	return favs, nil
}
