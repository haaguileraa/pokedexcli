package main

import "fmt"


func commandInspect(c *config, pokemonName string) error {
	if pokemonName == "" {
		return fmt.Errorf("cannot retreive empty pokemon name")
	}

	pokemon, ok := c.pokedex[pokemonName]

	if !ok {
		return fmt.Errorf("you have not caught that pokemon");
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Println("	-", pokemonType.Type.Name)
	}
	return nil
}
