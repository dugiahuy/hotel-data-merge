package hotel

import (
	"time"

	"github.com/dugiahuy/hotel-data-merge/src/model"
	"github.com/dugiahuy/hotel-data-merge/src/repository"
)

type Usecase interface {
	Get(id string) (*model.Hotel, error)
	GetByDestination(id int64) ([]model.Hotel, error)
}

func New(r repository.HotelStorage, timeout time.Duration) Usecase {
	return &usecase{
		repo:    r,
		timeout: timeout,
	}
}
