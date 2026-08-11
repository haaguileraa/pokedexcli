package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) GetLocations(customURL string) (ResponseLocations, error) {
	var url string
	if customURL == "" {
		url = fmt.Sprintf("%s%s/?limit=%d", pokeapiURL, locationArea, defaultLimit)
	} else {
		url = customURL
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocations{}, fmt.Errorf("error creating get request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocations{}, fmt.Errorf("error doing request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocations{}, fmt.Errorf("could not read request body: %w", err)
	}
	
	var resp ResponseLocations
	
	if err := json.Unmarshal(data, &resp); err != nil {
		return ResponseLocations{}, err
	}

	return resp, nil
}
