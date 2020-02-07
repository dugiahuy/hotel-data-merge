package updater

import (
	"reflect"
	"testing"

	"github.com/dugiahuy/hotel-data-merge/src/model"
)

func Test_sanitize(t *testing.T) {
	type args struct {
		input   string
		pattern string
		title   bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sanitize a title",
			args: args{
				input:   "test  ascenda",
				pattern: paragraph,
				title:   true,
			},
			want: "Test Ascenda",
		},
		{
			name: "have error when comple regexp",
			args: args{
				input:   "cant compile",
				pattern: "[[A-Z][^A-Z]*",
				title:   false,
			},
			want: "cant compile",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitize(tt.args.input, tt.args.pattern, tt.args.title); got != tt.want {
				t.Errorf("sanitize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeMaps(t *testing.T) {
	type args struct {
		maps []map[string]interface{}
	}
	mapData0 := map[string]interface{}{
		"name":        "Beach Villas Singapore",
		"lng":         103.824006,
		"address":     " 8 Sentosa Gateway, Beach Villas ",
		"description": "  This 5 star hotel is located on the coastline of Singapore.",
		"id":          "iJhz",
		"lat":         1.264751,
		"city":        "Singapore",
		"country":     "SG",
		"general": []string{
			"businesscenter",
			"drycleaning",
			"wifi",
			"breakfast",
		},
		"room":           []string{},
		"destination_id": 5432,
	}

	mapData1 := map[string]interface{}{
		"room": []string{
			"tv",
			"kettle",
			"hairdryer",
			"iron",
		},
		"rooms": []model.ImageDetail{
			model.ImageDetail{
				Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg",
				Description: "Double room",
			},
			model.ImageDetail{
				Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/3.jpg",
				Description: "Double room",
			},
		},
		"booking_conditions": []string{
			"All children are welcome. One child under 12 years stays free of charge when using existing beds. One child under 2 years stays free of charge in a child's cot/crib. One child under 4 years stays free of charge when using existing beds. One older child or adult is charged SGD 82.39 per person per night in an extra bed. The maximum number of children's cots/cribs in a room is 1. There is no capacity for extra beds in the room.",
			"Pets are not allowed.",
			"WiFi is available in all areas and is free of charge.",
			"Free private parking is possible on site (reservation is not needed).",
			"Guests are required to show a photo identification and credit card upon check-in. Please note that all Special Requests are subject to availability and additional charges may apply. Payment before arrival via bank transfer is required. The property will contact you after you book to provide instructions. Please note that the full amount of the reservation is due before arrival. Resorts World Sentosa will send a confirmation with detailed payment information. After full payment is taken, the property's details, including the address and where to collect keys, will be emailed to you. Bag checks will be conducted prior to entry to Adventure Cove Waterpark. === Upon check-in, guests will be provided with complimentary Sentosa Pass (monorail) to enjoy unlimited transportation between Sentosa Island and Harbour Front (VivoCity). === Prepayment for non refundable bookings will be charged by RWS Call Centre. === All guests can enjoy complimentary parking during their stay, limited to one exit from the hotel per day. === Room reservation charges will be charged upon check-in. Credit card provided upon reservation is for guarantee purpose. === For reservations made with inclusive breakfast, please note that breakfast is applicable only for number of adults paid in the room rate. Any children or additional adults are charged separately for breakfast and are to paid directly to the hotel.",
		},
		"id":          "iJhz",
		"address":     "8 Sentosa Gateway, Beach Villas, 098269",
		"description": "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex and a 2-minute walk from the Waterfront train station. Featuring sundecks and pool, garden or sea views, the plush 1- to 3-bedroom villas offer free Wi-Fi and flat-screens, as well as free-standing baths, minibars, and tea and coffeemaking facilities. Upgraded villas add private pools, fridges and microwaves; some have wine cellars. A 4-bedroom unit offers a kitchen and a living room. There's 24-hour room and butler service. Amenities include posh restaurant, plus an outdoor pool, a hot tub, and free parking.",
		"general": []string{
			"outdoorpool",
			"indoorpool",
			"businesscenter",
			"childcare",
		},
		"destination_id": 5432,
		"name":           "Beach Villas Singapore",
		"country":        "Singapore",
		"sites": []model.ImageDetail{
			model.ImageDetail{
				Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/1.jpg",
				Description: "Front",
			},
		},
	}
	mapData2 := map[string]interface{}{
		"amenities": []model.ImageDetail{
			model.ImageDetail{
				Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg",
				Description: "RWS",
			},
			model.ImageDetail{
				Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/6.jpg",
				Description: "Sentosa Gateway",
			},
		},
		"destination_id": 5432,
		"lat":            1.264751,
		"address":        "8 Sentosa Gateway, Beach Villas, 098269",
		"description":    "Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters. Guests will find themselves in paradise with this series of exquisite tropical sanctuaries, making it the perfect setting for an idyllic retreat. Within each villa, guests will discover living areas and bedrooms that open out to mini gardens, private timber sundecks and verandahs elegantly framingeither lush greenery or an expanse of sea. Guests are assured of a superior slumber with goose feather pillows and luxe mattresses paired with 400 thread count Egyptian cotton bed linen, tastefully paired with a full complement of luxurious in-room amenities and bathrooms boasting rain showers and free-standing tubs coupled with an exclusive array of ESPA amenities and toiletries. Guests also get to enjoy complimentary day access to the facilities at Asia’s flagship spa – the world-renowned ESPA.",
		"room":           []string{},
		"id":             "iJhz",
		"name":           "Beach Villas Singapore",
		"lng":            103.824006,
		"general":        []string{},
		"rooms": []model.ImageDetail{
			model.ImageDetail{
				Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg",
				Description: "Double room",
			},
			model.ImageDetail{
				Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg",
				Description: "Bathroom",
			},
		},
	}

	// successModel := model.Hotel{
	// ID:            "iJhz",
	// DestinationID: 5432,
	// Name:          "Beach Villas Singapore",
	// Location: struct{
	// 	Lat:     1.264751,
	// 	Lng:     103.824006,
	// 	Address: "8 Sentosa Gateway, Beach Villas, 098269",
	// 	City:    "Singapore",
	// 	Country: "Singapore",
	// },
	// Description: "Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters. Guests will find themselves in paradise with this series of exquisite tropical sanctuaries, making it the perfect setting for an idyllic retreat. Within each villa, guests will discover living areas and bedrooms that open out to mini gardens, private timber sundecks and verandahs elegantly framing either lush greenery or an expanse of sea. Guests are assured of a superior slumber with goose feather pillows and luxe mattresses paired with 400 thread count Egyptian cotton bed linen, tastefully paired with a full complement of luxurious in-room amenities and bathrooms boasting rain showers and free-standing tubs coupled with an exclusive array of ESPA amenities and toiletries. Guests also get to enjoy complimentary day access to the facilities at Asia’s flagship spa – the world-renowned ESPA.",
	// Amenities: {
	// 	General: []string{
	// 		"businesscenter",
	// 		"drycleaning",
	// 		"wifi",
	// 		"breakfast",
	// 		"outdoorpool",
	// 		"indoorpool",
	// 		"childcare",
	// 	},
	// 	Room: []string{
	// 		"tv",
	// 		"kettle",
	// 		"hairdryer",
	// 		"iron",
	// 	},
	// },
	// Images: {
	// 	Rooms: []model.ImageDetail{
	// 		model.ImageDetail{
	// 			Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg",
	// 			Description: "Double room",
	// 		},
	// 		model.ImageDetail{
	// 			Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/3.jpg",
	// 			Description: "Double room",
	// 		},
	// 		model.ImageDetail{
	// 			Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg",
	// 			Description: "Bathroom",
	// 		},
	// 	},
	// 	Sites: []model.ImageDetail{
	// 		model.ImageDetail{
	// 			Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/1.jpg",
	// 			Description: "Front",
	// 		},
	// 	},
	// 	Amenities: []model.ImageDetail{
	// 		model.ImageDetail{
	// 			Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg",
	// 			Description: "RWS",
	// 		},
	// 		model.ImageDetail{
	// 			Link:        "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/6.jpg",
	// 			Description: "Sentosa Gateway",
	// 		},
	// 	},
	// },
	// BookingConditions: []string{
	// 	"All children are welcome. One child under 12 years stays free of charge when using existing beds. One child under 2 years stays free of charge in a child's cot/crib. One child under 4 years stays free of charge when using existing beds. One older child or adult is charged SGD 82.39 per person per night in an extra bed. The maximum number of children's cots/cribs in a room is 1. There is no capacity for extra beds in the room.",
	// 	"Pets are not allowed.",
	// 	"WiFi is available in all areas and is free of charge.",
	// 	"Free private parking is possible on site (reservation is not needed).",
	// 	"Guests are required to show a photo identification and credit card upon check-in. Please note that all Special Requests are subject to availability and additional charges may apply. Payment before arrival via bank transfer is required. The property will contact you after you book to provide instructions. Please note that the full amount of the reservation is due before arrival. Resorts World Sentosa will send a confirmation with detailed payment information. After full payment is taken, the property's details, including the address and where to collect keys, will be emailed to you. Bag checks will be conducted prior to entry to Adventure Cove Waterpark. Upon check-in, guests will be provided with complimentary Sentosa Pass (monorail) to enjoy unlimited transportation between Sentosa Island and Harbour Front (VivoCity). Prepayment for non refundable bookings will be charged by RWS Call Centre. All guests can enjoy complimentary parking during their stay, limited to one exit from the hotel per day. Room reservation charges will be charged upon check-in. Credit card provided upon reservation is for guarantee purpose. For reservations made with inclusive breakfast, please note that breakfast is applicable only for number of adults paid in the room rate. Any children or additional adults are charged separately for breakfast and are to paid directly to the hotel.",
	// },
	// 	}
	// 	successModel.Location = struct{{
	// 		Lat:     1.264751,
	// 		Lng:     103.824006,
	// 		Address: "8 Sentosa Gateway, Beach Villas, 098269",
	// 		City:    "Singapore",
	// 		Country: "Singapore",
	// 	},
	// }
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success case",
			args: args{maps: []map[string]interface{}{mapData0, mapData1, mapData2}},
			want: "8 Sentosa Gateway, Beach Villas, 098269",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeMaps(tt.args.maps...); !reflect.DeepEqual(got.Location.Address, tt.want) {
				t.Errorf("mergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_mergeMapsLocation(t *testing.T) {
// 	type args struct {
// 		maps []map[string]interface{}
// 	}
// 	mapData0 := map[string]interface{}{
// 		"name":        "Beach Villas Singapore",
// 		"lng":         103.824006,
// 		"address":     " 8 Sentosa Gateway, Beach Villas ",
// 		"description": "  This 5 star hotel is located on the coastline of Singapore.",
// 		"id":          "iJhz",
// 		"lat":         1.264751,
// 		"city":        "Singapore",
// 		"country":     "SG",
// 	}
// 	mapData1 := map[string]interface{}{
// 		"address":        "8 Sentosa Gateway, Beach Villas, 098269",
// 		"description":    "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex and a 2-minute walk from the Waterfront train station. Featuring sundecks and pool, garden or sea views, the plush 1- to 3-bedroom villas offer free Wi-Fi and flat-screens, as well as free-standing baths, minibars, and tea and coffeemaking facilities. Upgraded villas add private pools, fridges and microwaves; some have wine cellars. A 4-bedroom unit offers a kitchen and a living room. There's 24-hour room and butler service. Amenities include posh restaurant, plus an outdoor pool, a hot tub, and free parking.",
// 		"destination_id": 5432,
// 		"name":           "Beach Villas Singapore",
// 		"country":        "Singapore",
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want model.Hotel
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := mergeMapsLocation(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("mergeMapsLocation() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_mergeMapsName(t *testing.T) {
	type args struct {
		mapKey interface{}
		input  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				mapKey: "8 Sentosa Gateway, Beach Villas, 098269",
				input:  "8 Sentosa Gateway",
			},
			want: "8 Sentosa Gateway, Beach Villas, 098269",
		},
		{
			name: "case 2",
			args: args{
				mapKey: "8 Sentosa Gateway, Beach Villas",
				input:  "8 Sentosa Gateway, Beach Villas, 09826999999",
			},
			want: "8 Sentosa Gateway, Beach Villas, 09826999999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeMapsName(tt.args.mapKey, tt.args.input); got != tt.want {
				t.Errorf("mergeMapsName() = %v, want %v", got, tt.want)
			}
		})
	}
}
