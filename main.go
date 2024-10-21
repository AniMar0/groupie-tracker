package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
// Artist structure for each artist with DatesLocations added
type Artist struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	StartYear  int      `json:"start_year"`
	FirstAlbum string   `json:"first_album"`
	Members    []string `json:"members"`
	Loca       string   `json:"locations"`
	Locations  Location
	Dates      []string `json:"dates"`
}

// Location structure for concert locations
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Global variables to hold the API data
var Artists []Artist


func main() {

	// Fetch artist
	FetchArtists()
	Artists[0].FetchLocation()

	// Display the first artist as a test
	if len(Artists) > 0 {
		fmt.Println("First Artist:", Artists[0].Locations.ID)
	}
}

// FetchArtists retrieves artist data from the API and stores it
func FetchArtists() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}
	defer response.Body.Close()

	ArtistsData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading artists data:", err)
		return
	}

	if err := json.Unmarshal(ArtistsData, &Artists); err != nil {
		fmt.Println("Error unmarshalling artists data:", err)
	}
}

// FetchLocations retrieves location data and links it to the artists
func (ar *Artist) FetchLocation() {
	response, err := http.Get(ar.Loca)
	if err != nil {
		fmt.Println("Error fetching API URL:", err)
		return
	}
	defer response.Body.Close()

	APIUrl, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading API URL data:", err)
		return
	}

	if err := json.Unmarshal(APIUrl, &ar.Locations); err != nil {
		fmt.Println("Error unmarshalling API URL data:", err)
		return
	}
}
