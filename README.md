# Pokedex CLI

A command-line interface Pokedex application written in Go that allows users to explore Pokemon locations, catch Pokemon, and manage their collection.

## Features

- Explore Pokemon locations
- Catch Pokemon with a chance-based system
- Inspect caught Pokemon's details
- View your Pokedex inventory
- Caching system for API responses

## Commands

- `help`: Displays all available commands
- `map`: Display the next 20 Pokemon locations
- `mapb`: Display the previous 20 Pokemon locations
- `explore [location-name]`: View Pokemon that can be found in a specific location
- `catch [pokemon-name]`: Attempt to catch a Pokemon
- `inspect [pokemon-name]`: View detailed information about a caught Pokemon
- `pokedex`: List all Pokemon in your collection
- `exit`: Close the Pokedex application

## Technical Details

- Built in Go
- Uses the [PokeAPI](https://pokeapi.co/) for Pokemon data
- Implements an in-memory cache system with automatic cleanup
- Thread-safe cache implementation with mutex locks
- Command-line interface with real-time user input

## Cache System

The application implements a sophisticated caching system that:
- Stores API responses for 5 seconds
- Automatically cleans up expired entries
- Reduces API calls and improves response times
- Uses thread-safe operations for concurrent access

## How to Use

1. Start the application
2. Use the `map` command to view available locations
3. Use `explore [location-name]` to see Pokemon in that area
4. Try to catch Pokemon using `catch [pokemon-name]`
5. View your caught Pokemon with the `pokedex` command
6. Inspect Pokemon details using `inspect [pokemon-name]`

Note: Pokemon catching is based on a random chance system that factors in the Pokemon's base experience.