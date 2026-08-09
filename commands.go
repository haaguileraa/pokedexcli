package main

import(
	"fmt"
	"os"
)

type cliCommand struct {
	name 		string
	description	string
	callback 	func(*config) error
}

const defaultLimit = 20
const pokeapiURL = "https://pokeapi.co/api/v2"

func getSupportedCommands() map[string]cliCommand {
	return  map[string]cliCommand{
		"exit" : {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
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

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\n")
	supportedCommands := getSupportedCommands()
	for _, clicmd := range supportedCommands {
		fmt.Printf("%s: %s\n", clicmd.name, clicmd.description)
	}
	return nil
}
