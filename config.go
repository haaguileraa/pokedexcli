package main

import (
	"github.com/haaguileraa/pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient	pokeapi.Client
	commands	map[string]cliCommand
	next		string
	previous	string
}


func NewConfig(timeout, interval time.Duration) *config {
	return &config{
		pokeapiClient: pokeapi.NewClient(timeout, interval),
		commands: getSupportedCommands(),
		next: "",
		previous: "",
	}
}
