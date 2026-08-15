package main

import(
	"fmt"
	"os"
)

type cliCommand struct {
	name 		string
	description	string
	callback 	func(*config, string) error
}


func getSupportedCommands() map[string]cliCommand {
	return  map[string]cliCommand{
		"catch" : {
			name: "catch",
			description: "displays the interaction of trying to catch a Pokemon given the name",
			callback: commandCatch,
		},
		"exit" : {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"explore" : {
			name: "explore",
			description: "displays the Pokemons' names in a given location",
			callback: commandExplore,
		},

		"help" : {
			name: "help",
			description: "displays a help message",
			callback: commandHelp,
		},
		"map" : {
			name: "map",
			description: "displays the next 20 locations",
			callback: commandMap,
		},
		"mapb" : {
			name: "mapb",
			description: "displays the previous 20 locations",
			callback: commandMapBack,
		},
	}
}

func commandExit(c *config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, arg string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	fmt.Println()
	supportedCommands := getSupportedCommands()
	for _, clicmd := range supportedCommands {
		fmt.Printf("%s: %s\n", clicmd.name, clicmd.description)
	}
	return nil
}
