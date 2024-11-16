package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiResponse struct {
	Results struct {
		Sunrise                 string `json:"sunrise"`
		Sunset                  string `json:"sunset"`
		SolarNoon               string `json:"solar_noon"`
		DayLength               int    `json:"day_length"`
		CivilTwilightBegin      string `json:"civil_twilight_begin"`
		CivilTwilightEnd        string `json:"civil_twilight_end"`
		NauticalTwilightBegin   string `json:"nautical_twilight_begin"`
		NauticalTwilightEnd     string `json:"nautical_twilight_end"`
		AstronomicalTwilightBegin string `json:"astronomical_twilight_begin"`
		AstronomicalTwilightEnd   string `json:"astronomical_twilight_end"`
	} `json:"results"`
	Status string `json:"status"`
}

func main() {
	// Location details
	latitude := "58.3776"
	longitude := "26.7290"
	locationName := "Tartu, Estonia"

	// API URL
	apiURL := fmt.Sprintf("https://api.sunrise-sunset.org/json?lat=%s&lng=%s&formatted=0", latitude, longitude)

	// Make an HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the response
	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Check status
	if apiResponse.Status != "OK" {
		fmt.Println("API returned an error:", apiResponse.Status)
		return
	}

	// Define Tartu's timezone
	location, err := time.LoadLocation("Europe/Tallinn")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Helper function to parse and format time
	formatTime := func(utcTime string) string {
		parsedTime, err := time.Parse(time.RFC3339, utcTime)
		if err != nil {
			return "Invalid time"
		}
		localTime := parsedTime.In(location)
		return localTime.Format("15:04") // Only show hours and minutes
	}

	// Print location and times
	fmt.Println("Location:", locationName)
	fmt.Println("Sunrise:", formatTime(apiResponse.Results.Sunrise))
	fmt.Println("Sunset:", formatTime(apiResponse.Results.Sunset))
	fmt.Println("Solar Noon:", formatTime(apiResponse.Results.SolarNoon))
	fmt.Println("Civil Twilight Begin:", formatTime(apiResponse.Results.CivilTwilightBegin))
	fmt.Println("Civil Twilight End:", formatTime(apiResponse.Results.CivilTwilightEnd))
	fmt.Println("Nautical Twilight Begin:", formatTime(apiResponse.Results.NauticalTwilightBegin))
	fmt.Println("Nautical Twilight End:", formatTime(apiResponse.Results.NauticalTwilightEnd))
	fmt.Println("Astronomical Twilight Begin:", formatTime(apiResponse.Results.AstronomicalTwilightBegin))
	fmt.Println("Astronomical Twilight End:", formatTime(apiResponse.Results.AstronomicalTwilightEnd))
}
