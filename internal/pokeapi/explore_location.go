package pokeapi

import (
	"io"
	"net/http"
	"encoding/json"
)

func (c *Client) ExploreLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	// if pageURL != nil {
	// 	url = *pageURL
	// }

	if val, exists := c.cache.Get(url); exists {
		locationDetails := Location{}
		err := json.Unmarshal(val, &locationDetails)
		if err != nil {
			return Location{}, nil
		}

		return locationDetails, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	
	locationDetails := Location{}
	err = json.Unmarshal(dat, &locationDetails)
	if err != nil {
		return Location{}, err
	}

	return locationDetails, nil
} // End GetLocations() func