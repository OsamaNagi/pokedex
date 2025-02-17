package main

import "fmt"

func commandPokedex(config *Config, args []string) error {
	fmt.Println("Your Pokedex:")

	for _, pokemon := range config.Player {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
