package updater

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dugiahuy/hotel-data-merge/src/model"
)

const (
	onlyLetterAndNumber = "[^a-zA-Z0-9]+"
	paragraph           = "[^a-zA-Z0-9,. ]+"
	upperCase           = "[A-Z][^A-Z]*"
)

var remote = []string{
	"https://api.myjson.com/bins/gdmqa",
	"https://api.myjson.com/bins/1fva3m",
	"https://api.myjson.com/bins/j6kzm",
}

type remoteOneResp struct {
	RemoteOnes []*model.RemoteOne
}
type remoteTwoResp struct {
	RemoteTwos []*model.RemoteTwo
}
type remoteThreeResp struct {
	RemoteThrees []*model.RemoteThree
}

var client = http.Client{
	Timeout: time.Duration(time.Second * 20),
}

func buildHotelFromRemoteOne(src remoteOneResp) []model.Hotel {
	res := []model.Hotel{}
	for _, v := range src.RemoteOnes {
		hotel := model.Hotel{}
		hotel.ID = v.ID
		hotel.DestinationID = v.DestinationID
		hotel.Name = Santinize(v.Name, onlyLetterAndNumber)
		hotel.Location.Lat = v.Latitude
		hotel.Location.Lng = v.Longitude
		hotel.Location.Address = Santinize(v.Address, paragraph)
		hotel.Location.City = Santinize(v.City, onlyLetterAndNumber)
		hotel.Location.Country = Santinize(v.City, onlyLetterAndNumber)
		hotel.Description = Santinize(v.Description, paragraph)
		res = append(res, hotel)
	}

	return res
}

// Make a Regex to say we only want letters and numbers
func Santinize(input, pattern string) string {
	input = strings.ReplaceAll(input, "  ", " ")
	input = strings.Trim(input, " ")
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return input
	}
	return reg.ReplaceAllString(input, "")
}

func Convert(inputs []string) ([]string, []string) {
	re := regexp.MustCompile(upperCase)
	general := []string{}
	room := []string{}
	for _, v := range inputs {
		general = append(general,
			strings.ToLower(
				strings.Join(
					re.FindAllString(v, -1),
					" ",
				),
			),
		)
	}

	return general, room
}

func makeRequest(url string, into interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode == 404 {
		return fmt.Errorf("API not found")
	}

	if err := json.NewDecoder(r.Body).Decode(&into); err != nil {
		return err
	}

	return nil
}
