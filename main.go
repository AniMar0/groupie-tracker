package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// API structure to store the base URLs for each API endpoint
type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

// Artist structure for each artist with DatesLocations added
type Artist struct {
	ID             int                 `json:"id"`
	Name           string              `json:"name"`
	Image          string              `json:"image"`
	StartYear      int                 `json:"start_year"`
	FirstAlbum     string              `json:"first_album"`
	Members        []string            `json:"members"`
	Locations      []string            `json:"locations"`
	DatesURL       string              `json:"dates"`
	Dates          []string            `json:"dates"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Location structure for concert locations
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// Date structure for concert dates
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Relation structure to link DatesLocations with artist
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Global variables to hold the API data
var (
	Apis      API
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
)

func main() {
	// Get the API URLs
	Apis.Url()

	// Fetch the data from all the endpoints
	Fetch()

	// Example: Display the DatesLocations for the first artist
	fmt.Println(Relations[1])
}

// Function to fetch data from the API and store it in the respective variables
func Fetch() {
	// Fetch and unmarshal artists
	ArtistsResponse, err := http.Get(Apis.Artists)
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}
	defer ArtistsResponse.Body.Close()
	ArtistsData, err := io.ReadAll(ArtistsResponse.Body)
	if err != nil {
		fmt.Println("Error reading artists data:", err)
		return
	}
	json.Unmarshal(ArtistsData, &Artists)

	// Fetch and unmarshal locations
	LocationsResponse, err := http.Get(Apis.Locations)
	if err != nil {
		fmt.Println("Error fetching locations:", err)
		return
	}
	defer LocationsResponse.Body.Close()
	LocationsData, err := io.ReadAll(LocationsResponse.Body)
	if err != nil {
		fmt.Println("Error reading locations data:", err)
		return
	}
	json.Unmarshal(LocationsData, &Locations)

	// Fetch and unmarshal dates
	DatesResponse, err := http.Get(Apis.Dates)
	if err != nil {
		fmt.Println("Error fetching dates:", err)
		return
	}
	defer DatesResponse.Body.Close()
	DatesData, err := io.ReadAll(DatesResponse.Body)
	if err != nil {
		fmt.Println("Error reading dates data:", err)
		return
	}
	json.Unmarshal(DatesData, &Dates)

	// Fetch and unmarshal relations (DatesLocations)
	RelationResponse, err := http.Get(Apis.Relation)
	if err != nil {
		fmt.Println("Error fetching relations:", err)
		return
	}
	defer RelationResponse.Body.Close()
	RelationData, err := io.ReadAll(RelationResponse.Body)
	if err != nil {
		fmt.Println("Error reading relations data:", err)
		return
	}
	json.Unmarshal(RelationData, &Relations)

	// Link DatesLocations to the respective artist
	// for i := range Artists {

	// 	if Artists[i].ID == Relations[i].ID {
	// 		Artists[i].DatesLocations = Relations[i].DatesLocations
	// 	}
	// 	if Artists[i].ID == Dates[i].ID {
	// 		Artists[i].Dates = Dates[i].Dates
	// 	}
	// 	if Artists[i].ID == Locations[i].ID {
	// 		Artists[i].Locations = Locations[i].Locations
	// 	}

	// }
}

// Fetch the API base URLs from the main endpoint
func (ap *API) Url() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
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

	json.Unmarshal(APIUrl, &Apis)
}
