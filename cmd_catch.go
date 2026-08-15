package main

import (
	"fmt"
	"math/rand/v2"
)

const baseExperienceFactor = 1000.0 // highest Base experience Blissey = 608

func commandCatch(c* config, pokemonName string) error {
	
	if pokemonName == "" {
		return fmt.Errorf("cannot try to catch without a target")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := c.pokeapiClient.GetPokemon(pokemonName)
	
	if err != nil {
		return err
	}
	probability := float64(pokemon.BaseExperience) / baseExperienceFactor

	if rand.Float64() > probability {
		c.pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}

