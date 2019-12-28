package hotel

import (
	"time"

	"github.com/dugiahuy/hotel-data-merge/src/model"
	"github.com/dugiahuy/hotel-data-merge/src/repository"
)

type usecase struct {
	repo    repository.HotelStorage
	timeout time.Duration
}

func (u *usecase) Get(id string) (*model.Hotel, error) {
	return u.repo.Get(id)
}

func (u *usecase) GetByDestination(id int64) ([]model.Hotel, error) {
	return u.repo.GetByDestination(id)
}
