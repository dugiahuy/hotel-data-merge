package updater

import (
	"regexp"
	"strings"

	"github.com/dugiahuy/hotel-data-merge/src/model"
	"github.com/mitchellh/mapstructure"
)

const (
	onlyLetterAndNumber = `[^a-zA-Z0-9]+`
	paragraph           = `[=]+`
	upperCase           = `[A-Z][^A-Z]*`
)

func sanitize(input, pattern string, title bool) string {
	input = strings.Trim(input, " ")
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return input
	}
	input = strings.ReplaceAll(reg.ReplaceAllString(input, ""), "  ", " ")
	if title {
		input = strings.Title(strings.ToLower(input))
	}
	return input
}

var (
	dictGeneral = []string{
		"outdoorpool",
		"indoorpool",
		"businesscenter",
		"childcare",
		"parking",
		"bar",
		"drycleaning",
		"wifi",
		"breakfast",
		"concierge",
	}
	dictRoom = []string{
		"tv",
		"coffee machine",
		"kettle",
		"hairdryer",
		"iron",
		"bathtub",
		"minibar",
		"aircon",
	}
)

func convert(inputs []string) ([]string, []string) {
	compare := make(map[string]struct{})
	for _, v := range inputs {
		key := strings.ToLower(strings.ReplaceAll(v, " ", ""))
		compare[key] = struct{}{}
	}

	generals := []string{}
	for _, v := range dictGeneral {
		if _, ok := compare[v]; ok {
			generals = append(generals, v)
		}
	}

	rooms := []string{}
	for _, v := range dictRoom {
		if _, ok := compare[v]; ok {
			rooms = append(rooms, v)
		}
	}

	return generals, rooms
}

func convertToImageDetail(input interface{}) []model.ImageDetail {
	res := []model.ImageDetail{}

	switch v := input.(type) {
	case imageRoom1:
		for _, item := range v {
			res = append(res, model.ImageDetail{
				Link:        item.Link,
				Description: item.Description,
			})
		}
	case imageRoom2:
		for _, item := range v {
			res = append(res, model.ImageDetail{
				Link:        item.Link,
				Description: item.Description,
			})
		}
	}

	return res
}

func mergeMaps(maps ...map[string]interface{}) model.Hotel {
	merged := make(map[string]interface{})
	for _, m := range maps {
		for key, val := range m {
			switch key {
			case "id":
				if v, ok := val.(string); ok {
					if merged[key] == nil {
						merged[key] = v
					}
				}
			case "destination_id":
				if v, ok := val.(int64); ok {
					if merged[key] == nil {
						merged[key] = v
					}
				}
			case "lat", "lng":
				if v, ok := val.(float64); ok {
					if merged[key] == nil || v != 0 {
						merged[key] = v
					}
				}
			case "name", "address", "country", "city":
				merged[key] = mergeMapsName(merged[key], val.(string))
				// if v, ok := val.(string); ok {
				// 	v = sanitize(v, paragraph, true)
				// 	if merged[key] == nil || len(v) > len(merged[key].(string)) {
				// 		merged[key] = v
				// 	}
				// }
			case "description":
				if v, ok := val.(string); ok {
					v = sanitize(v, paragraph, false)
					if merged[key] == nil || len(v) > len(merged[key].(string)) {
						merged[key] = v
					}
				}
			case "general", "room":
				if v, ok := val.([]string); ok {
					if merged[key] == nil {
						merged[key] = v
					} else {
						mergedVal := merged[key].([]string)
						added := map[string]struct{}{}
						for _, item := range mergedVal {
							added[item] = struct{}{}
						}
						for _, item := range v {
							if _, ok := added[item]; !ok {
								added[item] = struct{}{}
								mergedVal = append(mergedVal, item)
							}
						}
						merged[key] = mergedVal
					}
				}
			case "rooms", "sites", "amenities":
				if v, ok := val.([]model.ImageDetail); ok {
					if merged[key] == nil {
						merged[key] = v
					} else {
						mergedVal := merged[key].([]model.ImageDetail)
						added := map[string]struct{}{}
						for _, item := range mergedVal {
							added[item.Link] = struct{}{}
						}
						for _, item := range v {
							if _, ok := added[item.Link]; !ok {
								added[item.Link] = struct{}{}
								mergedVal = append(mergedVal, item)
							}
						}
						merged[key] = mergedVal
					}
				}
			case "booking_conditions":
				if v, ok := val.([]string); ok {
					conds := []string{}
					for _, cond := range v {
						conds = append(conds, sanitize(cond, paragraph, false))
					}
					if merged[key] == nil {
						merged[key] = conds
					}
				}
			}
		}
	}

	var hotel model.Hotel
	mapstructure.Decode(merged, &hotel)
	return hotel
}

func mergeMapsName(mapVal interface{}, input string) string {
	input = sanitize(input, paragraph, true)
	if mapVal == nil || len(input) > len(mapVal.(string)) {
		return input
	}

	return mapVal.(string)
}
