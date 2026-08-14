package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetLocations(customURL string) (LocationsResponse, error) {
	var url string
	if customURL == "" {
		url = fmt.Sprintf("%s%s/?limit=%d", pokeapiURL, locationArea, defaultLimit)
	} else {
		url = customURL
	}

	data, err := c.DoRequest(url)

	if err != nil {
		return LocationsResponse{}, err
	}
		
	var resp LocationsResponse
	
	if err := json.Unmarshal(data, &resp); err != nil {
		return LocationsResponse{}, err
	}

	return resp, nil
}
