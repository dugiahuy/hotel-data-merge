package repository

import (
	"github.com/dugiahuy/hotel-data-merge/src/model"
)

type HotelStorage interface {
	Fetch() (model.Hotels, error)
	Get(id string) (*model.Hotel, error)
	GetByDestination(id int64) ([]model.Hotel, error)
	Store([]model.Hotel) error
}

func NewStorage() HotelStorage {
	return &repo{}
}
