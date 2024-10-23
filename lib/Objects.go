package TRC

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Artist represents the structure for each artist with additional info like dates and locations
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsApi string   `json:"locations"`
	DatesApi     string   `json:"concertDates"`
	RelationsApi string   `json:"relations"`
	Locations    Locations
	Dates        Dates
	Relations    Relations
}

type LocationsData struct {
	Index []Locations `json:"index"`
}

// OtherInfo holds extra details like locations and dates for each artist
type Locations struct {
	Locations []string `json:"locations"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
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
	if err := json.Unmarshal(data, &ar.Locations); err != nil {
		fmt.Println("Error unmarshalling locations data:", err)
	}
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
	if err := json.Unmarshal(data, &ar.Dates); err != nil {
		fmt.Println("Error unmarshalling dates data:", err)
	}
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
	if err := json.Unmarshal(data, &ar.Relations); err != nil {
		fmt.Println("Error unmarshalling relations data:", err)
	}
}

func (a *Artist) FetchOtherData() {
	a.FetchLocations()
	a.FetchDates()
	a.FetchRelations()
}
