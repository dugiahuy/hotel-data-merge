package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/dugiahuy/hotel-data-merge/src/model"
	"github.com/k0kubun/pp"
)

var remote = []string{
	"https://api.myjson.com/bins/gdmqa",
	"https://api.myjson.com/bins/1fva3m",
	"https://api.myjson.com/bins/j6kzm",
}

var client = http.Client{
	Timeout: time.Duration(time.Second * 20),
}

func main() {
	r, err := client.Get(remote[0])
	if err != nil {
		log.Fatal("can't get ")
	}
	defer r.Body.Close()

	if r.StatusCode == 404 {
		log.Fatal("404")
		// return fmt.Errorf("API not found")
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("data")
	}

	paths := [][]string{
		[]string{"Id"},
		[]string{"DestinationId"},
		[]string{"Name"},
		[]string{"Latitude"},
		[]string{"Longitude"},
		[]string{"Address"},
		[]string{"City"},
		[]string{"PostalCode"},
		[]string{"Description"},
		[]string{"Facilities"},
	}

	var res model.Hotel
	idx := 0
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		item := "[" + strconv.Itoa(idx) + "]"
		idx += 1
		jsonparser.EachKey(data, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				res.ID = string(value)
				pp.Println("res", string(value))
				// case 1:
				// 	v, _ := jsonparser.ParseInt(value)
				// 	data.Tz = int(v)
				// case 2:
				// 	data.Ua, _ = value
				// case 3:
				// 	v, _ := jsonparser.ParseInt(value)
				// 	data.St = int(v)
			}
		}, paths...)
	})

	// pp.Println("res", string(data))

	// if err := json.NewDecoder(r.Body).Decode(&into); err != nil {
	// 	return err
	// }

}
