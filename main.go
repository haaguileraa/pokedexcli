package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			fmt.Println("sacanner could not receive more tokens")
			break
		}
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		command := cleanedInput[0]
		supportedCommands := getSupportedCommands()
		clicmd, ok := supportedCommands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := clicmd.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
