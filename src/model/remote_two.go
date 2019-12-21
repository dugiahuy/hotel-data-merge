package model

type RemoteTwo struct {
	ID            string `json:"hotel_id"`
	DestinationID int    `json:"destination_id"`
	Name          string `json:"hotel_name"`
	Location      struct {
		Address string `json:"address"`
		Country string `json:"country"`
	} `json:"location"`
	Description string `json:"details"`
	Amenities   struct {
		General []string `json:"general"`
		Room    []string `json:"room"`
	} `json:"amenities"`
	Images struct {
		Rooms []struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		} `json:"rooms"`
		Site []struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		} `json:"site"`
	} `json:"images"`
	BookingConditions []string `json:"booking_conditions"`
}

type RemoteTwos []RemoteTwo

func (hotels RemoteTwos) GroupByID() map[string]*RemoteTwo {
	res := make(map[string]*RemoteTwo)
	for i := range hotels {
		res[hotels[i].ID] = &hotels[i]
	}
	return res
}
