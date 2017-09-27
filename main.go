package main

import (
	"fmt"
	"flag"
	"strings"
	"time"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

var (
	units string
	days int
	apiKey string
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

func init() {
	flag.StringVar(&units, "units", "F", "Temperature units")
	flag.IntVar(&days, "days", 0, "Number of days")
	flag.StringVar(&apiKey, "key", "", "Dark Sky API key")
	flag.Parse()
}

func main() {
	geo, err := Locate()
	if err != nil {
		fmt.Errorf("failed %s", err)
	}
	var key string
	if apiKey != "" {
		err = ioutil.WriteFile("API_KEY", []byte(fmt.Sprintf("%s\n", apiKey)), 0666)
		fmt.Printf("Saved key %s to file API_KEY\n", apiKey)
		if err != nil {
			fmt.Errorf("failed %s", err)
		}
		key = apiKey
	} else {
		buffer, err := ioutil.ReadFile("API_KEY")
		if err != nil {
			fmt.Errorf("failed %s", err)
		}
		key = strings.TrimRight(string(buffer), "\n")
	}
	forecast, err := getForecast(fmt.Sprintf("https://api.darksky.net/forecast/%s/%f,%f", key, geo.Latitude, geo.Longitude))
	if err != nil {
		fmt.Errorf("failed %s", err)
	}
	if days > 0 {
		if days > 7 {
			days = 7
		}
		fmt.Printf("Weather in %s, %s, %s\n", geo.City, geo.Region, geo.Country)
		fmt.Printf("Showing weather for %d days", days)
		for i := 0; i < days; i++ {
			t := time.Unix(forecast.Daily.Data[i].Time, 0)
			year, month, day := t.Date()
			weekday := t.Weekday().String()
			fmt.Printf("Forecast for %s %d %s %d\n", weekday, day, month.String(), year)
			fmt.Printf("High (°%s): %f\n", units, forecast.Daily.Data[i].TemperatureHigh)
			fmt.Printf("Low (°%s): %f\n", units, forecast.Daily.Data[i].TemperatureLow)
			fmt.Println()
		}
	} else {
		now := forecast.Currently.Temperature
		feels := forecast.Currently.ApparentTemperature
		high := forecast.Daily.Data[0].TemperatureHigh
		low := forecast.Daily.Data[0].TemperatureLow
		if units == "C" {
			now = convert(now)
			feels = convert(feels)
			high = convert(high)
			low = convert(low)
		}
		fmt.Printf("Weather in %s, %s, %s\n", geo.City, geo.Region, geo.Country)
		fmt.Printf("%s\n", forecast.Currently.Summary)
		fmt.Printf("Temperature now (°%s): %f\n", units, now)
		fmt.Printf("Feels like (°%s): %f\n", units, feels)
		fmt.Printf("High (°%s): %f\n", units, high)
		fmt.Printf("Low (°%s): %f\n", units, low)
		fmt.Printf("Wind speed (mph): %f\n", forecast.Currently.WindSpeed)
	}
}

func convert(fahrenheit float64) float64 {
	return (fahrenheit-32.0) * (5.0/9.0)
}
