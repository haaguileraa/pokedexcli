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

	var data []byte

	if cachedResp, ok := c.cache.Get(url); ok {
		data = cachedResp
	} else {
		resp, err := c.DoRequest(url)
		if err != nil {
			return ResponseLocations{}, err
		}
		data = resp
	}
		
	var resp ResponseLocations
	
	if err := json.Unmarshal(data, &resp); err != nil {
		return ResponseLocations{}, err
	}

	return resp, nil
}

func (c *Client) DoRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating get request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read request body: %w", err)
	}
	c.cache.Add(url, data)
	return data, nil
}
