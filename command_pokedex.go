package main

import "fmt"

func commandPokedex(c *config, args ...string) error {

	if len(c.caughtPokemon) == 0 {
		fmt.Println("No Pokemon Caught Yet, use Catch command")
		return nil
	}
	fmt.Println("Your pokedex:")
	for pokemon := range c.caughtPokemon {
		fmt.Println(" - " + pokemon)
	}
	return nil
}
