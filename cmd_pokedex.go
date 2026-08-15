package main

import "fmt"

func commandPokedex(c *config, arg string) error {
	if len(c.pokedex) == 0 {
		return fmt.Errorf("Your Pokedex is currently empty")
	}
	fmt.Println("Your Pokedex:")
	for pokemon := range c.pokedex {
		fmt.Println("	-", pokemon)
	}
	return nil
}
