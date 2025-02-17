package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
}

type PokemonResponse struct {
	Results  []Pokemon `json:"results"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
}

func commandMap(config *Config) error {
	baseURL := "https://pokeapi.co/api/v2/location-area"

	url := baseURL
	if config.Next != "" {
		url = config.Next
	}

	if v, ok := config.Cache.Get(url); ok {
		fmt.Println("**[Using cached response]**")
		var pokemonResp PokemonResponse
		if err := json.Unmarshal(v, &pokemonResp); err != nil {
			fmt.Println("error unmarshaling JSON:", err)
			return err
		}

		config.Next = pokemonResp.Next
		config.Previous = pokemonResp.Previous
		for _, location := range pokemonResp.Results {
			fmt.Println(location.Name)
		}

		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error making request:", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("error reading response body:", err)
		return err
	}

	var pokemonResp PokemonResponse
	if err := json.Unmarshal(body, &pokemonResp); err != nil {
		fmt.Println("error unmarshaling JSON:", err)
		return err
	}

	config.Next = pokemonResp.Next
	config.Previous = pokemonResp.Previous

	for _, location := range pokemonResp.Results {
		fmt.Println(location.Name)
	}

	config.Cache.Add(url, body)
	return nil
}

func commandMapb(config *Config) error {
	baseURL := "https://pokeapi.co/api/v2/location-area"

	url := baseURL
	if config.Previous != "" {
		url = config.Previous
	}

	if v, ok := config.Cache.Get(url); ok {
		fmt.Println("**[Using cached response]**")
		var pokemonResp PokemonResponse
		if err := json.Unmarshal(v, &pokemonResp); err != nil {
			fmt.Println("error unmarshaling JSON:", err)
			return err
		}

		config.Next = pokemonResp.Next
		config.Previous = pokemonResp.Previous
		for _, location := range pokemonResp.Results {
			fmt.Println(location.Name)
		}

		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error making request:", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("error reading response body:", err)
		return err
	}

	var pokemonResp PokemonResponse
	if err := json.Unmarshal(body, &pokemonResp); err != nil {
		fmt.Println("error unmarshaling JSON:", err)
		return err
	}

	config.Next = pokemonResp.Next
	config.Previous = pokemonResp.Previous

	for _, location := range pokemonResp.Results {
		fmt.Println(location.Name)
	}

	config.Cache.Add(url, body)

	return nil
}
