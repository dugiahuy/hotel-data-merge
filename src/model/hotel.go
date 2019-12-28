package model

type Hotel struct {
	ID            string `json:"id"`
	DestinationID int64  `json:"destination_id" mapstructure:"destination_id"`
	Name          string `json:"name"`
	Location      struct {
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
		Address string  `json:"address"`
		City    string  `json:"city"`
		Country string  `json:"country"`
	} `json:"location" mapstructure:",squash"`
	Description string `json:"description"`
	Amenities   struct {
		General []string `json:"general"`
		Room    []string `json:"room"`
	} `json:"amenities" mapstructure:",squash"`
	Images struct {
		Rooms     []ImageDetail `json:"rooms"`
		Sites     []ImageDetail `json:"sites"`
		Amenities []ImageDetail `json:"amenities"`
	} `json:"images" mapstructure:",squash"`
	BookingConditions []string `json:"booking_conditions" mapstructure:"booking_conditions"`
}

type ImageDetail struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}

type Hotels []Hotel

func (hotels Hotels) GetByID(id string) *Hotel {
	for i := range hotels {
		if hotels[i].ID == id {
			return &hotels[i]
		}
	}
	return nil
}

func (hotels Hotels) GetByDestinationID(dest int64) []Hotel {
	var res []Hotel
	for i := range hotels {
		if hotels[i].DestinationID == dest {
			res = append(res, hotels[i])
		}
	}
	return res
}
