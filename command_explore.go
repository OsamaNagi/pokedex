package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}

type LocationAreaResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func commandExplore(config *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a location area name to explore.")
		return nil
	}

	baseURL := "https://pokeapi.co/api/v2/location-area/" + args[0]

	if cached, ok := config.Cache.Get(baseURL); ok {
		fmt.Println("**[Using cached response]**")
		return processLocationAreaResponse(cached, args[0])
	}

	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("error making request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return fmt.Errorf("location area '%s' not found", args[0])
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return err
	}

	config.Cache.Add(baseURL, body)
	return processLocationAreaResponse(body, args[0])
}

func processLocationAreaResponse(data []byte, args string) error {
	var response LocationAreaResponse
	if err := json.Unmarshal(data, &response); err != nil {
		fmt.Println("error unmarshaling JSON:", err)
		return err
	}

	fmt.Printf("Exploring location area: %s\n", args)
	fmt.Println("Found Pokemon:")
	for _, encounter := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	if len(response.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found in this area!")
	}

	return nil
}
