package pokeapi

import (
	"fmt"
	"github.com/haaguileraa/pokedexcli/internal/pokecache"
	"io"
	"net/http"
	"time"
)


type Client struct {
	httpClient	http.Client
	cache		pokecache.Cache

}


func NewClient(timeout, interval time.Duration) Client {
	return Client {
		httpClient: http.Client {
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
}

func (c *Client) DoRequest(url string) ([]byte, error) {

	if cachedResp, ok := c.cache.Get(url); ok {
		return cachedResp, nil
	} 

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
