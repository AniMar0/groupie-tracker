package TRC

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Artists stores all artist data fetched from the API
var (
	Artists  []Artist
	Location LocationsData
)

// FetchArtists retrieves the list of artists from the API
func FetchArtists() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading artists data:", err)
		return
	}

	// Unmarshal the JSON data into the Artists slice
	if err := json.Unmarshal(data, &Artists); err != nil {
		fmt.Println("Error unmarshalling artists data:", err)
	}
}

func FetchLocations() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading artists data:", err)
		return
	}

	// Unmarshal the JSON data into the Artists slice
	if err := json.Unmarshal(data, &Location); err != nil {
		fmt.Println("Error unmarshalling artists data:", err)
	}
}
