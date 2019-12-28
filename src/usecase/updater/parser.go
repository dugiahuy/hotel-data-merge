package updater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/buger/jsonparser"
)

var linkRemote = []string{
	"https://api.myjson.com/bins/gdmqa",
	"https://api.myjson.com/bins/1fva3m",
	"https://api.myjson.com/bins/j6kzm",
}

func fetchData(url string, parser func([]byte) map[string]map[string]interface{}) (map[string]map[string]interface{}, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode == 404 {
		return nil, fmt.Errorf("API not found")
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return parser(data), nil
}

func parser0(data []byte) map[string]map[string]interface{} {
	idx := 0
	mapDatas := make(map[string]map[string]interface{})
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		item := "[" + strconv.Itoa(idx) + "]"
		idx++

		id, _ := jsonparser.GetString(data, item, "Id")
		dest, _ := jsonparser.GetInt(data, item, "DestinationId")
		name, _ := jsonparser.GetString(data, item, "Name")
		lat, _ := jsonparser.GetFloat(data, item, "Latitude")
		lng, _ := jsonparser.GetFloat(data, item, "Longitude")
		addr, _ := jsonparser.GetString(data, item, "Address")
		city, _ := jsonparser.GetString(data, item, "City")
		country, _ := jsonparser.GetString(data, item, "Country")
		desc, _ := jsonparser.GetString(data, item, "Description")
		faciByte, _, _, _ := jsonparser.Get(data, item, "Facilities")
		var facilities []string
		json.Unmarshal(faciByte, &facilities)
		generals, rooms := convert(facilities)

		mapData := map[string]interface{}{
			"id":             id,
			"destination_id": dest,
			"name":           name,
			"lat":            lat,
			"lng":            lng,
			"address":        addr,
			"city":           city,
			"country":        country,
			"description":    desc,
			"general":        generals,
			"room":           rooms,
		}
		mapDatas[id] = mapData
	})

	return mapDatas
}

type imageRoom1 []struct {
	Link        string `json:"link"`
	Description string `json:"caption"`
}

func parser1(data []byte) map[string]map[string]interface{} {
	idx := 0
	mapDatas := make(map[string]map[string]interface{})
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		item := "[" + strconv.Itoa(idx) + "]"
		idx++

		id, _ := jsonparser.GetString(data, item, "hotel_id")
		dest, _ := jsonparser.GetInt(data, item, "destination_id")
		name, _ := jsonparser.GetString(data, item, "hotel_name")
		addr, _ := jsonparser.GetString(data, item, "location", "address")
		country, _ := jsonparser.GetString(data, item, "location", "country")
		desc, _ := jsonparser.GetString(data, item, "details")

		var imageRooms imageRoom1
		imageRoom, _, _, _ := jsonparser.Get(data, item, "images", "rooms")
		json.Unmarshal(imageRoom, &imageRooms)
		imgDetailRooms := convertToImageDetail(imageRooms)
		var imageSites imageRoom1
		imageSite, _, _, _ := jsonparser.Get(data, item, "images", "site")
		json.Unmarshal(imageSite, &imageSites)
		imgDetailSites := convertToImageDetail(imageSites)

		var bookConds []string
		bookCond, _, _, _ := jsonparser.Get(data, item, "booking_conditions")
		json.Unmarshal(bookCond, &bookConds)

		var generalStr []string
		general, _, _, _ := jsonparser.Get(data, item, "amenities", "general")
		json.Unmarshal(general, &generalStr)
		var roomStr []string
		room, _, _, _ := jsonparser.Get(data, item, "amenities", "room")
		json.Unmarshal(room, &roomStr)
		amentities := generalStr
		amentities = append(amentities, roomStr...)
		generals, rooms := convert(amentities)

		mapData := map[string]interface{}{
			"id":                 id,
			"destination_id":     dest,
			"name":               name,
			"address":            addr,
			"country":            country,
			"description":        desc,
			"general":            generals,
			"room":               rooms,
			"rooms":              imgDetailRooms,
			"sites":              imgDetailSites,
			"booking_conditions": bookConds,
		}
		mapDatas[id] = mapData
	})

	return mapDatas
}

type imageRoom2 []struct {
	Link        string `json:"url"`
	Description string `json:"description"`
}

func parser2(data []byte) map[string]map[string]interface{} {
	idx := 0
	mapDatas := make(map[string]map[string]interface{})
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		item := "[" + strconv.Itoa(idx) + "]"
		idx++

		id, _ := jsonparser.GetString(data, item, "id")
		dest, _ := jsonparser.GetInt(data, item, "destination")
		name, _ := jsonparser.GetString(data, item, "name")
		lat, _ := jsonparser.GetFloat(data, item, "lat")
		lng, _ := jsonparser.GetFloat(data, item, "lng")
		addr, _ := jsonparser.GetString(data, item, "address")
		desc, _ := jsonparser.GetString(data, item, "info")

		var amenities []string
		general, _, _, _ := jsonparser.Get(data, item, "amenities", "general")
		json.Unmarshal(general, &amenities)
		generals, rooms := convert(amenities)

		var imageRooms imageRoom2
		imageRoom, _, _, _ := jsonparser.Get(data, item, "images", "rooms")
		json.Unmarshal(imageRoom, &imageRooms)
		imgDetailRooms := convertToImageDetail(imageRooms)

		var imageAmenities imageRoom2
		imageAmennity, _, _, _ := jsonparser.Get(data, item, "images", "amenities")
		json.Unmarshal(imageAmennity, &imageAmenities)
		imgDetailAmenities := convertToImageDetail(imageAmenities)

		mapData := map[string]interface{}{
			"id":             id,
			"destination_id": dest,
			"name":           name,
			"lat":            lat,
			"lng":            lng,
			"address":        addr,
			"description":    desc,
			"general":        generals,
			"room":           rooms,
			"rooms":          imgDetailRooms,
			"amenities":      imgDetailAmenities,
		}
		mapDatas[id] = mapData
	})

	return mapDatas
}
