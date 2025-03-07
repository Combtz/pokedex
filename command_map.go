package main

import (
	"fmt"

	"github.com/Combtz/pokedex/internal/pokeapi"
)

func commandMap(c *config) error {
	if c.next == "" {
		data, err := pokeapi.GetLocationAreaJSON("https://pokeapi.co/api/v2/location-area/")
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
	data, err := pokeapi.GetLocationAreaJSON(c.next)
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
func commandMapb(c *config) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	data, err := pokeapi.GetLocationAreaJSON(c.previous)
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
