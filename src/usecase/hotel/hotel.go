package hotel

import (
	"time"

	"github.com/dugiahuy/hotel-data-merge/src/repository"
)

type usecase struct {
	repo    repository.HotelRepo
	timeout time.Duration
}

func (u *usecase) HealthCheck() (bool, error) {
	return u.repo.HealthCheck()
}
