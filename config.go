package main

import (
	"github.com/haaguileraa/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient	pokeapi.Client
	next		string
	previous	string
}
