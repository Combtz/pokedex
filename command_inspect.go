package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemon := args[0]

	data, ok := c.caughtPokemon[pokemon]
	if !ok {
		fmt.Println("you have no caught this pokemon or pokemon name invalid")
		return nil
	}
	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %d\n", data.Height)
	fmt.Printf("Weight: %d\n", data.Weight)
	fmt.Println("Stats:")
	for _, stat := range data.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range data.Types {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}

	return nil
}
