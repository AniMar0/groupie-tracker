package TRC

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
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

func (a *Artist) Search(searchWord string) *Artist {
	if strings.Contains(strings.ToLower(a.Name), searchWord) ||
		strings.Contains(strings.ToLower(a.FirstAlbum), searchWord) ||
		strings.Contains(strconv.Itoa(a.CreationDate), searchWord) {
		return a
	}

	for membeId := range a.Members {
		if strings.Contains(strings.ToLower(a.Members[membeId]), searchWord) {
			return a
		}
	}

	for locationId := range Location.Index[a.ID-1].Locations {
		if strings.Contains(strings.ToLower(Location.Index[a.ID-1].Locations[locationId]), searchWord) {
			return a
		}
	}

	return nil
}

func (a *Artist) GetData() []string {
	var data []string
	data = append(data, a.Name+" - artist/band", (a.FirstAlbum)+" - First Album", strconv.Itoa(a.CreationDate)+" - Creation Date")

	for membeId := range a.Members {
		data = append(data, a.Members[membeId]+" - member")
	}

	for locationId := range Location.Index[a.ID-1].Locations {
		data = append(data, Location.Index[a.ID-1].Locations[locationId]+" - Location")
	}

	return data
}
