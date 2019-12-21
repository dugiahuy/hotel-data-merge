package model

type RemoteThree struct {
	ID            string   `json:"id"`
	DestinationID int      `json:"destination"`
	Name          string   `json:"name"`
	Latitude      float64  `json:"lat"`
	Longitude     float64  `json:"lng"`
	Address       string   `json:"address"`
	Description   string   `json:"info"`
	Amenities     []string `json:"amenities"`
	Images        struct {
		Rooms []struct {
			URL         string `json:"url"`
			Description string `json:"description"`
		} `json:"rooms"`
		Amenities []struct {
			URL         string `json:"url"`
			Description string `json:"description"`
		} `json:"amenities"`
	} `json:"images"`
}

type RemoteThrees []RemoteThree

func (hotels RemoteThrees) GroupByID() map[string]*RemoteThree {
	res := make(map[string]*RemoteThree)
	for i := range hotels {
		res[hotels[i].ID] = &hotels[i]
	}
	return res
}
