package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Forecast struct {
	Latitutde float64 `json:"latitutde"`
	Longitude float64 `json:"longitude"`
	Currently Point `json:"currently"`
	Timezone string `json:"timezone"`
	Minutely Block `json:"minutely"`
	Hourly Block `json:"hourly"`
	Daily Block `json:"daily"`
	Alerts []Alert `json:"alerts"`
	Flags Flags `json:"flags"`
}

type Point struct {
	ApparentTemperature float64 `json:"apparentTemperature"`
	ApparentTemperatureHigh float64 `json:apparentTemperatureHigh"`
	ApparentTemperatureHighTime int64 `json:"apparentTemperatureHighTime"`
	ApparentTemperatureLow float64 `json:apparentTemperatureLow"`
	ApparentTemperatureLowTime int64 `json:"apparentTemperatureLowTime"`
	CloudCover float64 `json:"cloudCover"`
	Humidity float64 `json:"humidity"`
	Icon string `json:"icon"`
	PrecipAccumulation float64 `json:"precipAccumulation"`
	PrecipIntensity float64 `json:"precipIntensity"`
	PrecipIntensityMax float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime int64 `json:"precipIntensityMaxTime"`
	PrecipType string `json:"precipType"`
	Pressure float64 `json:"pressure"`
	Summary string `json:"summary"`
	SunriseTime int64 `json:"sunriseTime"`
	SunsetTime int64 `json:"sunsetTime"`
	Temperature float64 `json:"temperature"`
	TemperatureHigh float64 `json:"temperatureHigh"`
	TemperatureHighTime float64 `json:"temperatureHighTime"`
	TemperatureLow float64 `json:"temperatureLow"`
	TemperatureLowTime float64 `json:"temperatureLowTime"`
	Time int64 `json:"time"`
	WindBearing float64 `json:"windBearing"`
	WindGust float64 `json:"windGust"`
	WindSpeed float64 `json:"windSpeed"`
}

type Block struct {
	Data []Point `json:"data"`
	Summary string `json:"summary"`
	Icon string `json:"icon"`
}

type Alert struct {
	Description string `json:"description"`
	Expires int64 `json:"expires"`
	Regions []string `json:"regions"`
	Severity string `json:"severity"`
	Time int64 `json:"time"`
	Title string `json:"title"`
	uri string `json:"uri"`
}

type Flags struct {
	Sources []string `json:"sources"`
	Units string `json:"units"`
}

func getForecast(url string) (forecast Forecast, err error) {
	res, err := http.Get(url)

	defer res.Body.Close()

	if err != nil {
		return forecast, fmt.Errorf("failed %s", err)
	}
	if res.Status != "200 OK" {
		return forecast, fmt.Errorf("failed status %s", res.Status)
	}

	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&forecast); err != nil {
		return forecast, fmt.Errorf("failed %s", err)
	}

	return forecast, nil
}
