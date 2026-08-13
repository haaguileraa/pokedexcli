package pokeapi

import (
	"github.com/haaguileraa/pokedexcli/internal/pokecache"
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
