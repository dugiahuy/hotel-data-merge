package repository

import (
	"encoding/json"
	"io/ioutil"

	"github.com/dugiahuy/hotel-data-merge/src/model"
)

type repo struct{}

func (r *repo) Store(hotels []model.Hotel) error {
	hotelJSON, err := json.Marshal(hotels)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("data.json", hotelJSON, 0644); err != nil {
		return err
	}
	return nil
}

func (r *repo) Fetch() (model.Hotels, error) {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var hotels []model.Hotel
	if err := json.Unmarshal(data, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}

func (r *repo) Get(id string) (*model.Hotel, error) {
	hotels, err := r.Fetch()
	if err != nil {
		return nil, err
	}
	return hotels.GetByID(id), nil
}

func (r *repo) GetByDestination(id int64) ([]model.Hotel, error) {
	hotels, err := r.Fetch()
	if err != nil {
		return nil, err
	}
	return hotels.GetByDestinationID(id), nil
}
