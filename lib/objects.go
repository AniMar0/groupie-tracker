package TRC

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func test1() {
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
}

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

func test2() {
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
}
