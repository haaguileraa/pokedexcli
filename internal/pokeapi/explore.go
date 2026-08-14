package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ExploreLocation(area string) (ExplorationResponse, error) {
	if area == "" {
		return ExplorationResponse{}, fmt.Errorf("cannot retrieve empty area")
	}
	url := fmt.Sprintf("%s%s/%s", pokeapiURL, locationArea, area)

	data, err := c.DoRequest(url)

	if err != nil {
		return ExplorationResponse{}, err
	}
		
	var resp ExplorationResponse
	
	if err := json.Unmarshal(data, &resp); err != nil {
		return ExplorationResponse{}, err
	}

	return resp, nil
}
