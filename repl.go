package main

import 	(
	"fmt"
	"os"
	"strings"
)


func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\n")
	supportedCommands := getSupportedCommands()
	for _, clicmd := range supportedCommands {
		fmt.Printf("%s: %s\n", clicmd.name, clicmd.description)
	}
	return nil
}

