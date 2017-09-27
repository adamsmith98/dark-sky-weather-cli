package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	apiKey string
	days int
	units string
)

func init() {
	flag.StringVar(&apiKey, "key", "", "Dark Sky API key")
	flag.StringVar(&apiKey, "k", "", "Dark Sky API key")
	flag.IntVar(&days, "days", 0, "Number of days")
	flag.IntVar(&days, "d", 0, "Number of days")
	flag.StringVar(&units, "units", "F", "Temperature units")
	flag.StringVar(&units, "u", "F", "Temperature units")

	flag.Parse()
}

func main() {
	geo, err := Locate()
	if err != nil {
		exitOnError(err)
	}

	var key string
	if apiKey != "" {
		err = ioutil.WriteFile("API_KEY", []byte(fmt.Sprintf("%s\n", apiKey)), 0666)
		fmt.Printf("Saved key %s to file API_KEY\n", apiKey)
		if err != nil {
			exitOnError(err)
		}
		key = apiKey
	} else {
		buffer, err := ioutil.ReadFile("API_KEY")
		if err != nil {
			exitOnError(err)
		}
		key = strings.TrimRight(string(buffer), "\n")
	}

	forecast, err := getForecast(fmt.Sprintf("https://api.darksky.net/forecast/%s/%f,%f", key, geo.Latitude, geo.Longitude))
	if err != nil {
		exitOnError(err)
	}

	if days > 0 {
		if days > 7 {
			days = 7
		}
		printDays(forecast, geo)

	} else {
		printNow(forecast, geo)
	}
}

func exitOnError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
