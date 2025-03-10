package main

import (
	"fmt"
)

func commandMap(c *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if c.next == "" {
		data, err := c.pokeapiClient.GetLocationAreaJSON(url)
		if err != nil {
			return err
		}
		c.next = data.Next
		c.previous = data.Previous

		for _, location := range data.Results {
			fmt.Println(location.Name)
		}
		return nil
	}
	data, err := c.pokeapiClient.GetLocationAreaJSON(c.next)
	if err != nil {
		return err
	}
	c.next = data.Next
	c.previous = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	return nil
}
func commandMapb(c *config, args ...string) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	data, err := c.pokeapiClient.GetLocationAreaJSON(c.previous)
	if err != nil {
		return err
	}
	c.next = data.Next
	c.previous = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	return nil
}
