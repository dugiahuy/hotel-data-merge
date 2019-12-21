package model

type RemoteOne struct {
	ID            string   `json:"Id"`
	DestinationID int      `json:"DestinationId"`
	Name          string   `json:"Name"`
	Latitude      float64  `json:"Latitude"`
	Longitude     float64  `json:"Longitude"`
	Address       string   `json:"Address"`
	City          string   `json:"City"`
	Country       string   `json:"Country"`
	PostalCode    string   `json:"PostalCode"`
	Description   string   `json:"Description"`
	Facilities    []string `json:"Facilities"`
}

type RemoteOnes []RemoteOne

func (hotels RemoteOnes) GroupByID() map[string]*RemoteOne {
	res := make(map[string]*RemoteOne)
	for i := range hotels {
		res[hotels[i].ID] = &hotels[i]
	}
	return res
}
