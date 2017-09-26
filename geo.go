package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

const url string = "https://freegeoip.net/json/"

type Geocode struct {
	Country string `json:"country_name"`
	Region string `json:"region_name"`
	City string `json:"city"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func Locate() (geocode Geocode, err error) {
	res, err := http.Get(url)

	defer res.Body.Close()

	if err != nil {
		return geocode, fmt.Errorf("failed %s", err)
	}
	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&geocode); err != nil {
		return geocode, fmt.Errorf("failed %s", err)
	}
	return geocode, nil
}
