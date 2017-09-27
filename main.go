package main

import (
	"fmt"
	"flag"
	"strings"
	"time"
	"io/ioutil"
)

var (
	units string
	days int
	apiKey string
)

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
