package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := fmt.Sprintf("%s%s/%s", pokeapiURL, pokemonEndpoint, pokemonName)

	data, err := c.DoRequest(url)

	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
} 
