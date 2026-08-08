package main

type cliCommand struct {
	name 		string
	description	string
	callback 	func() error
}

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
	}
}
