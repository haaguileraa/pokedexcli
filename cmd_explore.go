package main

import (
	"fmt"
)

func commandExplore(c* config, arg string) error {
	if arg == "" {
		return fmt.Errorf("cannot retrieve empty area")
	}

	fmt.Printf("Exploring %s...\n", arg)
	resp, err := c.pokeapiClient.ExploreLocation(arg)
	
	if err != nil {
		return err
	}

	
	encounters := resp.PokemonEncounters
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
