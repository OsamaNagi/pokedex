package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonStats struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func commandInspect(config *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name to inspect")
		return nil
	}

	pokemonName := args[0]
	found := false

	for _, poke := range config.Player {
		if poke.Name == pokemonName {
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("You haven't caught %s yet\n", pokemonName)
		return nil
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	var body []byte
	if cached, ok := config.Cache.Get(url); ok {
		fmt.Println("**[Using cached response]**")
		body = cached
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		config.Cache.Add(url, body)
	}

	var pokemon PokemonStats
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
