package hotel

import (
	"time"

	"github.com/dugiahuy/hotel-data-merge/src/repository"
)

type Usecase interface {
	HealthCheck() (bool, error)
}

func New(r repository.HotelRepo, timeout time.Duration) Usecase {
	return &usecase{
		repo:    r,
		timeout: timeout,
	}
}
