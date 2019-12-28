package updater

import (
	"go.uber.org/zap"

	"github.com/dugiahuy/hotel-data-merge/src/repository"
)

type Usecase interface {
	CollectData() error
}

func New(r repository.HotelStorage, logger *zap.Logger) Usecase {
	return &usecase{
		repo:   r,
		logger: logger,
	}
}
