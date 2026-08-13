package main

import (
	"github.com/haaguileraa/pokedexcli/internal/pokeapi"
	"github.com/haaguileraa/pokedexcli/internal/pokecache"
	"time"
)

type config struct {
	pokeapiClient	pokeapi.Client
	cache		pokecache.Cache
	commands	map[string]cliCommand
	next		string
	previous	string
}


func NewConfig(timeout, interval time.Duration) *config {
	return &config{
		pokeapiClient: pokeapi.NewClient(timeout),
		cache: pokecache.NewCache(interval),
		commands: getSupportedCommands(),
		next: "",
		previous: "",
	}
}
