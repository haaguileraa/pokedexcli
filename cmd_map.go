package main

import (
	"fmt"
)

func commandMap(c* config) error {

	resp, err := c.pokeapiClient.GetLocations(c.next)
	
	if err != nil {
		return err
	}
	
	c.next = resp.Next
	c.previous = resp.Previous
	results := resp.Results
	
	for _, result := range results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapBack(c *config) error {
	oldNext := c.next
	c.next = c.previous
	c.previous = oldNext
	if c.next == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	return commandMap(c)
}
