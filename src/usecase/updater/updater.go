package updater

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/dugiahuy/hotel-data-merge/src/model"
	"github.com/dugiahuy/hotel-data-merge/src/repository"
	"github.com/k0kubun/pp"
)

type usecase struct {
	repo   repository.HotelStorage
	logger *zap.Logger
}

func (u *usecase) CollectData() error {
	dataRemote0, err := fetchData(linkRemote[0], parser0)
	if err != nil {
		u.logger.Error(fmt.Sprintf("CollectData[%s]", linkRemote[0]), zap.Error(err))
		return err
	}
	dataRemote1, err := fetchData(linkRemote[1], parser1)
	if err != nil {
		u.logger.Error(fmt.Sprintf("CollectData[%s]", linkRemote[1]), zap.Error(err))
		return err
	}
	dataRemote2, err := fetchData(linkRemote[2], parser2)
	if err != nil {
		u.logger.Error(fmt.Sprintf("CollectData[%s]", linkRemote[2]), zap.Error(err))
		return err
	}

	var hotels []model.Hotel
	parsedMap := make(map[string]struct{})
	for id := range dataRemote0 {
		if _, ok := parsedMap[id]; !ok {
			parsedMap[id] = struct{}{}
			hotels = append(hotels, mergeMaps(dataRemote0[id], dataRemote1[id], dataRemote2[id]))
			pp.Println(dataRemote0[id])
			pp.Println(dataRemote1[id])
			pp.Println(dataRemote2[id])
			pp.Println(hotels)
		}
	}
	for id := range dataRemote1 {
		if _, ok := parsedMap[id]; !ok {
			parsedMap[id] = struct{}{}
			hotels = append(hotels, mergeMaps(dataRemote0[id], dataRemote1[id], dataRemote2[id]))
		}
	}
	for id := range dataRemote2 {
		if _, ok := parsedMap[id]; !ok {
			parsedMap[id] = struct{}{}
			hotels = append(hotels, mergeMaps(dataRemote0[id], dataRemote1[id], dataRemote2[id]))
		}
	}

	if err := u.repo.Store(hotels); err != nil {
		u.logger.Error("CollectData.Store", zap.Error(err))
		return err
	}

	return nil
}
