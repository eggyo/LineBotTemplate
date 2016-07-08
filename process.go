package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func callGoogleElev(lat float64, lon float64) string {
	resp, err := http.Get("https://maps.googleapis.com/maps/api/elevation/json?locations=" + FloatToString(lat) + "," + FloatToString(lon) + "&key=AIzaSyAn9cWoce9zGEfGjDzMg6r_uTTUw3WoMOg")
	if err != nil {
		println(err.Error())
		return string(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	elev, err := getElev([]byte(body))

	return FloatToString(elev.Results[0].Elev)
}

type ELEV struct {
	Elev float64 `json:"elevation"`
	Res  float64 `json:"resolution"`
	Loc  struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"lng"`
	} `json:"location"`
}

type ResultElev struct {
	Results []ELEV `json:"results"`
	Status  string `json:"status"`
}

func getElev(body []byte) (*ResultElev, error) {
	var s = new(ResultElev)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
