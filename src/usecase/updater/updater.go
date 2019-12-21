package updater

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/dugiahuy/hotel-data-merge/src/model"
)

type usecase struct {
	// repo   repository.HotelRepo
	logger *zap.Logger
}

func (u *usecase) CollectData() ([]model.Hotel, error) {
	var resp remoteOneResp
	if err := makeRequest(remote[0], &resp); err != nil {
		u.logger.Error(fmt.Sprintf("CollectData[%s]", remote[0]), zap.Error(err))
		return nil, err
	}

	return buildHotelFromRemoteOne(resp), nil
}
