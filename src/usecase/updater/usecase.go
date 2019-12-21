package updater

import (
	"go.uber.org/zap"

	"github.com/dugiahuy/hotel-data-merge/src/model"
)

type Usecase interface {
	CollectData() ([]model.Hotel, error)
}

// func New(r repository.HotelRepo, logger *zap.Logger) Usecase {
// 	return &usecase{
// 		repo:   r,
// 		logger: logger,
// 	}
// }

func New(logger *zap.Logger) Usecase {
	return &usecase{
		logger: logger,
	}
}
