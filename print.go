package main

import (
	"fmt"
	"time"
)

func printDays(forecast Forecast, geo Geocode) {
	fmt.Printf("Weather in %s, %s, %s\n", geo.City, geo.Region, geo.Country)
	fmt.Printf("Showing weather for %d days\n", days)
	fmt.Println()

	for i := 0; i < days; i++ {
		t := time.Unix(forecast.Daily.Data[0].Time, 0)
		year, month, day := t.Date()
		weekday := t.Weekday().String()
		fmt.Printf("Forecast for %s %d %s %d\n", weekday, day, month.String(), year)
		fmt.Printf("High (°%s): %f\n", units, forecast.Daily.Data[0].TemperatureHigh)
		fmt.Printf("Low (°%s): %f\n", units, forecast.Daily.Data[0].TemperatureLow)
		fmt.Println()
	}
}

func printNow(forecast Forecast, geo Geocode) {
	now := forecast.Currently.Temperature
	feels := forecast.Currently.ApparentTemperature
	high := forecast.Daily.Data[0].TemperatureHigh // High today
	low := forecast.Daily.Data[0].TemperatureLow // Low today

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

func convert(fahrenheit float64) float64 {
	return (fahrenheit-32.0) * (5.0/9.0)
}
