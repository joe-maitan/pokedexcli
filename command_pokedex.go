package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) <= 0 {
		fmt.Println("You have not caught any pokemon yet. Get out there and explore!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, p := range cfg.pokedex {
		fmt.Printf(" - %s\n", p.Name)
	}

	return nil
} // End commandMapf() func