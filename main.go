package main

import (
	"time"

	"github.com/OsamaNagi/pokedex/internals"
)

func main() {
	config := &Config{
		Cache:  internals.NewCache(5 * time.Second),
		Player: make([]Pokemon, 0),
	}
	startRepl(config)
}
