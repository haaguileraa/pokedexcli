package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func listen(c *config){
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
		clicmd, ok := c.commands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := clicmd.callback(c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}
