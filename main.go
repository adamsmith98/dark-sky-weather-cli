package main

import (
	"fmt"
	"os"
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
	ApparentTemperatureMax float64 `json:apparentTemperatureMax"`
	ApparentTemperatureMaxTime int64 `json:"apparentTemperatureMaxTime"`
	ApparentTemperatureMin float64 `json:apparentTemperatureMin"`
	ApparentTemperatureMinTime int64 `json:"apparentTemperatureMinTime"`
	CloudCover float64 `json:"cloudCover"`
	DewPoint float64 `json:"dewPoint"`
	Humidity float64 `json:"humidity"`
	Icon string `json:"icon"`
	MoonPhase float64 `json:"moonPhase"`
	NearestStormBearing float64 `json:nearestStormBearing"`
	NearestStormDistance float64 `json:"nearestStormDistance"`
	Ozone float64 `json:"ozone"`
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
	TemperatureMax float64 `json:"temperatureMax"`
	TemperatureMaxTime float64 `json:"temperatureMaxTime"`
	TemperatureMin float64 `json:"temperatureMin"`
	TemperatureMinTime float64 `json:"temperatureMinTime"`
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

	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&forecast); err != nil {
		return forecast, fmt.Errorf("failed %s", err)
	}
	if err != nil {
		return forecast, fmt.Errorf("failed %s", err)
	}

	fmt.Println(forecast)
	return forecast, err
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please provide a url")
	} else {
		getForecast(args[0])
	}
}
