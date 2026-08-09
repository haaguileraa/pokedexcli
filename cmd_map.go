package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

)

func commandMap(c* config) error {
	if c.next == "" {
		c.next = fmt.Sprintf("%s%s/?limit=%d", pokeapiURL, "/location-area", defaultLimit)
	}
	res, err := http.Get(c.next)
	if err != nil {
		return fmt.Errorf("error creating get request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not read request body: %w", err)
	}
	
	var resp Response
	
	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}
	fmt.Println(c.next)
	fmt.Println(c.previous)
	c.next = resp.Next
	c.previous = resp.Previous
	results := resp.Results
	
	for _, result := range results {
	if name, ok := result["name"]; ok {
		fmt.Println(name)
		}
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
