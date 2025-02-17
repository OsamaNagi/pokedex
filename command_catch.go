package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type BaseExperience struct {
	BaseExperience int `json:"base_experience"`
}

func commandCatch(config *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a location area name to explore.")
		return nil
	}

	baseURL := "https://pokeapi.co/api/v2/pokemon/" + args[0]

	if cached, ok := config.Cache.Get(baseURL); ok {
		fmt.Println("**[Using cached response]**")
		return processBaseExperienceResponse(config, cached, args[0])
	}

	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("error making request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return fmt.Errorf("pokemon '%s' not found", args[0])
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return err
	}

	config.Cache.Add(baseURL, body)
	return processBaseExperienceResponse(config, body, args[0])
}

func processBaseExperienceResponse(config *Config, data []byte, args string) error {
	var response BaseExperience
	if err := json.Unmarshal(data, &response); err != nil {
		fmt.Println("error unmarshaling JSON:", err)
		return err
	}

	fmt.Printf("Throwing a pokeball at %s...\n", args)

	experience := rand.Intn(response.BaseExperience)

	if experience > 50 {
		fmt.Printf("%s escaped!\n", args)
	} else {
		fmt.Printf("%s was caught!\n", args)
		fmt.Println("You may now inspect it with the inspect command.")
		config.Player = append(config.Player, Pokemon{Name: args})
	}

	return nil
}
