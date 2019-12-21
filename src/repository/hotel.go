package repository

import (
	"context"

	"github.com/dugiahuy/hotel-data-merge/src/model"
)

type HotelRepo interface {
	HealthCheck() (bool, error)
	Store(context.Context, *model.Hotel) error
}
