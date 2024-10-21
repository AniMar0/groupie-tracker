package TRC

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Artist represents the structure for each artist with additional info like dates and locations
type Artist struct {
	ID                       int      `json:"id"`
	Name                     string   `json:"name"`
	Image                    string   `json:"image"`
	Members                  []string `json:"members"`
	CreationDate             int      `json:"creationDate"`
	FirstAlbum               string   `json:"firstAlbum"`
	LocationsApi             string   `json:"locations"`
	DatesApi                 string   `json:"concertDates"`
	RelationsApi             string   `json:"relations"`
	Locations                []string
	Dates                    []string
	Relations                map[string][]string
	OtherLocationsInfos      OtherLocationsInfo
	OtherDatesInfos          OtherDatesInfo
	OtherDatesLocationsInfos OtherDatesLocationsInfo
}

// OtherInfo holds extra details like locations and dates for each artist
type OtherLocationsInfo struct {
	Locations []string `json:"locations"`
}
type OtherDatesInfo struct {
	Dates []string `json:"dates"`
}
type OtherDatesLocationsInfo struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Artists stores all artist data fetched from the API
var Artists []Artist

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

// FetchLocations retrieves the location data for the artist
func (ar *Artist) FetchLocations() {
	response, err := http.Get(ar.LocationsApi)
	if err != nil {
		fmt.Println("Error fetching locations:", err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading locations data:", err)
		return
	}

	// Unmarshal the JSON data into the artist's OtherInfo
	if err := json.Unmarshal(data, &ar.OtherLocationsInfos); err != nil {
		fmt.Println("Error unmarshalling locations data:", err)
	}
	ar.Locations = ar.OtherLocationsInfos.Locations
}

// FetchDates retrieves the concert dates data for the artist
func (ar *Artist) FetchDates() {
	response, err := http.Get(ar.DatesApi)
	if err != nil {
		fmt.Println("Error fetching dates:", err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading dates data:", err)
		return
	}

	// Unmarshal the JSON data into the artist's OtherInfo
	if err := json.Unmarshal(data, &ar.OtherDatesInfos); err != nil {
		fmt.Println("Error unmarshalling dates data:", err)
	}
	ar.Dates = ar.OtherDatesInfos.Dates
}

// FetchRelations retrieves the relations data (dates and locations combined) for the artist
func (ar *Artist) FetchRelations() {
	response, err := http.Get(ar.RelationsApi)
	if err != nil {
		fmt.Println("Error fetching relations:", err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading relations data:", err)
		return
	}

	// Unmarshal the JSON data into the artist's DatesLocations map
	if err := json.Unmarshal(data, &ar.OtherDatesLocationsInfos); err != nil {
		fmt.Println("Error unmarshalling relations data:", err)
	}
	ar.Relations = ar.OtherDatesLocationsInfos.DatesLocations
}
