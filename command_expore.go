package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	url := "https://pokeapi.co/api/v2/location-area/"
	areaURL := url + args[0]

	data, err := c.pokeapiClient.GetLocationPokemonJson(areaURL)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + args[0] + "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}
