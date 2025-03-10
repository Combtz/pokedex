package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	pokemonURL := baseURL + args[0]
	data, err := c.pokeapiClient.GetPokemonJson(pokemonURL)
	if err != nil {
		return err
	}
	fmt.Println("Throwing a Pokeball at " + args[0] + "...")
	chance := rand.Float64()
	catchChance := 1.0 / (1.0 + float64(data.BaseExperience)/100.0)
	if chance < float64(catchChance) {
		fmt.Println(args[0] + " was caught!")
	} else {
		fmt.Println(args[0] + " escaped!")
	}
	return nil
}
